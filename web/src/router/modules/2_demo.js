export default [
  {
    path: '/demo',
    name: 'demo',
    component: () => import('@/views/demo/index.vue'),
    meta: {
      title: 'demo',
      sortMenu: 2,
    },
  },
];
