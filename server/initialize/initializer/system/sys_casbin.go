package system

import (
	"context"
	"errors"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/initialize/initializer"
	systemModel "github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/jasvtfvan/oms-admin/server/model/system/response"
	initializeService "github.com/jasvtfvan/oms-admin/server/service/initialize"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

// 初始化顺序
const initOrderSysCasbin = global.InitOrderSysCasbin

type initSysCasbin struct{}

// DataInserted implements initialize.Initializer.
func (i *initSysCasbin) DataInserted(ctx context.Context) bool {
	rootRoleCode := initializer.GetRootRoleCode()
	rootOrgCode := initializer.GetRootGroupCode()
	casbinInfos := response.DefaultCasbinSource()
	if len(casbinInfos) <= 0 {
		return true
	}
	// 通过最后一条数据判断，是否完全插入
	lastCasbin := casbinInfos[len(casbinInfos)-1]
	return initializer.DataInserted(ctx, &systemModel.SysCasbin{},
		"ptype = 'p' and v0 = ? and v1 = ? and v2 = ? and v3 = ?",
		rootRoleCode, rootOrgCode, lastCasbin.Path, lastCasbin.Method)
}

// InitializeData implements initialize.Initializer.
func (i *initSysCasbin) InitializeData(ctx context.Context) (next context.Context, err error) {
	db := global.OMS_DB

	rootRoleCode := initializer.GetRootRoleCode()
	rootOrgCode := initializer.GetRootGroupCode()
	slices := []systemModel.SysCasbin{}
	casbinInfos := response.DefaultCasbinSource()

	var sysCasbinWorkerId int64 = global.SysCasbinWorkerId
	snowflakeWorker := utils.NewSnowflakeWorker(sysCasbinWorkerId)

	for _, v := range casbinInfos {
		slices = append(slices, systemModel.SysCasbin{
			ID:    uint(snowflakeWorker.NextId()),
			Ptype: "p",
			V0:    rootRoleCode,
			V1:    rootOrgCode,
			V2:    v.Path,
			V3:    v.Method,
		})
	}
	if err = db.Create(&slices).Error; err != nil {
		return ctx, errors.New(err.Error() + ": " + i.InitializerName() + " 表数据初始化失败")
	}
	next = context.WithValue(ctx, i.InitializerName(), slices)
	return next, err
}

// InitializerName implements initialize.Initializer.
func (i *initSysCasbin) InitializerName() string {
	return (&systemModel.SysCasbin{}).TableName()
}

// MigrateTable implements initialize.Initializer.
func (i *initSysCasbin) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return initializer.MigrateTable(ctx, &systemModel.SysCasbin{})
}

// TableCreated implements initialize.Initializer.
func (i *initSysCasbin) TableCreated(ctx context.Context) bool {
	return initializer.TableCreated(ctx, &systemModel.SysCasbin{})
}

// auto run
func init() {
	initializeService.RegisterInit(initOrderSysCasbin, &initSysCasbin{}, &systemModel.SysCasbin{})
}
