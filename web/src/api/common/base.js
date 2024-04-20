import request from '@/api/request';

// 比较登录用的加密数据
export function postCompareSecret(data) {
  return request.post({
    url: '/base/compare-secret',
    data,
    authorization: true,
    loading: true,
  });
}

// 登录接口
export function postLogin(data) {
  return Promise.resolve({
    "code": 200,
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTE4NjQ2ODIwNTA2MzM3MjgsIlVzZXJuYW1lIjoib21zX2FkbWluIiwiTG9nT3BlcmF0aW9uIjp0cnVlLCJHcm91cHMiOlsib21zIl0sIlJvbGVzIjpbIm9tc19hZG1pbiJdLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoiRlZhbiIsImF1ZCI6WyJPTVMiXSwiZXhwIjoxNzE0MDMzNDIzLCJuYmYiOjE3MTM0Mjg2MjN9.VpfRgh_S94FQhVSt-SNx5Se2NRlyQnmpG12YoDa-5f4",
    "msg": "登录成功!"
  })
  // return request.post({
  //   url: '/base/login',
  //   data,
  //   authorization: false, // 不使用该字段 == false
  //   loading: false, // 不使用该字段 == false
  // });
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
    authorization: false, // 不使用该字段 == false
    loading: false, // 不使用该字段 == false
  });
}
