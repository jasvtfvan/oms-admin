import { LOGIN_NAME } from './constant';

export default [
  {
    path: '/login',
    name: LOGIN_NAME,
    component: () => import('@/views/login/index.vue'),
    meta: {
      title: '登录',
      hideInMenu: true,
      hideInTabs: true,
      hideInBreadcrumb: true,
    },
  },
]
