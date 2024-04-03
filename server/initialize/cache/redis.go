package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/jasvtfvan/oms-admin/server/config"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

var redisClient *RedisClient
var redisOnce sync.Once

func GetRedisClient(cfg config.Redis) *RedisClient {
	if redisClient == nil {
		redisOnce.Do(func() {
			client := getConnect(cfg)
			redisClient = &RedisClient{
				Client: client,
			}
		})
	}
	return redisClient
}

func getConnect(cfg config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("redis connect ping failed, err: %s", err))
	}
	return client
}

func (rc *RedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return rc.Client.Get(ctx, key)
}

// Set方法必须传入有效时间，避免代码编写问题不可排查
func (rc *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if expiration <= 0 {
		// TTL(Time To Live): -2代表失效，-1代表永久有效，正整数代表剩余时间(秒)
		panic(errors.New("redis client couldn't set expiration <= 0"))
	}
	return rc.Client.Set(ctx, key, value, expiration)
}

// 单独提炼出永久保存方法，该方法不推荐使用
func (rc *RedisClient) SetWithNoExpire(ctx context.Context, key string, value interface{}) *redis.StatusCmd {
	return rc.Client.Set(ctx, key, value, 0)
}

func (rc *RedisClient) Del(ctx context.Context, key string) *redis.IntCmd {
	return rc.Client.Del(ctx, key)
}
