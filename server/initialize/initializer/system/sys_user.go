package system

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/initializer"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

// 初始化顺序
const initOrderSysUser = global.InitOrderSysUser

type initSysUser struct{}

// DataInserted implements initialize.Initializer.
func (i *initSysUser) DataInserted(ctx context.Context) bool {
	rootUsername := global.OMS_CONFIG.System.Username
	return initializer.DataInserted(ctx, &systemModel.SysUser{}, "username = ?", rootUsername)
}

// InitializeData implements initialize.Initializer.
func (i *initSysUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	rootUsername := global.OMS_CONFIG.System.Username
	rootPassword := global.OMS_CONFIG.System.Password
	db := global.OMS_DB

	password := utils.BcryptHash(rootPassword)
	slices := []systemModel.SysUser{
		{
			Username:     rootUsername,
			Password:     password,
			NickName:     "超级管理员",
			Avatar:       "",
			Phone:        "",
			Email:        "",
			IsAdmin:      true,
			Enable:       true,
			LogOperation: true,
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements initializeService.Initializer.
func (i *initSysUser) InitializerName() string {
	return (&systemModel.SysUser{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initSysUser) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &systemModel.SysUser{})
}

// TableCreated implements initialize.Initializer.
func (i *initSysUser) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &systemModel.SysUser{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderSysUser, &initSysUser{}, &systemModel.SysUser{})
}
