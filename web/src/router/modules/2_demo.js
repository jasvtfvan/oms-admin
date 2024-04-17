export default [
  {
    path: '/demo',
    name: 'demo',
    component: () => import('@/views/demo/index.vue'),
    meta: {
      title: 'demo',
      sortMenu: 102,
      notAdminDefault: true, // 管理员默认不选中
    },
  },
  {
    path: '/401',
    name: '401',
    component: () => import('@/views/error/401.vue'),
    meta: {
      title: '401',
      sortMenu: 103,
      notAdminDefault: true, // 管理员默认不选中
    },
  },
];
