package system

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/initialize/initializer/system/internal"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	systemService "github.com/jasvtfvan/oms-admin/server/service/system"
	"gorm.io/gorm"
)

const initOrderSysRole = initOrderSysGroup + 1

type initSysRole struct{}

// DataInserted implements system.Initializer.
func (i *initSysRole) DataInserted(ctx context.Context) bool {
	return internal.DataInserted(ctx, &systemModel.SysRole{}, "role_code = ?", "admin")
}

// InitializeData implements system.Initializer.
func (i *initSysRole) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, systemService.ErrMissingDBContext
	}
	slices := []systemModel.SysRole{
		{
			RoleName: "超级管理员",
			RoleCode: "admin",
			Sort:     0,
			Comment:  "超级管理员",
			Enable:   true,
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements system.Initializer.
func (i *initSysRole) InitializerName() string {
	return (&systemModel.SysRole{}).TableName()
}

// MigrateTable implements system.Initializer.
func (i *initSysRole) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return internal.MigrateTable(ctx, &systemModel.SysRole{})
}

// TableCreated implements system.Initializer.
func (i *initSysRole) TableCreated(ctx context.Context) bool {
	return internal.TableCreated(ctx, &systemModel.SysRole{})
}

// auto run
func init() {
	systemService.RegisterInit(initOrderSysUser, &initSysRole{})
}
