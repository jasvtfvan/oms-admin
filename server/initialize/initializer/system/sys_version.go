package system

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/initializer"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
)

// 初始化顺序
const initOrderSysVersion = global.InitOrderSysVersion

type initSysVersion struct{}

// DataInserted implements initialize.Initializer.
func (i *initSysVersion) DataInserted(ctx context.Context) bool {
	return initializer.DataInserted(ctx, &systemModel.SysVersion{}, "version_name = ?", "oms_version")
}

// InitializeData implements initialize.Initializer.
func (i *initSysVersion) InitializeData(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB

	version := global.OMS_CONFIG.Version
	slices := []systemModel.SysVersion{
		{
			VersionName: "oms_version",
			Version:     version,
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements initialize.Initializer.
func (i *initSysVersion) InitializerName() string {
	return (&systemModel.SysVersion{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initSysVersion) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &systemModel.SysVersion{})
}

// TableCreated implements initialize.Initializer.
func (i *initSysVersion) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &systemModel.SysVersion{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderSysVersion, &initSysVersion{}, &systemModel.SysVersion{})
}
