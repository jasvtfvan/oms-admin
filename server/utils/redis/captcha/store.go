package captcha

import (
	"context"
	"sync"
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

/*
单例模式
*/

var redisStore *RedisStore
var once sync.Once

func GetRedisStore() *RedisStore {
	if redisStore == nil {
		once.Do(func() {
			redisStore = &RedisStore{
				Expiration: time.Minute * 3,
				PreKey:     "CAPTCHA_",
				Context:    context.TODO(),
			}
		})
	}
	return redisStore
}

/*
记录验证码次数
*/
func (rs *RedisStore) GetCount(key string) (int, error) {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + "COUNT_" + key)).Int()
	return val, err
}
func (rs *RedisStore) InitCount(key string) {
	timeout := global.OMS_CONFIG.Captcha.OpenCaptchaTimeout
	err := global.OMS_REDIS.Set(rs.Context, (rs.PreKey + "COUNT_" + key), 1, time.Second*time.Duration(timeout)).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha RedisStore InitCount Error:", zap.Error(err))
	}
}
func (rs *RedisStore) AddCount(key string) {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + "COUNT_" + key)).Int()
	if val == 0 || err != nil {
		rs.InitCount(key)
	} else {
		timeout := global.OMS_CONFIG.Captcha.OpenCaptchaTimeout
		err := global.OMS_REDIS.Set(rs.Context, (rs.PreKey + "COUNT_" + key), val+1, time.Second*time.Duration(timeout)).Err()
		if err != nil {
			global.OMS_LOG.Error("Captcha RedisStore AddCount Error:", zap.Error(err))
		}
	}
}
func (rs *RedisStore) DelCount(key string) {
	err := global.OMS_REDIS.Del(rs.Context, (rs.PreKey + "COUNT_" + key)).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha RedisStore DelCount Error:", zap.Error(err))
	}
}
