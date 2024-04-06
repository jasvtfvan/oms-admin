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
	FreshCasbin() error
	Casbin() *casbin.SyncedCachedEnforcer
}

type CasbinServiceImpl struct{}

func (csi *CasbinApiServiceImpl) BatchSavePolicies(roleCode, groupCode string, casbinInfos []request.CasbinInfo) error {
	var casbinRules []system.SysCasbin
	for _, v := range casbinInfos {
		casbinRules = append(casbinRules, system.SysCasbin{
			Ptype: "p",
			V0:    roleCode,
			V1:    groupCode,
			V2:    v.Path,
			V3:    v.Method,
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
	if err := casbinDao.BatchInsert(tx, casbinRules); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	// 操作 e delete缓存，insert缓存
	return nil
}

func (csi *CasbinServiceImpl) FreshCasbin() (err error) {
	e := csi.Casbin()
	err = e.LoadPolicy()
	return err
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
