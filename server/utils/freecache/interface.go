package freecache

type String string
type Bool bool

type CacheInterface interface {
	Get(key string, valStruct interface{}) interface{}
	Set(key string, value interface{}) error
	Verify(key string, value interface{}, valStruct interface{}) bool
	Del(key string) bool
	Clear()
}
