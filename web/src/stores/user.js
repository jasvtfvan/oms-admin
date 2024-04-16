import { ref } from 'vue'
import { defineStore } from 'pinia'
import { localCache, sessionCache } from '@/utils/cache'
import { postLogin, postLogout, postCaptcha } from '@/api/common/base'
import { getMenus } from '@/api/common/user'
import { rsaEncrypt } from '@/utils/rsaEncryptOAEP'
import { addDynamicRoutes } from '@/router/helper/routeHelper'

export const useUserStore = defineStore('user', () => {
  // token登录凭证
  const token = ref(localCache.get('token') || '')
  const setToken = (val) => {
    localCache.set('token', val)
    token.value = val
  }
  const removeToken = () => {
    localCache.remove('token')
    token.value = ''
  }
  // group当前选中的组织
  const group = ref(sessionCache.get('group') || '')
  // 大写开头action，对外暴露
  const SetGroup = (val) => {
    sessionCache.set('group', val)
    group.value = val
  }
  // 小写开头action，不对外暴露
  const removeGroup = () => {
    sessionCache.remove('group')
    group.value = ''
  }
  // userInfo用户信息
  const userInfo = ref(sessionCache.get('userInfo') || {})
  const setUserInfo = (val) => {
    sessionCache.set('userInfo', val)
    userInfo.value = val
  }
  const removeUserInfo = () => {
    sessionCache.remove('userInfo')
    userInfo.value = {}
  }
  // 菜单
  const menus = ref([])
  const setMenus = (val) => {
    sessionCache.set('menus', val)
    menus.value = val
  }
  const removeMenus = () => {
    sessionCache.remove('menus')
    menus.value = []
  }

  // 清空当前登录状态(userInfo,token,....)
  const ClearLoginStatus = () => {
    removeGroup();
    removeMenus();
    removeToken();
    removeUserInfo();
  };
  // 获取菜单
  const GetMenus = async () => {
    try {
      const res = await getMenus();
      const menus = res.data || [];
      addDynamicRoutes(menus);
      setMenus(menus);
      return menus;
    } catch (error) {
      addDynamicRoutes();
      return error
    }
  };
  // 登录
  const Login = async (data) => {
    const { username, password, captcha, captchaId } = data
    const secret = await rsaEncrypt(JSON.stringify({ username, password }))
    return new Promise((resolve, reject) => {
      postLogin({ secret, captcha, captchaId })
        .then(res => {
          let { user, token } = res.data;
          // token登录凭证
          if (!token) token = '';
          setToken(token);
          // user用户信息
          if (!user) user = {};
          setUserInfo(user);
          // group默认群组
          let sysGroups = user.sysGroups || [];
          if (sysGroups[0]) {
            SetGroup(sysGroups[0].orgCode);
          }
          resolve(user);
        })
        .catch(error => {
          reject(error);
        });
    });
  }
  // 退出
  const Logout = async () => {
    return new Promise((resolve) => {
      postLogout().then(() => {
        ClearLoginStatus();
        resolve();
      }).catch(() => {
        resolve();
      });
    });
  }
  // 获取验证码
  const Captcha = async (data) => {
    return new Promise((resolve, reject) => {
      postCaptcha(data)
        .then(res => {
          const {
            captchaId,
            picPath,
            captchaLength,
            openCaptcha,
          } = res.data;
          resolve({
            captchaId,
            picPath,
            captchaLength,
            openCaptcha,
          });
        })
        .catch(error => {
          reject(error);
        });
    });
  }

  return {
    group,
    menus,
    token,
    userInfo, // 小写开头，getter
    Captcha,
    ClearLoginStatus,
    GetMenus,
    Login,
    Logout,
    SetGroup, // 大写开头，对外暴露的action
  }
})
