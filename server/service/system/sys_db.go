package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"

	"github.com/jasvtfvan/oms-admin/server/global"
	"gorm.io/gorm"
)

const (
	Mysql           = "mysql"
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	InitDataExist   = "\n[%v] --> %v 的初始数据已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 初始数据成功!\n"
)

const (
	InitOrderSystem   = 10
	InitOrderInternal = 1000
	InitOrderExternal = 100000
)

var (
	ErrMissingDBContext        = errors.New("Missing db in context")
	ErrMissingDependentContext = errors.New("Missing dependent value in context")
	ErrDBTypeMismatch          = errors.New("Db type mismatch")
)

// TypedDbInitHandler 执行传入的 initializer
type TypedDbInitHandler interface {
	EnsureDB(ctx context.Context) (context.Context, error) // 建数据库，失败属于 fatal error，让它 panic
	WriteConfig(ctx context.Context) error                 // 回写配置
	InitTables(ctx context.Context, inits initSlice) error // 建表
	InitData(ctx context.Context, inits initSlice) error   // 建数据
}

// Initializer 提供 dao/*/init() 使用的接口，每个 initializer 完成一个初始化过程
type Initializer interface {
	InitializerName() string // 需要初始化注册名，可以是表名，也可以是其他
	TableCreated(ctx context.Context) bool
	MigrateTable(ctx context.Context) (next context.Context, err error)
	InitializeData(ctx context.Context) (next context.Context, err error)
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
	cache        map[string]*orderedInitializer
)

// RegisterInit 注册要执行的初始化过程，InitDB() 时根据注册的 initializer 进行初始化
func RegisterInit(order int, i Initializer) {
	if initializers == nil {
		initializers = initSlice{}
	}
	if cache == nil {
		cache = map[string]*orderedInitializer{}
	}
	name := i.InitializerName()
	if _, existed := cache[name]; existed {
		panicStr := fmt.Sprintf("InitializerName conflict on %s", name)
		fmt.Println(panicStr)
		panic(panicStr)
	}
	oi := orderedInitializer{order, i}
	initializers = append(initializers, &oi)
	cache[name] = &oi
}

/* ------ * service * ------ */

type DbService interface {
	InitDB() error
	CheckDB() error
}

type DbServiceImpl struct{}

// 检查数据连接
func (*DbServiceImpl) CheckDB() (err error) {
	if global.OMS_DB == nil {
		return errors.New("DB为空，数据库未创建")
	} else {
		db, _ := global.OMS_DB.DB()
		if err = db.Ping(); err != nil {
			return err
		}
	}
	return err
}

// 初始化数据
func (dbService *DbServiceImpl) InitDB() (err error) {
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
		ctx = context.WithValue(ctx, "dbType", "mysql")
	default:
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbType", "mysql")
	}

	ctx, err = initHandler.EnsureDB(ctx)
	if err != nil {
		return err
	}
	db := ctx.Value("db").(*gorm.DB)
	global.OMS_DB = db

	if err = initHandler.WriteConfig(ctx); err != nil {
		return err
	}
	if err = initHandler.InitTables(ctx, initializers); err != nil {
		return err
	}
	if err = initHandler.InitData(ctx, initializers); err != nil {
		return err
	}

	initializers = initSlice{}
	cache = map[string]*orderedInitializer{}

	return err
}

// createDatabase 创建数据库（ EnsureDB() 中调用 ）
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			// 数据库链接没关闭成功，会导致内存泄漏，需要排查
			global.OMS_LOG.Fatal(err.Error())
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
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
