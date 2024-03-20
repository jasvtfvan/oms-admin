package system

import (
	"context"
	"errors"
	"strconv"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/initializer/system/internal"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	systemService "github.com/jasvtfvan/oms-admin/server/service/system"
	"gorm.io/gorm"
)

// 初始化顺序
const initOrderSysUserRole = initOrderSysUserGroup + 1

type initSysUserRole struct{}

// DataInserted implements system.Initializer.
func (i *initSysUserRole) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		global.OMS_LOG.Error(systemService.ErrMissingDBContext.Error())
		return false
	}

	// user表已经初始化完成
	sysUser := &systemModel.SysUser{}
	db.Where("username = ?", "admin").First(sysUser)
	// role表已经初始化完成
	sysRole := &systemModel.SysRole{}
	db.Where("role_code = ?", "admin").First(sysRole)

	return internal.DataInserted(
		ctx, &systemModel.SysUserRole{},
		"sys_user_id = ? and sys_role_id = ?",
		strconv.Itoa(int(sysUser.ID)),
		strconv.Itoa(int(sysRole.ID)),
	)
}

// InitializeData implements system.Initializer.
func (i *initSysUserRole) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, systemService.ErrMissingDBContext
	}

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

// InitializerName implements system.Initializer.
func (i *initSysUserRole) InitializerName() string {
	return (&systemModel.SysUserRole{}).TableName()
}

// MigrateTable implements system.Initializer.
func (i *initSysUserRole) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return internal.MigrateTable(ctx, &systemModel.SysUserRole{})
}

// TableCreated implements system.Initializer.
func (i *initSysUserRole) TableCreated(ctx context.Context) bool {
	return internal.TableCreated(ctx, &systemModel.SysUserRole{})
}

// auto run
func init() {
	// systemService.RegisterInit(initOrderSysUserRole, &initSysUserRole{})
}
