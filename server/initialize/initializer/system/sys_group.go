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
const initOrderSysGroup = initOrderSysVersion + 1

type initSysGroup struct{}

// DataInserted implements initialize.Initializer.
func (i *initSysGroup) DataInserted(ctx context.Context) bool {
	return initializer.DataInserted(ctx, &systemModel.SysGroup{}, "org_code = ?", "root")
}

// InitializeData implements initialize.Initializer.
func (i *initSysGroup) InitializeData(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB
	slices := []systemModel.SysGroup{
		{
			ShortName: "根组织",
			OrgCode:   "root",
			ParentID:  0,
			Sort:      0,
			Enable:    true,
			// SysRoles: []systemModel.SysRole{
			// 	{
			// 		RoleName: "超级管理员",
			// 		RoleCode: "admin",
			// 		Sort:     0,
			// 		Comment:  "超级管理员",
			// 		Enable:   true,
			// 	},
			// },
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements initialize.Initializer.
func (i *initSysGroup) InitializerName() string {
	return (&systemModel.SysGroup{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initSysGroup) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &systemModel.SysGroup{})
}

// TableCreated implements initialize.Initializer.
func (i *initSysGroup) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &systemModel.SysGroup{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderSysGroup, &initSysGroup{})
}
