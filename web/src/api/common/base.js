import request from '@/api/request';

// 登录接口
export function postLogin(data) {
  // return request.post({
  //   url: '/base/login',
  //   data,
  //   authorization: false, // 不使用该字段 == false
  //   loading: false, // 不使用该字段 == false
  // });
  return Promise.resolve({
    "code": 200,
    "data": {
      "user": {
        "username": "oms_admin",
        "nickName": "超级管理员",
        "avatar": "https://foruda.gitee.com/avatar/1710471233758250270/2074074_jasvtfvan_1710471233.png!avatar200",
        "phone": "",
        "email": "",
        "isRootAdmin": true,
        "logOperation": true,
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
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTE4NjQ2ODIwNTA2MzM3MjgsIlVzZXJuYW1lIjoib21zX2FkbWluIiwiTG9nT3BlcmF0aW9uIjp0cnVlLCJHcm91cHMiOlt7Ik9yZ0NvZGUiOiJvbXMiLCJTaG9ydE5hbWUiOiLmoLnnu4Tnu4cifV0sIlJvbGVzIjpbeyJSb2xlQ29kZSI6Im9tc19hZG1pbiIsIlJvbGVOYW1lIjoi6LaF57qn566h55CG5ZGYIn1dLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoiRlZhbiIsImF1ZCI6WyJPTVMiXSwiZXhwIjoxNzEzOTM5OTEzLCJuYmYiOjE3MTMzMzUxMTN9.H7bDsQ8-yen48oudrZEq4qaVgDMhuYrg0Q5xymSZdZQ"
    },
    "msg": "登录成功!"
  })
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
