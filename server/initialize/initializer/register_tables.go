package initializer

import (
	"context"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/migrate"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
	"go.uber.org/zap"
)

// 初始化顺序
const initOrderRegisterTables = global.InitOrderRegisterTables

type initRegisterTables struct{}

// DataInserted implements initialize.Initializer.
func (i *initRegisterTables) DataInserted(ctx context.Context) bool {
	return true
}

// InitializeData implements initialize.Initializer.
func (i *initRegisterTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, err
}

// InitializerName implements initialize.Initializer.
func (i *initRegisterTables) InitializerName() string {
	return "init_register_tables"
}

// MigrateTable implements initialize.Initializer.
func (i *initRegisterTables) MigrateTable(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB
	tables := migrate.InitMigrateTables

	for _, tb := range tables {
		// AutoMigrate() 基本无需考虑错误
		err = db.AutoMigrate(&tb)
		if err != nil {
			global.OMS_LOG.Error("表结构初始化失败", zap.Error(err))
		}
	}
	if err == nil {
		global.OMS_LOG.Info("表结构初始化成功")
	}

	next = context.WithValue(ctx, i.InitializerName(), len(tables))
	return next, err
}

// TableCreated implements initialize.Initializer.
func (i *initRegisterTables) TableCreated(ctx context.Context) bool {
	db := global.OMS_DB
	tables := migrate.InitMigrateTables
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}

// auto run
func init() {
	tables := migrate.InitMigrateTables
	initializeService.RegisterInit(initOrderRegisterTables, &initRegisterTables{}, tables...)
}
