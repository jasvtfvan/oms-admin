package system

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/service/initialize"
)

// 更新顺序
const updateOrderSysVersion = global.UpdateOrderSysVersion

type updateSysVersion struct{}

// UpdateData implements initialize.Updater.
func (u *updateSysVersion) UpdateData(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB

	version := global.OMS_CONFIG.Version
	result := db.Model(system.SysVersion{}).Where("version_name = ?", "oms_version").
		Updates(system.SysVersion{Version: version})
	if err = result.Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + u.UpdaterName() + " 表数据更新失败")
	}

	next = context.WithValue(ctx, u.UpdaterName(), result.RowsAffected)
	return next, err
}

// UpdateTable implements initialize.Updater.
func (u *updateSysVersion) UpdateTable(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB
	return ctx, db.AutoMigrate(&system.SysVersion{})
}

// UpdaterName implements initialize.Updater.
func (u *updateSysVersion) UpdaterName() string {
	return (&system.SysVersion{}).TableName()
}

// auto run
func init() {
	initialize.RegisterUpdate(updateOrderSysVersion, &updateSysVersion{})
}
