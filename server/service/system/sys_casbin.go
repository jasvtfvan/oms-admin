package system

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/model/system/request"
	"go.uber.org/zap"
)

/*
https://casbin.org/zh/editor/
RBAC with Domains

# Model
```
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
```

#Policy
```
p, admin, domain1, data1, read
p, admin, domain1, data1, write
p, admin, domain2, data2, read
p, admin, domain2, data2, write

g, alice, admin, domain1
g, bob, admin, domain2
```

#Request
```
# r是入参，p是policy
alice, domain1, data1, read
# g(r.sub, p.sub, r.dom) -> g(alice, ?, domain1) -> admin
# g(alice, admin, domain1) -> (alice, domain1) or (admin, domain1)
# alice, domain1, data1, read -> (alice, domain1, data1, read) or (admin, domain1, data1, read)
```

#Enforcement Result
```
true
```

举例：
alice就是用户名，admin就是角色，domain1就是群组，data1就是接口，read就是请求接口的方法POST/GET等等
*/

type CasbinInstance struct {
	CasbinService
}

var CasbinServiceApp = &CasbinInstance{
	CasbinService: new(CasbinServiceImpl),
}

type CasbinService interface {
	GetPolicyByCasbinInfo(roleCode, groupCode string, casbinInfos []request.CasbinInfo) [][]string
	GetCasbinInfoByPolicy(rules [][]string) []request.CasbinInfo
	GetPolicyInEnforcer(roleCode, groupCode string) [][]string
	Casbin() *casbin.SyncedCachedEnforcer
}

type CasbinServiceImpl struct{}

// 根据casbinInfo获取policy（去重）
func (csi *CasbinServiceImpl) GetPolicyByCasbinInfo(roleCode, groupCode string, casbinInfos []request.CasbinInfo) [][]string {
	rules := [][]string{}
	//做权限去重处理
	deduplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := roleCode + groupCode + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, []string{roleCode, groupCode, v.Path, v.Method})
		}
	}
	return rules
}

// 根据policy的rule获取casbinInfo（去重）
func (csi *CasbinServiceImpl) GetCasbinInfoByPolicy(rules [][]string) []request.CasbinInfo {
	casbinInfos := []request.CasbinInfo{}
	// 去重处理
	deduplicateMap := make(map[string]bool)
	for _, v := range rules {
		key := v[0] + v[1] + v[2] + v[3]
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			casbinInfos = append(casbinInfos, request.CasbinInfo{
				Path:   v[2],
				Method: v[3],
			})
		}
	}
	return casbinInfos
}

// 通过roleCode和groupCode获取policy
func (csi *CasbinServiceImpl) GetPolicyInEnforcer(roleCode, groupCode string) [][]string {
	e := csi.Casbin()
	var target [][]string
	rules := e.GetFilteredPolicy(0, roleCode)
	// 去重处理
	deduplicateMap := make(map[string]bool)
	for _, v := range rules {
		key := v[0] + v[1] + v[2] + v[3]
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			if v[1] == groupCode {
				target = append(target, v)
			}
		}
	}
	return target
}

// 批量保存，先删除存在的，再批量插入所有
func (csi *CasbinServiceImpl) BatchSavePolicies(roleCode, groupCode string, casbinInfos []request.CasbinInfo) error {
	rules := csi.GetPolicyByCasbinInfo(roleCode, groupCode, casbinInfos)
	rulesNamed := []system.SysCasbin{}
	for _, v := range rules {
		rulesNamed = append(rulesNamed, system.SysCasbin{
			Ptype: "p",
			V0:    v[0],
			V1:    v[1],
			V2:    v[2],
			V3:    v[3],
		})
	}

	tx := global.OMS_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := casbinDao.BatchDelete(tx, roleCode, groupCode); err != nil {
		tx.Rollback()
		return err
	}
	if err := casbinDao.BatchInsert(tx, rulesNamed); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	/*
		更新缓存
		不通过e.savePolicy，通过上边的事务保存，既确保原子性，又防止整个缓存同步到DB引发性能问题。（savePolicy会将整个缓存同步到DB）
		不通过e.loadPolicy，通过下边的RemovePolicy和AddPolicy，防止整个DB同步到缓存引发性能问题。（loadPolicy会将整个DB同步到缓存）
	*/
	e := csi.Casbin()
	var toRemove [][]string = csi.GetPolicyInEnforcer(roleCode, groupCode)
	_, err := e.RemovePolicies(toRemove)
	if err != nil {
		return err
	}
	_, err = e.AddPolicies(rules)
	if err != nil {
		return err
	}
	return nil
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	casbinOnce           sync.Once
)

func (*CasbinServiceImpl) Casbin() *casbin.SyncedCachedEnforcer {
	casbinOnce.Do(func() {
		sysCasbin := &system.SysCasbin{}
		adapter, err :=
			gormadapter.NewAdapterByDBWithCustomTable(global.OMS_DB, sysCasbin, sysCasbin.TableName())
		if err != nil {
			global.OMS_LOG.Error("适配数据库失败，请检查casbin表是否为InnoDB引擎", zap.Error(err))
			return
		}
		txt := `
		[request_definition]
		r = sub, dom, obj, act

		[policy_definition]
		p = sub, dom, obj, act

		[role_definition]
		g = _, _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
		`
		md, err := model.NewModelFromString(txt)
		if err != nil {
			global.OMS_LOG.Error("字符串加载模型失败，请检查格式是否正确", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(md, adapter)
		syncedCachedEnforcer.SetExpireTime(3600) // 设置缓存3600秒(1小时)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}
