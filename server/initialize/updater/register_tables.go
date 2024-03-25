package updater

import (
	"context"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/migrate"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
	"go.uber.org/zap"
)

// 更新顺序
const updateOrderRegisterTables = global.UpdateOrderRegisterTables

type updateRegisterTables struct{}

// UpdateData implements initialize.Updater.
func (u *updateRegisterTables) UpdateData(ctx context.Context) (next context.Context, err error) {
	return ctx, err
}

// UpdateTable implements initialize.Updater.
func (u *updateRegisterTables) UpdateTable(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB
	tables := migrate.UpdateMigrateTables

	for _, tb := range tables {
		// AutoMigrate() 基本无需考虑错误
		err = db.AutoMigrate(&tb)
		if err != nil {
			global.OMS_LOG.Error("表结构更新失败", zap.Error(err))
		}
	}
	if err == nil {
		global.OMS_LOG.Info("表结构更新成功")
	}

	next = context.WithValue(ctx, u.UpdaterName(), len(tables))
	return next, db.AutoMigrate()
}

// UpdaterName implements initialize.Updater.
func (u *updateRegisterTables) UpdaterName() string {
	return "update_register_tables"
}

// auto run
func init() {
	initializeService.RegisterUpdate(updateOrderRegisterTables, &updateRegisterTables{})
}
