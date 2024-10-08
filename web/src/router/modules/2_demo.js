export default [
  {
    path: '/demo',
    name: 'demo',
    component: () => import('@/views/demo/index.vue'),
    meta: {
      title: 'demo',
      sortMenu: 102,
      hideInTabs: true,
    },
  },
  {
    path: '/401',
    name: '401',
    component: () => import('@/views/error/401.vue'),
    meta: {
      title: '401',
      sortMenu: 103,
      notAdminDefault: true, // 管理员默认没有该菜单，需要从后台读取
    },
  },
];
