import { union, isEqual, unionWith } from 'lodash-es';
import { useUserStore } from '@/stores/user';
import { clearPwd } from './cryptoLoginSecret';

export const doCommonLogout = async () => {
  const userStore = useUserStore()
  await userStore.Logout()
  clearPwd()
  // window.location.href = url;
  // 为了重新实例化vue-router对象 避免bug
  window.location.reload()
}

// 合并2个数组，去掉重复项
export const mergeArray = (arr1, arr2) => {
  return union(arr1, arr2)
}

// 深度合并2个数组，去掉重复项
export const mergeArrayDeep = (arr1, arr2) => {
  return unionWith(arr1, arr2, isEqual)
}

// 判断密码是否符合规则
export const passwordErrorMessage = '密码长度不能小于8，必须同时包含数字、大写字母和小写字母'
export const isValidPassword = (password) => {
  // 检查长度
  if (password.length < 8) {
    return false;
  }
  // 检查是否包含小写字母
  if (!/[a-z]/.test(password)) {
    return false;
  }
  // 检查是否包含大写字母
  if (!/[A-Z]/.test(password)) {
    return false;
  }
  // 检查是否包含数字
  if (!/\d/.test(password)) {
    return false;
  }
  return true;
}
