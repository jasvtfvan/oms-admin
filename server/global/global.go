package global

import (
	"github.com/coocood/freecache"
	"github.com/jasvtfvan/oms-admin/server/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	OMS_VP        *viper.Viper     // 加载config.yaml配置文件 -> OMS_CONFIG
	OMS_LOG       *zap.Logger      // zap日志
	OMS_DB        *gorm.DB         // 数据库DB
	OMS_CONFIG    config.Server    // config.yaml配置文件所有属性
	OMS_REDIS     *redis.Client    // redis缓存
	OMS_FREECACHE *freecache.Cache // freecache单机缓存
)
