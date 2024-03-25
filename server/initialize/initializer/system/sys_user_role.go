package system

import (
	"context"
	"errors"
	"strconv"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/initializer"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
)

// 初始化顺序
const initOrderSysUserRole = global.InitOrderSysUserRole

type initSysUserRole struct{}

// DataInserted implements initialize.Initializer.
func (i *initSysUserRole) DataInserted(ctx context.Context) bool {
	db := global.OMS_DB

	// user表已经初始化完成
	sysUser := &systemModel.SysUser{}
	db.Where("username = ?", "admin").First(sysUser)
	// role表已经初始化完成
	sysRole := &systemModel.SysRole{}
	db.Where("role_code = ?", "admin").First(sysRole)

	return initializer.DataInserted(
		ctx, &systemModel.SysUserRole{},
		"sys_user_id = ? and sys_role_id = ?",
		strconv.Itoa(int(sysUser.ID)),
		strconv.Itoa(int(sysRole.ID)),
	)
}

// InitializeData implements initialize.Initializer.
func (i *initSysUserRole) InitializeData(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB

	// user表已经初始化完成
	sysUser := &systemModel.SysUser{}
	db.Where("username = ?", "admin").First(sysUser)
	// role表已经初始化完成
	sysRole := &systemModel.SysRole{}
	db.Where("role_code = ?", "admin").First(sysRole)

	slices := []systemModel.SysUserRole{
		{
			SysUserID: sysUser.ID,
			SysRoleID: sysRole.ID,
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements initialize.Initializer.
func (i *initSysUserRole) InitializerName() string {
	return (&systemModel.SysUserRole{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initSysUserRole) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &systemModel.SysUserRole{})
}

// TableCreated implements initialize.Initializer.
func (i *initSysUserRole) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &systemModel.SysUserRole{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderSysUserRole, &initSysUserRole{}, &systemModel.SysUserRole{})
}
