import request from '@/api/request';

// 更改密码接口
export function postChangePwd(data) {
  return request.post({
    url: '/user/change-pwd',
    data,
    authorization: true,
    loading: true,
  });
}

// 获取菜单
export function getMenus(data) {
  return Promise.resolve({
    code: 200,
    data: [
      'organize',
      'organizeGroup',
      '401',
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
  return Promise.resolve({
    "code": 200,
    "data": {
      "username": "oms_admin",
      "nickName": "超级管理员",
      "avatar": "https://foruda.gitee.com/avatar/1710471233758250270/2074074_jasvtfvan_1710471233.png!avatar200",
      "phone": "",
      "email": "",
      "logOperation": true,
      "enable": false,
      "isRootAdmin": true,
      "sysGroups": [
        {
          "shortName": "根组织",
          "orgCode": "oms",
          "sort": 0,
          "sysRoles": [
            {
              "roleName": "超级管理员",
              "roleCode": "oms_admin",
              "isAdmin": true,
              "sort": 0
            }
          ]
        }
      ]
    },
    "msg": "查询成功!"
  })
  // return request.get({
  //   url: '/user/profile',
  //   params: data,
  //   authorization: true,
  //   loading: false, // 不使用该字段 == false
  // });
}
