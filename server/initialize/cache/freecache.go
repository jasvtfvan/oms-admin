package cache

import (
	"errors"
	"sync"

	"github.com/coocood/freecache"
)

// size单位为Byte(1Byte=8bit)，最小要求为512KB
var sizeDefault = 100 * 1024 * 1024 // 100M

type FreecacheClient struct {
	Cache *freecache.Cache
}

var freecacheClient *FreecacheClient
var freecacheOnce sync.Once

func GetFreecacheClient() *FreecacheClient {
	if freecacheClient == nil {
		freecacheOnce.Do(func() {
			freecacheClient = &FreecacheClient{
				Cache: freecache.NewCache(sizeDefault),
			}
		})
	}
	return freecacheClient
}

func (fc *FreecacheClient) Get(key []byte) ([]byte, error) {
	return fc.Cache.Get(key)
}

// Set方法必须传入有效时间，避免代码编写问题不可排查
func (fc *FreecacheClient) Set(key, value []byte, seconds int) error {
	if seconds <= 0 {
		// TTL(Time To Live): -2代表失效，-1代表永久有效，正整数代表剩余时间(秒)
		panic(errors.New("freecache client couldn't set expiration <= 0"))
	}
	return fc.Cache.Set(key, value, seconds)
}

// 单独提炼出永久保存方法，该方法不推荐使用
func (fc *FreecacheClient) SetWithNoExpire(key, value []byte) error {
	return fc.Cache.Set(key, value, 0)
}

func (fc *FreecacheClient) Del(key []byte) bool {
	return fc.Cache.Del(key)
}

func (fc *FreecacheClient) Clear() {
	fc.Cache.Clear()
}
