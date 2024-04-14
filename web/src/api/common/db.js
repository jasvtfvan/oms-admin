import request from '@/api/request';

// 判断是否需要初始化
export function postInitCheck(data) {
  return request.post({
    url: '/init/check',
    data,
    authorization: false, // 不使用该字段 == false
    loading: false, // 不使用该字段 == false
  });
}

// 初始化数据库
export function postInitDb(data) {
  return request.post({
    url: '/init/db',
    data,
    authorization: false, // 不使用该字段 == false
    loading: true,
  });
}

// 判断是否需要升级
export function postUpdateCheck(data) {
  return request.post({
    url: '/update/check',
    data,
    authorization: true,
    loading: true,
  });
}

// 升级数据库
export function postUpdateDb(data) {
  return request.post({
    url: '/update/db',
    data,
    authorization: true,
    loading: true,
  });
}
