package freecache

import (
	"encoding/json"
	"reflect"
	"sync"
	"time"

	"github.com/jasvtfvan/oms-admin/server/global"
	"github.com/jasvtfvan/oms-admin/server/utils"
	"go.uber.org/zap"
)

type StoreDefault struct {
	Expiration int
	PreKey     string
}

/*
单例模式
*/
var storeDefault *StoreDefault
var onceDefault sync.Once

func GetStoreDefault() *StoreDefault {
	if storeDefault == nil {
		onceDefault.Do(func() {
			exp, _ := utils.ParseDuration("1d")
			storeDefault = &StoreDefault{
				Expiration: int(exp / time.Second), // 超时时间转换成秒,
				PreKey:     "FREECACHE_",
			}
		})
	}
	return storeDefault
}

/*
通用方法区
*/
func (sd *StoreDefault) Get(key string, valStruct interface{}) interface{} {
	val, err := global.OMS_FREECACHE.Get([]byte(key))
	if err != nil {
		global.OMS_LOG.Error("freecache StoreDefault Get Error:", zap.Error(err))
		return nil
	}
	value := &valStruct
	// 尝试将缓存中的字节切片反序列化为传入的指针所指向的结构体
	if err := json.Unmarshal(val, value); err != nil {
		global.OMS_LOG.Error("freecache StoreDefault Get Error:", zap.Error(err))
		return nil
	}
	return *value
}
func (sd *StoreDefault) Set(key string, value interface{}) error {
	val, err := json.Marshal(value)
	if err != nil {
		global.OMS_LOG.Error("freecache StoreDefault Set Error:", zap.Error(err))
		return err
	}
	return global.OMS_FREECACHE.Set([]byte(key), val, sd.Expiration)
}
func (sd *StoreDefault) Verify(key string, value interface{}, valStruct interface{}) bool {
	valGet := sd.Get(key, valStruct)
	if valGet == nil {
		return false
	}
	return value == valGet || reflect.DeepEqual(value, valGet)
}
func (sd *StoreDefault) Del(key string) bool {
	return global.OMS_FREECACHE.Del([]byte(key))
}
func (sd *StoreDefault) Clear() {
	global.OMS_FREECACHE.Clear()
}
