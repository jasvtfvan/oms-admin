package captcha

import (
	"context"
	"sync"
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"go.uber.org/zap"
)

type BuildCountStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *BuildCountStore) InitCount(key string) {
	err := global.OMS_REDIS.Set(rs.Context, (rs.PreKey + key), 0, rs.Expiration).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha BuildCountStore InitCount Error:", zap.Error(err))
	}
}

func (rs *BuildCountStore) GetCount(key string) int {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + key)).Int()
	if err != nil {
		global.OMS_LOG.Error("Captcha BuildCountStore GetCount Error:", zap.Error(err))
		return 0
	}
	return val
}

func (rs *BuildCountStore) AddCount(key string) {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + key)).Int()
	value := 1
	if err == nil && val != 0 {
		value += val
	}
	err = global.OMS_REDIS.Set(rs.Context, (rs.PreKey + key), value, rs.Expiration).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha BuildCountStore AddCount Error:", zap.Error(err))
	}
}

func (rs *BuildCountStore) DelCount(key string) {
	err := global.OMS_REDIS.Del(rs.Context, (rs.PreKey + key)).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha BuildCountStore DelCount Error:", zap.Error(err))
	}
}

func (rs *BuildCountStore) UseWithCtx(ctx context.Context) *BuildCountStore {
	rs.Context = ctx
	return rs
}

/*
单例模式
*/

var buildCountCountStore *BuildCountStore
var buildCountOnce sync.Once

func GetBuildCountStore() *BuildCountStore {
	if buildCountCountStore == nil {
		buildCountOnce.Do(func() {
			timeout := global.OMS_CONFIG.Captcha.OpenCaptchaTimeout
			buildCountCountStore = &BuildCountStore{
				Expiration: time.Second * time.Duration(timeout),
				PreKey:     "CAPTCHA_BUILD_COUNT_",
				Context:    context.TODO(),
			}
		})
	}
	return buildCountCountStore
}
