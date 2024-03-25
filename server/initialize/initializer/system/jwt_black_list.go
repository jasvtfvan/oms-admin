package system

import (
	"context"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/initializer"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
)

// 初始化顺序
const initOrderJWTBlackList = global.InitOrderJWTBlackList

type initJWTBlackList struct{}

// DataInserted implements initialize.Initializer.
func (i *initJWTBlackList) DataInserted(ctx context.Context) bool {
	return true
}

// InitializeData implements initialize.Initializer.
func (i *initJWTBlackList) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, err
}

// InitializerName implements initialize.Initializer.
func (i *initJWTBlackList) InitializerName() string {
	return (&systemModel.JWTBlackList{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initJWTBlackList) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &systemModel.JWTBlackList{})
}

// TableCreated implements initialize.Initializer.
func (i *initJWTBlackList) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &systemModel.JWTBlackList{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderJWTBlackList, &initJWTBlackList{}, &systemModel.JWTBlackList{})
}
