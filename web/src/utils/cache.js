const DEFAULT_EXPIRE_TIME = 6 * 24 * 60 * 60; // 缓存单位秒，缓存6天
const DEFAULT_LOCAL_PREFIX_KEY = '_local_cache_';
const DEFAULT_SESSION_PREFIX_KEY = '_session_cache_';

// 创建本地缓存对象
export const createStorage = ({ prefixKey = '', storage = localStorage } = {}) => {
  const _prefixKey = Symbol('_prefixKey');
  const _storage = Symbol('_storage');
  const _getKey = Symbol('_getKey');

  class Storage {
    constructor(prefixKey, storage) {
      this[_prefixKey] = prefixKey
      this[_storage] = storage
    }
    [_getKey](key) {
      return `${this[_prefixKey]}${key}`.toUpperCase()
    }
    // 删除
    remove(key) {
      let storageType = 'unknown storage'
      if (this[_storage] instanceof localStorage) {
        storageType = 'localStorage'
      } else if (this[_storage] instanceof localStorage) {
        storageType = 'sessionStorage'
      }
      console.warn(`${storageType}.remove:`, this[_getKey](key))
      this[_storage].removeItem(this[_getKey](key))
    }
    // 设置缓存(单位秒)
    set(key, value, expire) {
      // expire：如果是0则永久有效；如果是null则使用默认时长；否则根据是否数字判断
      let expireValue
      if (expire === 0) {
        expireValue = 'infinite'
      } else if (expire == null || isNaN(expire * 1000)) {
        expireValue = new Date().getTime() + DEFAULT_EXPIRE_TIME * 1000
      } else {
        expireValue = new Date().getTime() + expire * 1000
      }
      const stringData = JSON.stringify({
        value,
        expire: expireValue,
      })
      this[_storage].setItem(this[_getKey](key), stringData);
    }
    // 读取
    get(key) {
      const item = this[_storage].getItem(this[_getKey](key));
      if (item) {
        try {
          const data = JSON.parse(item);
          const { value, expire } = data;
          if (expire === 'infinite' || expire >= Date.now()) {
            return value
          }
          this.remove(this[_getKey](key))
        } catch (_) {
          return null
        }
      }
      return null
    }
    // 清除缓存
    clear() {
      const len = this[_storage].length;
      for (let i = len - 1; i > -1; i--) {
        const key = this[_storage].key(i);
        if (key.includes(_prefixKey)) {
          this[_storage].removeItem(key);
        }
      }
    }
  };
  return new Storage(prefixKey, storage)
}

export const localCache = createStorage({ prefixKey: DEFAULT_LOCAL_PREFIX_KEY, storage: localStorage });
export const sessionCache = createStorage({ prefixKey: DEFAULT_SESSION_PREFIX_KEY, storage: sessionStorage });

export default {
  localCache,
  sessionCache,
};
