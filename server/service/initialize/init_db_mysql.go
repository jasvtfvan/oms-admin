package initialize

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

type MysqlInitHandler struct{}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

// WriteConfig implements TypedDbInitHandler.
func (h *MysqlInitHandler) WriteConfig(ctx context.Context) error {
	global.OMS_CONFIG.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(global.OMS_CONFIG)
	for k, v := range cs {
		global.OMS_VP.Set(k, v)
	}
	return global.OMS_VP.WriteConfig()
}

// InitTables implements system.TypedDbInitHandler.
func (h *MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, in := range inits {
		if in.TableCreated(next) {
			continue
		}
		if n, err := in.MigrateTable(next); err != nil {
			global.OMS_LOG.Error(fmt.Sprintf(InitTableFailed, InitMysql, in.InitializerName(), err.Error()))
			return err
		} else {
			next = n
		}
	}
	return nil
}

// InitData implements system.TypedDbInitHandler.
func (h *MysqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, in := range inits {
		if in.DataInserted(next) {
			// 已经插入过，写入warn日志
			global.OMS_LOG.Warn(fmt.Sprintf(InitDataExist, InitMysql, in.InitializerName()))
			continue
		}
		if n, err := in.InitializeData(next); err != nil {
			global.OMS_LOG.Error(fmt.Sprintf(InitDataFailed, InitMysql, in.InitializerName(), err.Error()))
			return err
		} else {
			// 数据初始化成功，写入info日志
			global.OMS_LOG.Info(fmt.Sprintf(InitDataSuccess, InitMysql, in.InitializerName()))
			next = n
		}
	}
	return nil
}
