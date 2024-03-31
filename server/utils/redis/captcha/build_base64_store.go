package captcha

import (
	"context"
	"sync"
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BuildBase64Store struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

// Get implements base64Captcha.Store.
func (rs *BuildBase64Store) Get(key string, clear bool) string {
	val, err := global.OMS_REDIS.Get(rs.Context, (rs.PreKey + "ANSWER_" + key)).Result()
	if err != nil {
		global.OMS_LOG.Error("Captcha BuildBase64Store Get Error:", zap.Error(err))
		return ""
	}
	if clear {
		err := global.OMS_REDIS.Del(rs.Context, (rs.PreKey + "ANSWER_" + key)).Err()
		if err != nil {
			global.OMS_LOG.Error("Captcha BuildBase64Store Get(Clear) Error:", zap.Error(err))
			return ""
		}
	}
	return val
}

// Set implements base64Captcha.Store.
func (rs *BuildBase64Store) Set(key string, value string) error {
	err := global.OMS_REDIS.Set(rs.Context, (rs.PreKey + "ANSWER_" + key), value, rs.Expiration).Err()
	if err != nil {
		global.OMS_LOG.Error("Captcha BuildBase64Store Set Error:", zap.Error(err))
		return err
	}
	return nil
}

// Verify implements base64Captcha.Store.
func (rs *BuildBase64Store) Verify(key string, answer string, clear bool) bool {
	v := rs.Get(key, clear)
	return v == answer
}

func (rs *BuildBase64Store) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

/*
单例模式
*/

var base64Store *BuildBase64Store
var base64Once sync.Once

func GetBuildBase64Store() *BuildBase64Store {
	if base64Store == nil {
		base64Once.Do(func() {
			base64Store = &BuildBase64Store{
				Expiration: time.Minute * 3,
				PreKey:     "CAPTCHA_BUILD_BASE64_",
				Context:    context.TODO(),
			}
		})
	}
	return base64Store
}
