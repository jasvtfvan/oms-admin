package system

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/gookit/color"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MysqlInitHandler struct{}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

// EnsureDB implements system.TypedDbInitHandler.
func (h *MysqlInitHandler) EnsureDB(ctx context.Context) (next context.Context, err error) {
	if s, ok := ctx.Value("dbType").(string); !ok || s != "mysql" {
		return ctx, ErrDBTypeMismatch
	}
	config := global.OMS_CONFIG.Mysql
	if config.DbName == "" {
		return ctx, errors.New("数据库名不能为空")
	}
	dsn := config.EmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", config.DbName)
	if err = createDatabase(dsn, "mysql", createSql); err != nil {
		return ctx, err
	}

	var db *gorm.DB
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Prefix,
			SingularTable: config.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Dsn(), // DSN data source name
		DefaultStringSize:         191,          // string 类型字段的默认长度
		SkipInitializeWithVersion: true,         // 根据版本自动配置
	}), gormConfig); err != nil {
		return ctx, err
	}

	db.InstanceSet("gorm:table_options", "ENGINE="+config.Engine)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	next = context.WithValue(ctx, "db", db)

	return next, err
}

// WriteConfig implements TypedDbInitHandler.
func (h *MysqlInitHandler) WriteConfig(ctx context.Context) error {
	global.OMS_CONFIG.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(global.OMS_CONFIG)
	for k, v := range cs {
		global.OMS_VP.Set(k, v)
	}
	return global.OMS_VP.WriteConfig()
}

// InitTables implements system.TypedDbInitHandler.
func (h *MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

// InitData implements system.TypedDbInitHandler.
func (h *MysqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, in := range inits {
		if in.DataInserted(next) {
			// 已经插入过，写入warn日志
			global.OMS_LOG.Warn(fmt.Sprintf(InitDataExist, Mysql, in.InitializerName()))
			color.Info.Printf(InitDataExist, Mysql, in.InitializerName())
			continue
		}
		if n, err := in.InitializeData(next); err != nil {
			// 数据初始化失败，写入fatal日志，因为系统数据不全，会导致程序不能正确运行
			global.OMS_LOG.Fatal(fmt.Sprintf(InitDataFailed, Mysql, in.InitializerName(), err.Error()))
			color.Info.Printf(InitDataFailed, Mysql, in.InitializerName(), err.Error())
			return err
		} else {
			// 数据初始化成功，写入info日志
			global.OMS_LOG.Info(fmt.Sprintf(InitDataSuccess, Mysql, in.InitializerName()))
			color.Info.Printf(InitDataSuccess, Mysql, in.InitializerName())
			next = n
		}
	}
	return nil
}
