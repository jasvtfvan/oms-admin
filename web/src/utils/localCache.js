import { clone } from 'lodash-es';

const DB_NAME = '_local_cache_';
const OBJECT_KEY = '__@OBJECT#__';

export default {
  clear() {
    const len = window.localStorage.length;
    for (let i = len - 1; i > -1; i--) {
      const key = window.localStorage.key(i);
      if (key.includes(DB_NAME)) {
        window.localStorage.removeItem(key);
      }
    }
  },
  contain(key) {
    const cacheKey = DB_NAME + key;
    const value = window.localStorage.getItem(cacheKey);
    return !!value;
  },
  get(key) {
    const cacheKey = DB_NAME + key;
    const value = window.localStorage.getItem(cacheKey);
    if (!value) return null;
    if (value.startsWith(OBJECT_KEY)) {
      return JSON.parse(value.substring(OBJECT_KEY.length));
    }
    return value;
  },
  getKeys() {
    const keys = [];
    for (let i = 0; i < window.localStorage.length; i++) {
      let key = window.localStorage.key(i);
      if (key.includes(DB_NAME)) {
        key = key.substring(DB_NAME.length);
        keys.push(key);
      }
    }
    return keys;
  },
  getLength() {
    let len = 0;
    for (let i = 0; i < window.localStorage.length; i++) {
      const key = window.localStorage.key(i);
      if (key.includes(DB_NAME)) {
        len++;
      }
    }
    return len;
  },
  pop(key) {
    const value = clone(this.get(key));
    this.remove(key);
    return value;
  },
  remove(key) {
    const cacheKey = DB_NAME + key;
    window.localStorage.removeItem(cacheKey);
  },
  set(key, value) {
    if (key == null || value == null) return;
    const cacheKey = DB_NAME + key;
    if (typeof value === 'object') {
      value = `${OBJECT_KEY}${JSON.stringify(value)}`;
    }
    window.localStorage.setItem(cacheKey, value);
  },
};
