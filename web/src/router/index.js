import { createRouter, createWebHistory } from 'vue-router'
import Nprogress from 'nprogress'
import { useUserStore } from '@/stores/user'
import { rootLayout } from './layout'
import unLayoutRoutes, { LOGIN_NAME } from './unLayout'
import { message } from 'ant-design-vue'
import { decryptPwd } from '@/utils/cryptoLoginSecret'
import { isValidPassword } from '@/utils/util'
import $bus from '@/utils/bus'
import { nextTick } from 'vue'

export const routes = [
  ...unLayoutRoutes,
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
function hasRoute(to, menuNames = []) {
  const passedRoutes = ['/', '/home', '/404']; // 所有默认加载的路由path
  if (passedRoutes.includes(to.path)) return true
  return menuNames.includes(to.name)
}

router.beforeEach(async (to, from, next) => {
  Nprogress.start()
  const userStore = useUserStore()
  // menuNames不包含/home等默认路由
  let { token, menuNames, userProfile } = userStore; // 获取token,menuNames
  if (token) { // token存在
    if (to.name == LOGIN_NAME) { // 登录页面直接跳到home
      next({ path: '/', replace: true });
    } else {
      // 如果菜单或用户没有加载，进行加载
      if (menuNames.length <= 0 || !userProfile.username) {
        try {
          const authRes = await userStore.GetAuthWithoutLogin()
          menuNames = authRes.menuNames
        } catch (error) {
          message.error(error.msg || '获取所有权限失败', 2, async () => {
            // 加载失败，退出
            await userStore.Logout()
            window.location.reload();
          });
        }
      }
      if (hasRoute(to, menuNames)) { // 有权限
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

/**
 * 验证是否需要改密码
 */
function judgeChangePassword() {
  const userStore = useUserStore()
  try {
    const password = decryptPwd();
    const forceChangePwd = isValidPassword(password) ? false : true;
    nextTick(() => {
      console.log('router.afterEach-forceChangePwd', forceChangePwd)
      $bus.emit('changePasswordForce', { open: forceChangePwd })
    })
  } catch (error) {
    message.error(error.message || '验证安全过程中发现异常', 2, async () => {
      // 加载失败，退出
      await userStore.Logout()
      window.location.reload();
    });
  }
}

router.afterEach(async (to) => {
  Nprogress.done();
  if (to.meta && to.meta.title) {
    document.title = to.meta.title;
  }
  // 不在白名单的，判断是否需要修改密码
  if (!whiteList.includes(to.path)) {
    judgeChangePassword()
  }
})

router.onError((error) => {
  console.error('路由错误:', error)
})


export default router
