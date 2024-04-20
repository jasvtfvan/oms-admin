import { union, isEqual, unionWith } from 'lodash-es';

export const mergeArray = (arr1, arr2) => {
  return union(arr1, arr2)
}

export const mergeArrayDeep = (arr1, arr2) => {
  return unionWith(arr1, arr2, isEqual)
}

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
