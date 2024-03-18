package system

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/model/common"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	systemService "github.com/jasvtfvan/oms-admin/server/service/system"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"gorm.io/gorm"
)

const initOrderSysUser = systemService.InitOrderSystem + 1

type initSysUser struct{}

// auto run
func init() {
	systemService.RegisterInit(initOrderSysUser, &initSysUser{})
}

// InitializerName implements systemService.Initializer.
func (i *initSysUser) InitializerName() string {
	return systemModel.SysUser{}.TableName()
}

// TableCreated implements system.Initializer.
func (i *initSysUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		global.OMS_LOG.Fatal(systemService.ErrMissingDBContext.Error())
		return false
	}
	return db.Migrator().HasTable(&systemModel.SysUser{})
}

// MigrateTable implements system.Initializer.
func (i *initSysUser) MigrateTable(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, systemService.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&systemModel.SysUser{})
}

// DataInserted implements system.Initializer.
func (i *initSysUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		global.OMS_LOG.Fatal(systemService.ErrMissingDBContext.Error())
		return false
	}
	if errors.Is(db.Where("Username = ?", "admin").First(&systemModel.SysUser{}).
		Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// InitializeData implements system.Initializer.
func (i *initSysUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, systemService.ErrMissingDBContext
	}
	password := utils.BcryptHash("Oms123Admin456")
	snowflakeWorker := utils.NewSnowflakeWorker(0)
	ID := snowflakeWorker.NextId()

	slices := []systemModel.SysUser{
		{
			BaseModel: common.BaseModel{
				ID: uint(ID),
			},
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
