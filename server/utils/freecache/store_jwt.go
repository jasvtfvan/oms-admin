package freecache

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"go.uber.org/zap"
)

type StoreJWT struct {
	Expiration int
	PreKey     string
	Context    context.Context
}

/*
单例模式
*/
var storeJWT *StoreJWT
var onceJWT sync.Once

func GetStoreJWT() *StoreJWT {
	if storeJWT == nil {
		onceJWT.Do(func() {
			exp, _ := utils.ParseDuration(global.OMS_CONFIG.JWT.ExpiresTime)
			storeJWT = &StoreJWT{
				Expiration: int(exp / time.Second), // 超时时间转换成秒,
				PreKey:     "F_JWT_",
				Context:    context.TODO(),
			}
		})
	}
	return storeJWT
}

func (sj *StoreJWT) Get(key string, clear bool) string {
	targetKey := []byte(sj.PreKey + key)
	val, err := global.OMS_FREECACHE.Get(targetKey)
	if err != nil {
		global.OMS_LOG.Error("JWT StoreJWT Get Error:", zap.Error(err))
		return ""
	}
	if clear {
		ok := global.OMS_FREECACHE.Del(targetKey)
		if !ok {
			global.OMS_LOG.Error("JWT StoreJWT Get(Clear) Error")
			return ""
		}
	}
	return string(val)
}

func (sj *StoreJWT) Set(key string, value string) error {
	targetKey := []byte(sj.PreKey + key)
	targetValue := []byte(value)
	err := global.OMS_FREECACHE.Set(targetKey, targetValue, sj.Expiration)
	if err != nil {
		global.OMS_LOG.Error("JWT StoreJWT Set Error:", zap.Error(err))
		return err
	}
	return nil
}

func (sj *StoreJWT) Del(key string) error {
	targetKey := []byte(sj.PreKey + key)
	ok := global.OMS_FREECACHE.Del(targetKey)
	if !ok {
		global.OMS_LOG.Error("JWT StoreJWT Del Error")
		return errors.New("JWT StoreJWT Del Error")
	}
	return nil
}

func (sj *StoreJWT) Verify(key string, answer string, clear bool) bool {
	v := sj.Get(key, clear)
	return v == answer
}

func (sj *StoreJWT) UseWithCtx(ctx context.Context) *StoreJWT {
	sj.Context = ctx
	return sj
}
