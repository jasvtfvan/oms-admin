package system

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/initialize/initializer/system/internal"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	systemService "github.com/jasvtfvan/oms-admin/server/service/system"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

const initOrderSysUser = initOrderSysRole + 1

type initSysUser struct{}

// DataInserted implements system.Initializer.
func (i *initSysUser) DataInserted(ctx context.Context) bool {
	return internal.DataInserted(ctx, &systemModel.SysUser{}, "username = ?", "admin")
}

// InitializeData implements system.Initializer.
func (i *initSysUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, systemService.ErrMissingDBContext
	}
	password := utils.BcryptHash("Oms123Admin456")
	slices := []systemModel.SysUser{
		{
			Username: "admin",
			Password: password,
			NickName: "超级管理员",
			Avatar:   "",
			Phone:    "",
			Email:    "",
			Enable:   true,
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements systemService.Initializer.
func (i *initSysUser) InitializerName() string {
	return (&systemModel.SysUser{}).TableName()
}

// MigrateTable implements system.Initializer.
func (i *initSysUser) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return internal.MigrateTable(ctx, &systemModel.SysUser{})
}

// TableCreated implements system.Initializer.
func (i *initSysUser) TableCreated(ctx context.Context) bool {
	return internal.TableCreated(ctx, &systemModel.SysUser{})
}

// auto run
func init() {
	systemService.RegisterInit(initOrderSysUser, &initSysUser{})
}
