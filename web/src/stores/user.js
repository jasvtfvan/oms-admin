import { ref } from 'vue'
import { defineStore } from 'pinia'
import localCache from '@/utils/localCache'
import sessionCache from '@/utils/sessionCache'
import { postLogin, postLogout } from '@/api/common/user'

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

  // 登录
  const Login = async (data) => {
    return new Promise((resolve, reject) => {
      postLogin(data)
        .then(res => {
          let { user, token } = res.data;
          if (!token) token = '';
          setToken(token);
          if (!user) user = {};
          setUserInfo(user);
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
        removeToken();
        removeGroup();
        removeUserInfo();
        resolve();
      }).catch(() => {
        resolve();
      });
    });
  }

  return {
    token, // 小写开头，getter
    group, // 小写开头，getter
    userInfo, // 小写开头，getter
    SetGroup, // 大写Set开头，对外暴露的action
    Login, // 大写开头，对外暴露的action
    Logout, // 大写开头，对外暴露的action
  }
})
