package demo

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/initializer"
	demoModel "github.com/jasvtfvan/oms-admin/server/model/demo"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
)

// 初始化顺序
const initOrderDemo = initializeService.InitOrderDemo + 1

type initDemo struct{}

// DataInserted implements initialize.Initializer.
func (i *initDemo) DataInserted(ctx context.Context) bool {
	return initializer.DataInserted(ctx, &demoModel.Demo{}, "name = ?", "demo")
}

// InitializeData implements initialize.Initializer.
func (i *initDemo) InitializeData(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB

	slices := []demoModel.Demo{
		{
			Name: "demo",
		},
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements initialize.Initializer.
func (i *initDemo) InitializerName() string {
	return (&demoModel.Demo{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initDemo) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &demoModel.Demo{})
}

// TableCreated implements initialize.Initializer.
func (i *initDemo) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &demoModel.Demo{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderDemo, &initDemo{})
}
