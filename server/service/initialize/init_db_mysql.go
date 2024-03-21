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
	return createTables(ctx, inits)
}

// InitData implements system.TypedDbInitHandler.
func (h *MysqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, in := range inits {
		if in.DataInserted(next) {
			// 已经插入过，写入warn日志
			global.OMS_LOG.Warn(fmt.Sprintf(InitDataExist, Mysql, in.InitializerName()))
			continue
		}
		if n, err := in.InitializeData(next); err != nil {
			// 数据初始化失败，写入fatal日志，因为系统数据不全，会导致程序不能正确运行
			global.OMS_LOG.Fatal(fmt.Sprintf(InitDataFailed, Mysql, in.InitializerName(), err.Error()))
			return err
		} else {
			// 数据初始化成功，写入info日志
			global.OMS_LOG.Info(fmt.Sprintf(InitDataSuccess, Mysql, in.InitializerName()))
			next = n
		}
	}
	return nil
}
