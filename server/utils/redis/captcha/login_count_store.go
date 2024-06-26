package captcha

import (
	"context"
	"sync"
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"go.uber.org/zap"
)

type LoginCountStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *LoginCountStore) InitCount(key string) {
	err := global.OMS_REDIS.Set(rs.Context, (rs.PreKey + key), 0, rs.Expiration).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha LoginCountStore InitCount Error:", zap.Error(err))
	}
}

func (rs *LoginCountStore) GetCount(key string) int {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + key)).Int()
	if err != nil {
		global.OMS_LOG.Error("Captcha LoginCountStore GetCount Error:", zap.Error(err))
		return 0
	}
	return val
}

func (rs *LoginCountStore) AddCount(key string) {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + key)).Int()
	value := 1
	if err == nil && val != 0 {
		value += val
	}
	err = global.OMS_REDIS.Set(rs.Context, (rs.PreKey + key), value, rs.Expiration).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha LoginCountStore AddCount Error:", zap.Error(err))
	}
}

func (rs *LoginCountStore) DelCount(key string) {
	err := global.OMS_REDIS.Del(rs.Context, (rs.PreKey + key)).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha LoginCountStore DelCount Error:", zap.Error(err))
	}
}

func (rs *LoginCountStore) UseWithCtx(ctx context.Context) *LoginCountStore {
	rs.Context = ctx
	return rs
}

/*
单例模式
*/

var loginCountStore *LoginCountStore
var loginOnce sync.Once

func GetLoginCountStore() *LoginCountStore {
	if loginCountStore == nil {
		loginOnce.Do(func() {
			loginCountStore = &LoginCountStore{
				// 由于在api声明时调用方法，拿不到global，选择在此写死
				Expiration: 1 * time.Hour,
				PreKey:     "CAPTCHA_LOGIN_COUNT_",
				Context:    context.TODO(),
			}
		})
	}
	return loginCountStore
}
