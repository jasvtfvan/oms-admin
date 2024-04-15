import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/index.vue'
import Nprogress from 'nprogress'
import { useUserStore } from '@/stores/user'

const moduleDir = import.meta.glob('./modules/*.js', { eager: true })
const modules = [];
for (const path in moduleDir) {
  const module = moduleDir[path];
  modules.push(...module.default);
}

const routes = [
  {
    path: '/401',
    name: '401',
    component: () => import('@/views/error/401.vue'),
    meta: {
      title: '401',
    },
  },
  {
    path: '/:pathMatch(.*)*',
    name: '404',
    component: () => import('@/views/error/404.vue'),
    meta: {
      title: '404',
    },
  },
  {
    path: '/',
    name: 'Index',
    component: Index,
  },
  ...modules,
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})


/**
 * 路由守卫
 */

function hasPermission(to) {
  // 获取用户权限列表
  const perms = []; // TODO 获取用户的权限，后边实现
  const routePerms = to.meta && to.meta.perms;
  if (!routePerms || routePerms.includes('*')) return true;
  if (routePerms) {
    return perms.some(perm => routePerms.includes(perm))
  }
}

const whiteList = [
  '/',
]

router.beforeEach(async (to, from, next) => {
  if (whiteList.includes(to.path)) { // 白名单
    Nprogress.start()
    return next()
  }
  const userStore = useUserStore()
  const { token } = userStore; // 获取token
  if (token) { // token存在
    if (hasPermission(to)) { // 有权限
      Nprogress.start()
      return next()
    } else { // 没有权限
      Nprogress.start()
      return next({ name: '401' })
    }
  } else { // token不存在,跳到/根路径
    let queryStr = '';
    if (to.query && Object.keys(to.query).length > 0) {
      Object.keys(to.query).forEach(key => {
        queryStr += `&${key}=${to.query[key]}`;
      });
    }
    Nprogress.start()
    return next(`/?redirect=${to.path}${queryStr}`);
  }
})

router.afterEach(async (to) => {
  if (to.meta && to.meta.title) {
    document.title = to.meta.title;
  }
  Nprogress.done();
})

router.onError((error) => {
  console.error('路由错误:', error)
})


export default router
