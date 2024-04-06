package initialize

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/utils/freecache"
	"gorm.io/gorm"
)

const (
	InitMysql       = "mysql"
	InitDataExist   = "\n[%v] --> %v 的初始数据已存在!"
	InitTableFailed = "\n[%v] --> %v 创建表失败! [err]: %+v"
	InitDataFailed  = "\n[%v] --> %v 初始数据失败! [err]: %+v"
	InitDataSuccess = "\n[%v] --> %v 初始数据成功!"
)

// TypedDbInitHandler 执行传入的 initializer
type TypedDbInitHandler interface {
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
type tableSlice []interface{}

var (
	initializers initSlice
	initCache    map[string]*orderedInitializer
	initTables   tableSlice
)

// RegisterInit 注册要执行的初始化过程，InitDB() 时根据注册的 initializer 进行初始化
func RegisterInit(order int, i Initializer, model ...interface{}) {
	if initTables == nil {
		initTables = tableSlice{}
	}
	initTables = append(initTables, model...)

	if initializers == nil {
		initializers = initSlice{}
	}
	if initCache == nil {
		initCache = map[string]*orderedInitializer{}
	}
	name := i.InitializerName()
	if _, existed := initCache[name]; existed {
		// 表名冲突，会导致程序不能正确运行，init函数在main函数前运行，出错后禁止启动
		panicStr := fmt.Sprintf("InitializerName conflict on %s", name)
		panic(panicStr)
	}
	oi := orderedInitializer{order, i}
	initializers = append(initializers, &oi)
	initCache[name] = &oi
}

/* ------ * service * ------ */

type InitDBService interface {
	CheckDB() error
	ClearInitializer()
	InitDB() error
}

type InitDBServiceImpl struct{}

// 检查数据连接
func (*InitDBServiceImpl) CheckDB() error {
	var readyStruct freecache.Bool
	isReady := cacheStore.Get("DBReady", readyStruct)
	if cacheStore.Verify("DBReady", true, isReady) {
		fmt.Println("[Golang] CheckDB from local cache")
		return nil
	}

	if global.OMS_DB == nil {
		return errors.New("DB为空，数据库未创建")
	} else {
		db := global.OMS_DB
		tableCreated := db.Migrator().HasTable(&system.SysVersion{})
		if !tableCreated {
			return errors.New("表结构尚未创建")
		}
		err := db.Where("version_code = ?", "oms_version").First(&system.SysVersion{}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("表数据尚未插入")
		}
	}

	cacheStore.Set("DBReady", true)
	return nil
}

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

	if err = initHandler.InitTables(ctx, initializers); err != nil {
		global.OMS_LOG.Error("初始化表结构失败，删除所有表:" + err.Error())
		if dErr := deleteTables(); dErr != nil {
			global.OMS_LOG.Error("表删除失败（初始化失败，需要删表重启，重新初始化），请手动删表:" + dErr.Error())
		}
		return err
	}
	global.OMS_LOG.Info("初始化表成功")
	if err = initHandler.InitData(ctx, initializers); err != nil {
		global.OMS_LOG.Error("初始化表数据失败，删除所有表:" + err.Error())
		if dErr := deleteTables(); dErr != nil {
			global.OMS_LOG.Error("表删除失败（初始化失败，需要删表重启，重新初始化），请手动删表:" + dErr.Error())
		}
		return err
	}
	global.OMS_LOG.Info("初始化数据成功")

	initializers = initSlice{}
	initCache = map[string]*orderedInitializer{}
	initTables = tableSlice{}

	return err
}

// 删除表，初始化过程失败。需要调整代码重新初始化，要删除所有表。
func deleteTables() (err error) {
	db := global.OMS_DB
	for _, table := range initTables {
		if err := db.Migrator().DropTable(table); err != nil {
			return err
		}
	}
	initTables = tableSlice{}
	return err
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
