import axios from 'axios';
import { useUserStore } from '@/stores/user';
import { message } from 'ant-design-vue';
import { useLoading } from '@/hooks/useLoading';
import { doCommonLogout } from '@/utils/util';

const { showLoading, hideLoading } = useLoading();

const STATUS_MAP = {
  200: '请求成功',
  400: '请求失败',
  401: '没有权限',
  404: '请求地址不存在',
  500: '服务器内部错误',
};

const BASE_API = import.meta.env.VITE_BASE_API

/**
 * 防止返回前重复请求
 */
const pending = [];
function addPending(url, method) {
  let completeUrl = url.startsWith('http') ? url : BASE_API + url;
  completeUrl += (`&${method.toUpperCase()}`);
  if (!pending.includes(completeUrl)) {
    pending.push(completeUrl);
  }
}
function removePending(url, method) {
  let completeUrl = url.startsWith('http') ? url : BASE_API + url;
  completeUrl += (`&${method.toUpperCase()}`);
  if (pending.includes(completeUrl)) {
    const index = pending.indexOf(completeUrl);
    pending.splice(index, 1);
  }
}
function isPending(url, method) {
  let completeUrl = url.startsWith('http') ? url : BASE_API + url;
  completeUrl += (`&${method.toUpperCase()}`);
  return pending.includes(completeUrl);
}

let activeLoadingCount = 0
let loadingTimer
// 执行loading
function doShowLoading() {
  activeLoadingCount++
  if (loadingTimer) clearTimeout(loadingTimer)
  loadingTimer = setTimeout(() => {
    if (activeLoadingCount > 0) {
      showLoading()
    }
  }, 400); // 如果请求在0.4秒就全部返回则不需要loading
}
// 隐藏loading
function doHideLoading() {
  activeLoadingCount--
  if (activeLoadingCount <= 0) {
    activeLoadingCount = 0;
    if (loadingTimer) clearTimeout(loadingTimer); // 0.4秒全部返回不需要loading
    hideLoading();
  }
}

// 提示信息关闭后
async function onToastClose(status) {
  if (status && /^401|425|429$/.test(status)) { // 401 425 429
    doCommonLogout()
  }
}

// 处理失败
function authorizationInvalidate(code, msg) {
  if (code == 401) {
    message.error('连接超时，请重新登录', 2, async () => {
      doCommonLogout()
    });
  } else {
    message.error(msg || '请求数据失败', 2);
  }
}

// 统一处理错误
function handleRequestError(err) {
  let code
  let msg
  if (err.response) {
    console.error('请求失败:', err.response)
    const { status, data, statusText } = err.response;
    if (data) { // data存在
      if (typeof data == 'object') {
        const retMsg = data.msg || data.message || data.ErrorMessage || statusText;
        code = data.code == null ? status : data.code
        msg = retMsg
        message.error(retMsg, 2, () => onToastClose(status));
      } else {
        code = status
        msg = data
        message.error(data, 2, () => onToastClose(status));
      }
    } else { // data不存在
      if (Object.hasOwnProperty.call(STATUS_MAP, status)) { // 已经定义了错误状态
        code = status
        msg = statusText || STATUS_MAP[status]
        message.error(statusText || STATUS_MAP[status], 2, () => onToastClose(status));
      } else { // 未定义错误状态
        code = status
        msg = statusText || '服务器响应错误'
        message.error(statusText || '服务器响应错误', 2, () => onToastClose(status));
      }
    }
    return Promise.reject({ code, msg, });
  } else if (err.request) {
    msg = '请求没有响应'
    console.error('请求没有响应:', err.request)
    message.error(msg, 2);
  } else {
    msg = '请求配置出错'
    console.error('请求配置出错:', err.message)
    message.error(msg, 2);
  }
  code = -1;
  msg = `${err.code} | ${err.message}`;
  return Promise.reject({ code, msg, });
}

// 下载blob二进制文件
async function downloadFile(response) {
  const url = window.URL.createObjectURL(new Blob([response.data]));
  const filename = response.headers['x-filename'];

  return axios.get(url, { responseType: 'blob' }).then((res) => {
    const blob = new Blob([res.data]);
    if (window.navigator.msSaveBlob) {
      // 兼容 IE，使用 msSaveBlob 方法进行下载
      window.navigator.msSaveBlob(blob, decodeURIComponent(filename));
    } else {
      // 创建一个 <a> 元素
      const link = document.createElement('a');
      link.href = window.URL.createObjectURL(blob);
      link.setAttribute('download', decodeURIComponent(filename));
      // 模拟点击下载
      link.click();
      // 清理 URL 和 <a> 元素
      link.remove();
      window.URL.revokeObjectURL(url);
    }
    return Promise.resolve();
  });
}

class HttpRequest {
  constructor() {
    this.timeout = 120 * 1000; // 120秒
  }

  static setInterceptors(instance, url, method) {
    instance.interceptors.request.use((config) => {
      const userStore = useUserStore();
      if (config.loading) {
        doShowLoading();
      }
      const conf = config;
      const { authorization } = conf;
      if (!conf.headers) {
        conf.headers = { Accept: 'application/json, text/plain, */*' };
      }
      const token = userStore.token;
      const group = userStore.group;
      if (authorization && token) {
        conf.headers['x-token'] = token;
        conf.headers['x-group'] = group;
      }
      return conf;
    }, (err) => {
      if (err.config.loading) {
        doHideLoading();
      }
      removePending(url, method);
      return Promise.reject(err);
    });

    instance.interceptors.response.use((res) => {
      if (res.config.loading) {
        doHideLoading();
      }
      removePending(url, method);
      /** 下载请求 */
      const contentType = res.headers['content-type'];
      if (contentType === 'application/octet-stream' ||
        contentType === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet') {
        // 注意：Blob类型文件下载需要请求头参数添加 responseType:'blob'  下载 导出等功能
        return downloadFile(res)
      } else {
        /** 非下载请求 */
        const { status, data, statusText } = res;
        if (!data) {
          return Promise.reject({ code: status, msg: statusText });
        }
        if (/^[2-3]0\d$/.test(status) && /^[2-3]0\d$/.test(data.code)) {
          return Promise.resolve(data);
        }
        authorizationInvalidate(data.code, (data.msg || statusText));
        return Promise.reject(data);
      }
    }, (err) => {
      if (err.config.loading) {
        doHideLoading();
      }
      removePending(url, method);
      return handleRequestError(err);
    });
  }

  // 合并选项
  mergeOptions(options) {
    const opts = options;
    const baseURL = BASE_API;
    return {
      // withCredentials: true,
      baseURL,
      timeout: this.timeout,
      ...opts
    };
  }

  // 创建请求
  request(options) {
    const { url, method } = options;
    addPending(url, method);
    const instance = axios.create();
    HttpRequest.setInterceptors(instance, url, method);
    const opts = this.mergeOptions(options);
    return instance(opts);
  }

  // get方法
  get(config) {
    const { url, data, ...opts } = config;
    if (!url) {
      return Promise.reject();
    }
    const force = opts.force || false; // 强制请求，可以重复提交
    if (!force && isPending(url, 'get')) {
      return Promise.reject(`${url} is pending`);
    }
    delete opts.force;
    return this.request({
      method: 'get',
      url,
      data,
      ...opts,
    });
  }

  // post方法
  post(config) {
    const { url, data, ...opts } = config;
    if (!url) {
      return Promise.reject();
    }
    const force = opts.force || false; // 强制请求，可以重复提交
    if (!force && isPending(url, 'post')) {
      return Promise.reject(`${url} is pending`);
    }
    delete opts.force;
    return this.request({
      method: 'post',
      url,
      data,
      ...opts,
    });
  }

  // put方法
  put(config) {
    const { url, data, ...opts } = config;
    if (!url) {
      return Promise.reject();
    }
    const force = opts.force || false; // 强制请求，可以重复提交
    if (!force && isPending(url, 'post')) {
      return Promise.reject(`${url} is pending`);
    }
    delete opts.force;
    return this.request({
      method: 'put',
      url,
      data,
      ...opts,
    });
  }

  // delete方法
  delete(config) {
    const { url, data, ...opts } = config;
    if (!url) {
      return Promise.reject();
    }
    const force = opts.force || false; // 强制请求，可以重复提交
    if (!force && isPending(url, 'delete')) {
      return Promise.reject(`${url} is pending`);
    }
    delete opts.force;
    return this.request({
      method: 'delete',
      url,
      data,
      ...opts,
    });
  }

}

export default new HttpRequest();
