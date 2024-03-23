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
const initOrderSysUserGroup = initOrderSysUser + 1

type initSysUserGroup struct{}

// DataInserted implements initialize.Initializer.
func (i *initSysUserGroup) DataInserted(ctx context.Context) bool {
	db := global.OMS_DB

	// user表已经初始化完成
	sysUser := &systemModel.SysUser{}
	db.Where("username = ?", "admin").First(sysUser)
	// group表已经初始化完成
	sysGroup := &systemModel.SysGroup{}
	db.Where("org_code = ?", "root").First(sysGroup)

	return initializer.DataInserted(
		ctx, &systemModel.SysUserGroup{},
		"sys_user_id = ? and sys_group_id = ?",
		strconv.Itoa(int(sysUser.ID)),
		strconv.Itoa(int(sysGroup.ID)),
	)
}

// InitializeData implements initialize.Initializer.
func (i *initSysUserGroup) InitializeData(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB

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

// InitializerName implements initialize.Initializer.
func (i *initSysUserGroup) InitializerName() string {
	return (&systemModel.SysUserGroup{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initSysUserGroup) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &systemModel.SysUserGroup{})
}

// TableCreated implements initialize.Initializer.
func (i *initSysUserGroup) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &systemModel.SysUserGroup{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderSysUserGroup, &initSysUserGroup{}, &systemModel.SysUserGroup{})
}
