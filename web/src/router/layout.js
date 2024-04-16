const moduleDir = import.meta.glob('./modules/*.js', { eager: true })
const modules = [];
for (const path in moduleDir) {
  const module = moduleDir[path];
  modules.push(...module.default);
}

export const layoutModules = modules;

export const rootLayout = {
  path: '/',
  name: 'Layout',
  redirect: '/home',
  component: () => import('@/layout/index.vue'),
  meta: {
    title: 'layout',
    hideInBreadcrumb: true,
    hideInMenu: true,
    hideInTabs: true,
  },
  children: [
    {
      path: '/:pathMatch(.*)*',
      name: '404',
      component: () => import('@/views/error/404.vue'),
      meta: {
        title: '404',
        hideInMenu: true,
      },
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('@/views/home/index.vue'),
      meta: {
        title: '首页',
        sortMenu: 0,
      },
    },
  ],
};
