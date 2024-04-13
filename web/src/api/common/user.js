import request from '@/api/request';

// 登录接口
export function postLogin(data) {
  return request.post({
    url: '/base/login',
    data,
    authorization: false,
    loading: true,
  });
}

// 退出接口
export function postLogout() {
  return request.post({
    url: '/base/logout',
    authorization: true,
    loading: true,
  });
}
