package global

import (
	"github.com/jasvtfvan/oms-admin/server/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	OMS_VP     *viper.Viper
	OMS_LOG    *zap.Logger
	OMS_DB     *gorm.DB
	OMS_CONFIG config.Server
	OMS_REDIS  *redis.Client
)
