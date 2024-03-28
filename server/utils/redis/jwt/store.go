package jwt

import (
	"context"
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"go.uber.org/zap"
)

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + key)).Result()
	if err != nil {
		global.OMS_LOG.Error("JWT RedisStore Get Error:", zap.Error(err))
		return ""
	}
	if clear {
		err := global.OMS_REDIS.Del(rs.Context, (rs.PreKey + key)).Err()
		if err != nil {
			global.OMS_LOG.Error("JWT RedisStore Get(Clear) Error:", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Set(key string, value string) error {
	err := global.OMS_REDIS.Set(rs.Context, (rs.PreKey + key), value, rs.Expiration).Err()
	if err != nil {
		global.OMS_LOG.Error("JWT RedisStore Set Error:", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Verify(key string, answer string, clear bool) bool {
	v := rs.Get(key, clear)
	return v == answer
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) *RedisStore {
	rs.Context = ctx
	return rs
}

func NewDefaultRedisStore() *RedisStore {
	exp, _ := utils.ParseDuration(global.OMS_CONFIG.JWT.ExpiresTime)
	return &RedisStore{
		Expiration: exp,
		PreKey:     "JWT_",
		Context:    context.TODO(),
	}
}
