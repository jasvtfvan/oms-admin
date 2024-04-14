import request from '@/api/request';

// 登录接口
export function postLogin(data) {
  return request.post({
    url: '/base/login',
    data,
    authorization: false,
    loading: false,
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

// 获取验证码
export function postCaptcha(data) {
  return request.post({
    url: '/base/captcha',
    data,
    authorization: false,
    loading: false,
  });
}
