import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { localCache, sessionCache } from '@/utils/cache'
import { postLogin, postLogout, postCaptcha } from '@/api/common/base'
import { getMenus, getUserProfile } from '@/api/common/user'
import { addDynamicRoutes, getAdminMenuNames, getLayoutMenus } from '@/router/helper/routeHelper'
import { mergeArray } from '@/utils/util'
import { decryptSecret, encryptSecret } from '@/utils/cryptoLoginSecret'

export const useUserStore = defineStore('user', () => {
  // token登录凭证
  const token = ref(localCache.get('token') || '')
  const setToken = (val = '') => {
    localCache.set('token', val)
    token.value = val
  }
  const removeToken = () => {
    localCache.remove('token')
    token.value = ''
  }
  // 登录成功后的账号密码加密
  const encryptedSecret = ref(localCache.get('encryptedSecret') || '')
  const setEncryptedSecret = (val = '') => {
    localCache.set('encryptedSecret', val)
    encryptedSecret.value = val
  }
  const removeEncryptedSecret = () => {
    localCache.remove('encryptedSecret')
    encryptedSecret.value = ''
  }
  // group当前选中的组织
  const group = ref(sessionCache.get('group') || '')
  // 大写开头action，对外暴露
  const SetGroup = (val = '') => {
    sessionCache.set('group', val)
    group.value = val
  }
  // 小写开头action，不对外暴露
  const removeGroup = () => {
    sessionCache.remove('group')
    group.value = ''
  }
  // userProfile用户信息
  const userProfile = ref(sessionCache.get('userProfile') || {})
  const setUserProfile = (val = {}) => {
    sessionCache.set('userProfile', val)
    userProfile.value = val
  }
  const removeUserProfile = () => {
    sessionCache.remove('userProfile')
    userProfile.value = {}
  }
  // 菜单，没有权限则隐藏，有权限原来隐藏的依然隐藏
  const menus = ref(sessionCache.get('menus') || [])
  const setMenus = (val = []) => {
    sessionCache.set('menus', val)
    menus.value = val
  }
  const removeMenus = () => {
    sessionCache.remove('menus')
    menus.value = []
  }
  // 菜单名，包含则有权访问路由，跟隐藏无关
  const menuNames = ref(sessionCache.get('menuNames') || [])
  const setMenuNames = (val = []) => {
    sessionCache.set('menuNames', val)
    menuNames.value = val
  }
  const removeMenuNames = () => {
    sessionCache.remove('menuNames')
    menuNames.value = []
  }
  // 动态路由是否ready
  const dynamicRoutesReady = ref(false)
  const setDynamicRoutesReady = (val) => {
    console.log('setDynamicRoutesReady', !!val)
    dynamicRoutesReady.value = !!val
  }

  const groups = computed(() => {
    return userProfile.value.sysGroups || [];
  })
  // 是否超级管理员
  const isRootAdmin = computed(() => {
    return userProfile.value.isRootAdmin
  })
  // 根据group，判断是否为管理员
  const isAdmin = computed(() => {
    const foundGroup = groups.value.find(grp => grp.orgCode == group);
    if (foundGroup && foundGroup.sysRoles && foundGroup.sysRoles.length) {
      return foundGroup.sysRoles.some(role => role.isAdmin == true);
    }
    return false;
  })

  // 清空当前登录状态(userProfile,token,....)
  const ClearLoginStatus = () => {
    setDynamicRoutesReady(false)
    removeEncryptedSecret();
    removeGroup();
    removeMenus();
    removeMenuNames();
    removeToken();
    removeUserProfile();
  };
  // 重新获取权限
  const RefreshAuth = async () => {
    try {
      const secret = decryptSecret(encryptedSecret.value)
      await Login({ secret, captcha: '', captchaId: '' })
      await GetAuthWithoutLogin()
    } catch (error) {
      return Promise.reject(error)
    }
  }
  // 获取权限（不包含token）
  const GetAuthWithoutLogin = async () => {
    try {
      const profile = await GetUserProfile()
      const { menus, menuNames } = await GetMenus()
      return { menus, menuNames, profile }
    } catch (error) {
      return Promise.reject(error)
    }
  }
  // 添加动态路由
  const AddDynamicRoutes = async (menuNames) => {
    if (!dynamicRoutesReady.value) {
      addDynamicRoutes(menuNames); // 添加动态路由
      setDynamicRoutesReady(true);
    }
  }
  // 获取菜单
  const GetMenus = async () => {
    let menuNames = [];
    let adminMenuNames = [];
    if (isRootAdmin.value || isAdmin.value) {
      adminMenuNames = getAdminMenuNames();
    }
    try {
      const res = await getMenus();
      menuNames = res.data || []; // 不包含/home等默认路由
      menuNames = mergeArray(adminMenuNames, menuNames);
      setMenuNames(menuNames);
      const menus = getLayoutMenus(menuNames) // 包含/home等默认路由（去掉component引入）
      setMenus(menus)
      return { menus, menuNames }
    } catch (error) {
      return Promise.reject(error)
    }
  };
  // 获取用户信息
  const GetUserProfile = async () => {
    return new Promise((resolve, reject) => {
      getUserProfile()
        .then(res => {
          const profile = res.data || {};
          setUserProfile(profile);
          // group默认群组
          let sysGroups = profile.sysGroups || [];
          if (sysGroups[0]) {
            SetGroup(sysGroups[0].orgCode);
          }
          resolve(profile);
        })
        .catch(error => {
          reject(error);
        });
    })
  }
  // 登录
  const Login = async (data) => {
    const { secret, captcha, captchaId } = data
    return new Promise((resolve, reject) => {
      postLogin({ secret, captcha, captchaId })
        .then(res => {
          let token = res.data;
          if (!token) token = '';
          setToken(token);
          setEncryptedSecret(encryptSecret(secret)); // 把账号密码加密保存起来
          resolve(token);
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
        ClearLoginStatus();
        console.warn('调用退出接口失败，直接resolve')
        resolve();
      })
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
    encryptedSecret,
    group,
    groups,
    isAdmin,
    isRootAdmin,
    menus,
    menuNames,
    token,
    userProfile, // 小写开头，getter
    AddDynamicRoutes,
    Captcha,
    ClearLoginStatus,
    GetAuthWithoutLogin,
    GetMenus,
    GetUserProfile,
    Login,
    Logout,
    RefreshAuth,
    SetGroup, // 大写开头，对外暴露的action
  }
})
