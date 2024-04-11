export default [
  {
    path: '/organize',
    component: () => import('@/views/organize/index.vue'),
    children: [
      {
        path: 'group',
        component: () => import('@/views/organize/group.vue'),
        meta: {
          title: '企业管理',
        },
      },
      {
        path: 'role',
        component: () => import('@/views/organize/role.vue'),
        meta: {
          title: '角色管理',
        },
      },
      {
        path: 'user',
        component: () => import('@/views/organize/user.vue'),
        meta: {
          title: '用户管理',
        },
      },
    ],
  },
];
