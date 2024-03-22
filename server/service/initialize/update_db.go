package initialize

import (
	"context"
	"errors"
	"fmt"

	"github.com/jasvtfvan/oms-admin/server/global"
)

const (
	UpdateMysql       = "mysql"
	UpdateDataFailed  = "\n[%v] --> %v 更新数据失败! [err]: %+v"
	UpdateDataSuccess = "\n[%v] --> %v 更新数据成功!"
)

const (
	UpdateOrderSystem = 10
	UpdateOrderDemo   = 1000
)

// TypedDbUpdateHandler 执行传入的 updater
type TypedDbUpdateHandler interface {
	UpdateConfig(ctx context.Context) error                  // 回写配置
	UpdateTables(ctx context.Context, inits initSlice) error // 建表
	UpdateData(ctx context.Context, inits initSlice) error   // 建数据
}

// Updater 提供 dao/*/init() 使用的接口，每个 updater 完成一个初始化过程
type Updater interface {
	UpdaterName() string // 需要初始化注册名，可以是表名，也可以是其他
	UpdateTable(ctx context.Context) (next context.Context, err error)
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
		panicStr := fmt.Sprintf("UpdaterName conflict on %s", name)
		// 此处不能panic，因为系统已经运行，此处记录日志，用于排查冲突，修复后再次升级
		global.OMS_LOG.Error(panicStr)
	}
	oup := orderedUpdater{order, up}
	updaters = append(updaters, &oup)
	updaterCache[name] = &oup
}

/* ------ * service * ------ */

type UpdateDBService interface {
	UpdateDB() error
	ClearUpdater()
}

type UpdateDBServiceImpl struct{}

// 已经升级，重启服务后，清除 updaters
func (s *UpdateDBServiceImpl) ClearUpdater() {
	updaters = updaterSlice{}
	updaterCache = map[string]*orderedUpdater{}
}

// 升级的前提是，部署了代码，部署代码一定会重新启动server，需要更新表结构，更新必要数据
func (s *UpdateDBServiceImpl) UpdateDB() (err error) {
	// ctx := context.Background()
	if len(updaters) == 0 {
		return errors.New("升级任务列表为空，请检查是否已执行完成")
	}

	return err
}
