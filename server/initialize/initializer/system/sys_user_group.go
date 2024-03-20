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
const initOrderSysUserGroup = initOrderSysUser + 1

type initSysUserGroup struct{}

// DataInserted implements system.Initializer.
func (i *initSysUserGroup) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		global.OMS_LOG.Error(systemService.ErrMissingDBContext.Error())
		return false
	}

	// user表已经初始化完成
	sysUser := &systemModel.SysUser{}
	db.Where("username = ?", "admin").First(sysUser)
	// group表已经初始化完成
	sysGroup := &systemModel.SysGroup{}
	db.Where("org_code = ?", "root").First(sysGroup)

	return internal.DataInserted(
		ctx, &systemModel.SysUserGroup{},
		"sys_user_id = ? and sys_group_id = ?",
		strconv.Itoa(int(sysUser.ID)),
		strconv.Itoa(int(sysGroup.ID)),
	)
}

// InitializeData implements system.Initializer.
func (i *initSysUserGroup) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, systemService.ErrMissingDBContext
	}

	// user表已经初始化完成
	sysUser := &systemModel.SysUser{}
	db.Where("username = ?", "admin").First(sysUser)
	// group表已经初始化完成
	sysGroup := &systemModel.SysGroup{}
	db.Where("org_code = ?", "root").First(sysGroup)

	slices := []systemModel.SysUserGroup{
		{
			SysUserID:  sysUser.ID,
			SysGroupID: sysGroup.ID,
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements system.Initializer.
func (i *initSysUserGroup) InitializerName() string {
	return (&systemModel.SysUserGroup{}).TableName()
}

// MigrateTable implements system.Initializer.
func (i *initSysUserGroup) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return internal.MigrateTable(ctx, &systemModel.SysUserGroup{})
}

// TableCreated implements system.Initializer.
func (i *initSysUserGroup) TableCreated(ctx context.Context) bool {
	return internal.TableCreated(ctx, &systemModel.SysUserGroup{})
}

// auto run
func init() {
	// systemService.RegisterInit(initOrderSysUserGroup, &initSysUserGroup{})
}
