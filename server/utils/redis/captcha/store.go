package captcha

import (
	"context"
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

// Get implements base64Captcha.Store.
func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + key)).Result()
	if err != nil {
		global.OMS_LOG.Error("Captcha RedisStore Get Error:", zap.Error(err))
		return ""
	}
	if clear {
		err := global.OMS_REDIS.Del(rs.Context, (rs.PreKey + key)).Err()
		if err != nil {
			global.OMS_LOG.Error("Captcha RedisStore Get(Clear) Error:", zap.Error(err))
			return ""
		}
	}
	return val
}

// Set implements base64Captcha.Store.
func (rs *RedisStore) Set(key string, value string) error {
	err := global.OMS_REDIS.Set(rs.Context, (rs.PreKey + key), value, rs.Expiration).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha RedisStore Set Error:", zap.Error(err))
		return err
	}
	return nil
}

// Verify implements base64Captcha.Store.
func (rs *RedisStore) Verify(key string, answer string, clear bool) bool {
	v := rs.Get(key, clear)
	return v == answer
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func NewDefaultRedisStore() *RedisStore {
	timeout := global.OMS_CONFIG.Captcha.OpenCaptchaTimeOut
	return &RedisStore{
		Expiration: time.Second * time.Duration(timeout),
		PreKey:     "CAPTCHA_",
		Context:    context.TODO(),
	}
}
