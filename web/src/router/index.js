import { createRouter, createWebHistory } from 'vue-router'
import Nprogress from 'nprogress'
import { useUserStore } from '@/stores/user'
import { rootLayout } from './layout'
import globRoutes, { LOGIN_NAME } from './unLayout'

export const routes = [
  ...globRoutes,
  rootLayout,
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})


/**
 * 路由守卫
 */
const whiteList = [ // 白名单不需token验证
  '/login',
]

// 是否跳转权限
function hasRoute(to, menus = []) {
  const passedRoutes = ['/', '/home', '/404'];
  if (passedRoutes.includes(to.path)) return true
  return menus.includes(to.name)
}

router.beforeEach(async (to, from, next) => {
  Nprogress.start()
  const userStore = useUserStore()
  const { token, menus } = userStore; // 获取token,menus
  if (token) { // token存在
    if (to.name == LOGIN_NAME) { // 登录页面直接跳到home
      next({ path: '/', replace: true });
    } else {
      if (hasRoute(to, menus)) { // 有权限
        next()
      } else { // 没有权限
        next({ name: '404' })
      }
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
