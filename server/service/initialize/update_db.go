package initialize

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
)

const (
	UpdateMysql       = "mysql"
	UpdateDataFailed  = "\n[%v] --> %v 更新数据失败! [err]: %+v"
	UpdateTableFailed = "\n[%v] --> %v 更新表结构失败! [err]: %+v"
	UpdateDataSuccess = "\n[%v] --> %v 更新数据成功!"
)

// TypedDbUpdateHandler 执行传入的 updater
type TypedDbUpdateHandler interface {
	UpdateTables(ctx context.Context, updaters updaterSlice) error // 更新表
	UpdateData(ctx context.Context, updaters updaterSlice) error   // 更新数据
}

// Updater 提供 updater/*/init() 使用的接口，每个 updater 完成一个初始化过程
type Updater interface {
	// 需要初始化注册名，可以是表名，也可以是其他
	UpdaterName() string
	// ！注意：这里是merge表结构，不会删除原有字段
	UpdateTable(ctx context.Context) (next context.Context, err error)
	// ！注意：这里主要是update数据，如果使用insert插入关联表，需要单独建立关联表updater
	UpdateData(ctx context.Context) (next context.Context, err error)
}

// orderedUpdater 组合一个顺序字段，用来排序
type orderedUpdater struct {
	order int
	Updater
}

// updaterSlice 供 updater 排序依赖时使用
type updaterSlice []*orderedUpdater

var (
	updaters     updaterSlice
	updaterCache map[string]*orderedUpdater
)

/*
	！注意：更新的逻辑，初始化也要对应同步，确保部署给其他企业时，无需重复初始化
*/
// RegisterUpdate 注册要执行的初始化过程，UpdateDB() 时根据注册的 updater 进行初始化
func RegisterUpdate(order int, up Updater) {
	if updaters == nil {
		updaters = updaterSlice{}
	}
	if updaterCache == nil {
		updaterCache = map[string]*orderedUpdater{}
	}
	name := up.UpdaterName()
	if _, existed := updaterCache[name]; existed {
		// 表名冲突，会导致程序不能正确运行，init函数在main函数前运行，出错后禁止启动
		panicStr := fmt.Sprintf("UpdaterName conflict on %s", name)
		panic(panicStr)
	}
	oup := orderedUpdater{order, up}
	updaters = append(updaters, &oup)
	updaterCache[name] = &oup
}

/* ------ * service * ------ */

type UpdateDBService interface {
	CheckUpdate() (string, string, error)
	ClearUpdater()
	UpdateDB() error
}

type UpdateDBServiceImpl struct{}

// 检查是否需要升级
func (*UpdateDBServiceImpl) CheckUpdate() (oldVersion, newVersion string, err error) {
	db := global.OMS_DB
	sysVersion := &system.SysVersion{}
	err = db.Where("version_code = ?", "oms_version").First(sysVersion).Error
	if err != nil {
		return "", "", err
	}
	return sysVersion.Version, global.OMS_CONFIG.Version, err
}

// 已经升级，重启服务后，清除 updaters
func (s *UpdateDBServiceImpl) ClearUpdater() {
	updaters = updaterSlice{}
	updaterCache = map[string]*orderedUpdater{}
}

// 升级的前提是，部署了代码，部署代码一定会重新启动server，需要更新表结构，更新必要数据
func (s *UpdateDBServiceImpl) UpdateDB() (err error) {
	ctx := context.Background()
	if len(updaters) == 0 {
		return errors.New("升级任务列表为空，请检查是否已执行完成")
	}
	/*
		保证有依赖的 initializer 排在后面执行
		Note: 若 initializer 只有单一依赖，可以写为 B=A+1, C=A+1; 由于 BC 之间没有依赖关系，所以谁先谁后并不影响初始化
		若存在多个依赖，可以写为 C=A+B, D=A+B+C, E=A+1;
		C必然>A|B，因此在AB之后执行，D必然>A|B|C，因此在ABC后执行，而E只依赖A，顺序与CD无关，因此E与CD哪个先执行并不影响
	*/
	sort.Sort(&updaters)

	var updateHandler TypedDbUpdateHandler
	switch global.OMS_CONFIG.System.DbType {
	case "mysql":
		updateHandler = NewMysqlUpdateHandler()
	default:
		updateHandler = NewMysqlUpdateHandler()
	}

	if err = updateHandler.UpdateTables(ctx, updaters); err != nil {
		return err
	}
	global.OMS_LOG.Info("更新表结构成功")
	if err = updateHandler.UpdateData(ctx, updaters); err != nil {
		return err
	}
	global.OMS_LOG.Info("更新数据成功")

	updaters = updaterSlice{}
	updaterCache = map[string]*orderedUpdater{}

	return err
}

/* -- sortable interface -- */

func (a updaterSlice) Len() int {
	return len(a)
}

func (a updaterSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a updaterSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
