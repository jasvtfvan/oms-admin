import { RouterView } from 'vue-router';
import { REDIRECT_NAME, PAGE_NOT_FOUND_NAME } from './constant';

const moduleDir = import.meta.glob('./modules/*.js', { eager: true })
const modules = [];
for (const path in moduleDir) {
  const module = moduleDir[path];
  modules.push(...module.default);
}

// 所有业务模块里的路由
export const dynamicModules = modules;
// 默认layout里的路由
export const defaultModules = [{
  path: '/home',
  name: 'home',
  component: () => import('@/views/home/index.vue'),
  meta: {
    icon: 'HomeOutlined',
    title: '首页',
    keepAlive: true,
    sortMenu: 0,
  },
},{
  path: '/static',
  name: 'static',
  component: () => import('@/views/home/static.vue'),
  meta: {
    icon: 'HomeOutlined',
    title: 'Static',
    keepAlive: true,
    sortMenu: 1,
  },
}];

export const rootLayout = {
  path: '/',
  name: 'Layout',
  redirect: '/home',
  component: () => import('@/layout/index.vue'),
  meta: {
    title: 'layout',
    hideInBreadcrumb: false,
    hideInMenu: false,
    hideInTabs: false,
  },
  children: [
    {
      path: '/:pathMatch(.*)*',
      name: PAGE_NOT_FOUND_NAME,
      component: () => import('@/views/error/404.vue'),
      meta: {
        title: '404',
        hideInMenu: true,
      },
    },
    ...defaultModules,
  ],
};

/**
 * 重定向路由 主要用于刷新当前页面
 */
export const redirectRoute = {
  path: '/redirect',
  name: 'RedirectTo',
  meta: {
    title: 'redirect',
    hideInBreadcrumb: true,
    hideInMenu: true,
    hideInTabs: true,
  },
  children: [
    {
      path: ':path(.*)',
      name: REDIRECT_NAME,
      component: RouterView,
      meta: {
        title: 'redirect',
        hideInMenu: true,
      },
      beforeEnter: (to) => {
        const { params, query } = to;
        const { path, redirectType = 'path' } = params;

        Reflect.deleteProperty(params, '_redirect_type');
        Reflect.deleteProperty(params, 'path');

        const _path = Array.isArray(path) ? path.join('/') : path;
        setTimeout(() => {
          if (redirectType === 'name') {
            router.replace({
              name: _path,
              query,
              params,
            });
          } else {
            router.replace({
              path: _path.startsWith('/') ? _path : `/${_path}`,
              query,
            });
          }
        });
        return true;
      },
    },
  ],
};
