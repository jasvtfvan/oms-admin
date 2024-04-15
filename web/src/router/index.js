import { createRouter, createWebHistory } from 'vue-router'
import Nprogress from 'nprogress'
import { useUserStore } from '@/stores/user'

const moduleDir = import.meta.glob('./modules/*.js', { eager: true })
const modules = [];
for (const path in moduleDir) {
  const module = moduleDir[path];
  modules.push(...module.default);
}

const LOGIN_NAME = 'login'; // 第一个页面，通常是登录
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
    path: '/login',
    name: LOGIN_NAME,
    component: () => import('@/views/login/index.vue'),
    meta: {
      title: '登录',
    },
  },
  {
    path: '/',
    name: 'Layout',
    redirect: '/dashboard',
    component: () => import('@/layout/index.vue'),
    children: [
      ...modules,
    ],
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})


/**
 * 路由守卫
 */
const whiteList = [
  '/login',
]

function hasPermission(to) {
  // 获取用户权限列表
  const perms = []; // TODO 获取用户的权限，后边实现
  const routePerms = to.meta && to.meta.perms;
  if (!routePerms || routePerms.includes('*')) return true;
  if (routePerms) {
    return perms.some(perm => routePerms.includes(perm))
  }
}

router.beforeEach(async (to, from, next) => {
  Nprogress.start()
  const userStore = useUserStore()
  const { token } = userStore; // 获取token
  if (token) { // token存在
    if (hasPermission(to)) { // 有权限
      next()
    } else { // 没有权限
      next({ name: '401' })
    }
  } else { // token不存在
    if (whiteList.includes(to.path)) { // 白名单
      next()
    } else {
      next({ name: LOGIN_NAME, query: { redirect: to.fullPath }, replace: true });
    }
  }
})

router.afterEach(async (to) => {
  Nprogress.done();
  if (to.meta && to.meta.title) {
    document.title = to.meta.title;
  }
})

router.onError((error) => {
  console.error('路由错误:', error)
})


export default router
