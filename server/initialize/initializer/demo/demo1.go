package demo

import (
	"context"

	"github.com/jasvtfvan/oms-admin/server/initialize/initializer"
	demoModel "github.com/jasvtfvan/oms-admin/server/model/demo"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
)

// 初始化顺序
const initOrderDemo1 = initOrderDemo + 1

type initDemo1 struct{}

// DataInserted implements initialize.Initializer.
func (i *initDemo1) DataInserted(ctx context.Context) bool {
	return true // 只初始化表结构
}

// InitializeData implements initialize.Initializer.
func (i *initDemo1) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, err // 只初始化表结构
}

// InitializerName implements initialize.Initializer.
func (i *initDemo1) InitializerName() string {
	return (&demoModel.Demo1{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initDemo1) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &demoModel.Demo1{})
}

// TableCreated implements initialize.Initializer.
func (i *initDemo1) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &demoModel.Demo1{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderDemo1, &initDemo1{})
}
