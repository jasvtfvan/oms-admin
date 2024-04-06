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
const initOrderSysRole = global.InitOrderSysRole

type initSysRole struct{}

// DataInserted implements initialize.Initializer.
func (i *initSysRole) DataInserted(ctx context.Context) bool {
	rootRoleCode := initializer.GetRootRoleCode()
	return initializer.DataInserted(ctx, &systemModel.SysRole{}, "role_code = ?", rootRoleCode)
}

// InitializeData implements initialize.Initializer.
func (i *initSysRole) InitializeData(ctx context.Context) (next context.Context, err error) {
	rootRoleCode := initializer.GetRootRoleCode()
	rootOrgCode := initializer.GetRootGroupCode()
	db := global.OMS_DB

	// group表已经初始化完成
	sysGroup := &systemModel.SysGroup{}
	db.Where("org_code = ?", rootOrgCode).First(sysGroup)

	slices := []systemModel.SysRole{
		{
			RoleName:   "超级管理员",
			RoleCode:   rootRoleCode,
			Sort:       0,
			Comment:    "超级管理员",
			Enable:     true,
			SysGroupID: sysGroup.ID,
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements initialize.Initializer.
func (i *initSysRole) InitializerName() string {
	return (&systemModel.SysRole{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initSysRole) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &systemModel.SysRole{})
}

// TableCreated implements initialize.Initializer.
func (i *initSysRole) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &systemModel.SysRole{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderSysRole, &initSysRole{}, &systemModel.SysRole{})
}
