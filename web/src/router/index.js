import { createRouter, createWebHistory } from 'vue-router'
import Nprogress from 'nprogress'
import { useUserStore } from '@/stores/user'
import { rootLayout, redirectRoute } from './layout'
import unLayoutRoutes from './unLayout'
import { message } from 'ant-design-vue'
import { decryptPwd } from '@/utils/cryptoLoginSecret'
import { doCommonLogout, isValidPassword } from '@/utils/util'
import $bus from '@/utils/bus'
import { nextTick } from 'vue'
import { useKeepAliveStore } from '@/stores/keepAlive'
import { LOGIN_NAME, whiteList, REDIRECT_NAME, passedRoutes, PAGE_NOT_FOUND_NAME } from './constant';

export const routes = [
  ...unLayoutRoutes,
  rootLayout,
  redirectRoute,
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})


/**
 * 路由守卫
 */

// 是否跳转权限
function hasRoute(to, menuNames = []) {
  if (passedRoutes.includes(to.path)) return true
  return menuNames.includes(to.name)
}

router.beforeEach(async (to, from, next) => {
  Nprogress.start()
  const userStore = useUserStore();
  const keepAliveStore = useKeepAliveStore();
  // 如果进入的是 Redirect 页面，则也将离开页面的缓存清空(刷新页面的操作)
  if (to.name == REDIRECT_NAME && from.name) {
    keepAliveStore.remove(from.name);
  }
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
            doCommonLogout()
          });
        }
      }
      if (hasRoute(to, menuNames)) { // 有权限
        const toName = to.name;
        if (to.meta && to.meta.keepAlive) {
          if (toName) keepAliveStore.add(toName);
        } else {
          if (toName) keepAliveStore.remove(toName);
        }
        next()
      } else { // 没有权限
        next({ name: PAGE_NOT_FOUND_NAME })
      }
    }
  } else { // token不存在
    keepAliveStore.clear();
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
  try {
    const password = decryptPwd();
    const forceChangePwd = isValidPassword(password) ? false : true;
    nextTick(() => {
      $bus.emit('changePasswordForce', { open: forceChangePwd })
    })
  } catch (error) {
    message.error(error.message || '验证安全过程中发现异常', 2, async () => {
      // 加载失败，退出
      doCommonLogout()
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
