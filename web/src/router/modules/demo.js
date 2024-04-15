export default [
  {
    path: '/demo',
    component: () => import('@/views/demo/index.vue'),
    meta: {
      title: 'demo',
    },
  },
];
