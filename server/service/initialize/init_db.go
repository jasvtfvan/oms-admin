package initialize

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/jasvtfvan/oms-admin/server/global"
)

const (
	InitMysql       = "mysql"
	InitDataExist   = "\n[%v] --> %v 的初始数据已存在!"
	InitDataFailed  = "\n[%v] --> %v 初始数据失败! [err]: %+v"
	InitDataSuccess = "\n[%v] --> %v 初始数据成功!"
)

const (
	InitOrderSystem = 10
	InitOrderDemo   = 1000
)

// TypedDbInitHandler 执行传入的 initializer
type TypedDbInitHandler interface {
	WriteConfig(ctx context.Context) error                 // 回写配置
	InitTables(ctx context.Context, inits initSlice) error // 建表
	InitData(ctx context.Context, inits initSlice) error   // 建数据
}

// Initializer 提供 initializer/*/init() 使用的接口，每个 initializer 完成一个初始化过程
type Initializer interface {
	InitializerName() string // 需要初始化注册名，可以是表名，也可以是其他
	TableCreated(ctx context.Context) bool
	MigrateTable(ctx context.Context) (next context.Context, err error)   // ！注意：这里是create表结构
	InitializeData(ctx context.Context) (next context.Context, err error) // ！注意：这里是insert数据
	DataInserted(ctx context.Context) bool
}

// orderedInitializer 组合一个顺序字段，用来排序
type orderedInitializer struct {
	order int
	Initializer
}

// initSlice 供 initializer 排序依赖时使用
type initSlice []*orderedInitializer

var (
	initializers initSlice
	initCache    map[string]*orderedInitializer
)

// RegisterInit 注册要执行的初始化过程，InitDB() 时根据注册的 initializer 进行初始化
func RegisterInit(order int, i Initializer) {
	if initializers == nil {
		initializers = initSlice{}
	}
	if initCache == nil {
		initCache = map[string]*orderedInitializer{}
	}
	name := i.InitializerName()
	if _, existed := initCache[name]; existed {
		panicStr := fmt.Sprintf("InitializerName conflict on %s", name)
		// 表名冲突，写入fatal日志，因为代码错误，会导致程序不能正确运行
		global.OMS_LOG.Fatal(panicStr)
	}
	oi := orderedInitializer{order, i}
	initializers = append(initializers, &oi)
	initCache[name] = &oi
}

/* ------ * service * ------ */

type InitDBService interface {
	InitDB() error
	ClearInitializer()
}

type InitDBServiceImpl struct{}

// 已经初始化，重启服务后，清除 initializers
func (s *InitDBServiceImpl) ClearInitializer() {
	initializers = initSlice{}
	initCache = map[string]*orderedInitializer{}
}

// 初始化数据
func (s *InitDBServiceImpl) InitDB() (err error) {
	ctx := context.Background()
	if len(initializers) == 0 {
		return errors.New("初始化任务列表为空，请检查是否已执行完成")
	}
	/*
		保证有依赖的 initializer 排在后面执行
		Note: 若 initializer 只有单一依赖，可以写为 B=A+1, C=A+1; 由于 BC 之间没有依赖关系，所以谁先谁后并不影响初始化
		若存在多个依赖，可以写为 C=A+B, D=A+B+C, E=A+1;
		C必然>A|B，因此在AB之后执行，D必然>A|B|C，因此在ABC后执行，而E只依赖A，顺序与CD无关，因此E与CD哪个先执行并不影响
	*/
	sort.Sort(&initializers)

	var initHandler TypedDbInitHandler
	switch global.OMS_CONFIG.System.DbType {
	case "mysql":
		initHandler = NewMysqlInitHandler()
	default:
		initHandler = NewMysqlInitHandler()
	}

	if err = initHandler.WriteConfig(ctx); err != nil {
		return err
	}
	global.OMS_LOG.Info("更新配置文件成功")
	if err = initHandler.InitTables(ctx, initializers); err != nil {
		return err
	}
	global.OMS_LOG.Info("初始化表成功")
	if err = initHandler.InitData(ctx, initializers); err != nil {
		return err
	}
	global.OMS_LOG.Info("初始化数据成功")

	initializers = initSlice{}
	initCache = map[string]*orderedInitializer{}

	return err
}

// createTables 创建表（默认 dbInitHandler.initTables 行为）
func createTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, in := range inits {
		if in.TableCreated(next) {
			continue
		}
		if n, err := in.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}

/* -- sortable interface -- */

func (a initSlice) Len() int {
	return len(a)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
