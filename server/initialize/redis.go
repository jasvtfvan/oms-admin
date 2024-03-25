package initialize

import (
	"context"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.OMS_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.OMS_LOG.Fatal("Redis connect ping failed, err:", zap.Error(err))
	} else {
		global.OMS_LOG.Info("Redis connect ping response:", zap.String("pong", pong))
		global.OMS_REDIS = client
	}
}
