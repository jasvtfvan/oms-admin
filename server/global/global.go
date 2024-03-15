package global

import (
	"github.com/jasvtfvan/oms-admin/server/config"
	"github.com/jasvtfvan/oms-admin/server/model/goods"
	"github.com/jasvtfvan/oms-admin/server/model/system"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	OMS_VP     *viper.Viper
	OMS_LOG    *zap.Logger
	OMS_DB     *gorm.DB
	OMS_CONFIG config.Server
)

// 自动更新的表结构切片
var Tables = []interface{}{
	&system.SysUser{},
	&system.SysGroup{},
	&system.SysRole{},
	&goods.GoodsOrder{},
	// 添加其他需要迁移的表结构
}
