export default [
  {
    path: '/organize',
    name: 'organize',
    meta: {
      title: '组织管理',
      sortMenu: 101,
    },
    redirect: '/organize/group',
    component: () => import('@/views/organize/index.vue'),
    children: [
      {
        path: '/organize/group',
        name: 'organizeGroup',
        component: () => import('@/views/organize/group.vue'),
        meta: {
          title: '企业管理',
          sortMenu: 10001,
          keepAlive: true,
        },
      },
      {
        path: '/organize/role',
        name: 'organizeRole',
        component: () => import('@/views/organize/role.vue'),
        meta: {
          title: '角色管理',
          sortMenu: 10002,
          keepAlive: true,
        },
      },
      {
        path: '/organize/user',
        name: 'organizeUser',
        component: () => import('@/views/organize/user.vue'),
        meta: {
          title: '用户管理',
          sortMenu: 10003,
        },
      },
    ],
  },
];
