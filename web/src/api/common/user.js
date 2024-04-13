import request from '@/api/request';

// 登录接口
export function login(data) {
  return request.post({
    url: '/api/GameLogin/Login',
    data,
    authorization: false,
    loading: true,
  });
}

// 退出接口
export function logout() {
  return request.get({
    url: '/api/GameLogin/Logout',
    authorization: true,
    loading: true,
  });
}
