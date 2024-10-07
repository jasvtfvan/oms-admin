export const LOGIN_NAME = 'login';

export const REDIRECT_NAME = 'redirect';

export const PAGE_NOT_FOUND_NAME = '404';

// 路由白名单不需token验证
export const whiteList = [
  '/login',
  '/redirect',
  '/404',
];

// 所有用户默认加载的路由path，不需要授权的路由path
export const passedRoutes = ['/', '/home', '/static', '/404'];
