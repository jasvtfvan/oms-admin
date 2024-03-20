package system

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/initialize/initializer/system/internal"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	systemService "github.com/jasvtfvan/oms-admin/server/service/system"
	"gorm.io/gorm"
)

const initOrderSysGroup = systemService.InitOrderSystem + 1

type initSysGroup struct{}

// DataInserted implements system.Initializer.
func (i *initSysGroup) DataInserted(ctx context.Context) bool {
	return internal.DataInserted(ctx, &systemModel.SysGroup{}, "org_code = ?", "root")
}

// InitializeData implements system.Initializer.
func (i *initSysGroup) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, systemService.ErrMissingDBContext
	}
	slices := []systemModel.SysGroup{
		{
			ShortName: "根组织",
			OrgCode:   "root",
			ParentID:  0,
			Sort:      0,
			Enable:    true,
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements system.Initializer.
func (i *initSysGroup) InitializerName() string {
	return (&systemModel.SysGroup{}).TableName()
}

// MigrateTable implements system.Initializer.
func (i *initSysGroup) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return internal.MigrateTable(ctx, &systemModel.SysGroup{})
}

// TableCreated implements system.Initializer.
func (i *initSysGroup) TableCreated(ctx context.Context) bool {
	return internal.TableCreated(ctx, &systemModel.SysGroup{})
}

// auto run
func init() {
	systemService.RegisterInit(initOrderSysGroup, &initSysGroup{})
}
