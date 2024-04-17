import request from '@/api/request';

// 获取菜单
export function getMenus(data) {
  return Promise.resolve({
    code: 200,
    data: [
      'organize',
      'organizeGroup',
      'demo',
    ],
    msg: 'ok',
  })
  // return request.get({
  //   url: '/user/menus',
  //   params: data,
  //   authorization: true,
  //   loading: false, // 不使用该字段 == false
  // });
}

// 获取登录用户信息
export function getUserProfile(data) {
  return request.get({
    url: '/user/profile',
    params: data,
    authorization: true,
    loading: false, // 不使用该字段 == false
  });
}
