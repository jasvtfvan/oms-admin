import { aesEncryptCBC, aesDecryptCBC } from '@/utils/aesCrypto'
import { localCache } from '@/utils/cache'

/**
 * 登录secret加密/解密
 */
const SECRET_AES_KEY = 'abcdefgh12345678abcdefgh12345678'
export const encryptSecret = (secret) => {
  return aesEncryptCBC(secret, SECRET_AES_KEY)
}
export const decryptSecret = (encryptedSecret) => {
  return aesDecryptCBC(encryptedSecret, SECRET_AES_KEY)
}

/**
 * 登录成功的密码加密/解密/对比secret
 */
// xor异或方法，最简单的加密方式
const XOR_KEY = '_yes_is_your_secret_'
function xorEncryptDecrypt(input, key = XOR_KEY) {
  let output = '';
  for (let i = 0; i < input.length; i++) {
    output += String.fromCharCode(input.charCodeAt(i) ^ key.charCodeAt(i % key.length));
  }
  return output;
}
// 扰乱真实密码方法
function adjustStringLength(str) {
  if (str.length < 6) {
    while (str.length < 6) {
      str += 'a';
    }
  } else if (str.length > 6) {
    str = str.substring(0, 6);
  }
  return str;
}
function disturbPwd(password, disturbStr = '123456') {
  if (password.length < 6) throw new Error('密码长度必须至少为6位');
  const arr = disturbStr.split('');
  // 在第1个位置插入'se'  
  let modifiedPassword = arr[0] + arr[1] + password.slice(0);
  // 在第3个位置插入'cre'（现在已经是第3位，因为前面插入了2位）  
  modifiedPassword = modifiedPassword.slice(0, 3) + arr[2] + arr[3] + arr[4] + modifiedPassword.slice(3);
  // 在第5个位置插入't'（现在已经是第7位，因为前面插入了5位）  
  modifiedPassword = modifiedPassword.slice(0, 5) + arr[5] + modifiedPassword.slice(5);
  return modifiedPassword;
}
function resetPwd(modifiedPassword, bits) {
  // 移除第7个位置的't'
  let restoredPassword = modifiedPassword.slice(0, 6) + modifiedPassword.slice(bits[0]);
  // 移除第4个位置的'cre'（现在已经是第3位，因为前面移除了1位）
  restoredPassword = restoredPassword.slice(0, 3) + restoredPassword.slice(bits[1]);
  // 移除前2个位置的'se'
  restoredPassword = restoredPassword.slice(bits[2]);
  return restoredPassword;
}

const PWD_TIMESTAMP_AES_KEY = '8765432187654321'
const PWD_OBJECT_AES_KEY = '87l6x543z2i1d8y765o432q1'
export const encryptPwd = (password) => {
  const timeStr = `${new Date().getTime()}` // 时间戳
  const xorTimeStr = xorEncryptDecrypt(`${timeStr}_762`) // xor时间戳
  const encryptXorTime = aesEncryptCBC(xorTimeStr, PWD_TIMESTAMP_AES_KEY) // 将xor时间戳aes加密
  localCache.set(`_${xorTimeStr}_is_not_your_secret`, encryptXorTime) // key:xor时间戳,value:xor+aes加密
  const disturbStr = adjustStringLength(xorTimeStr) // 根据xor时间戳获取6位扰乱key
  const disturbedPwd = disturbPwd(password, disturbStr) // 扰乱原始密码
  const xorEncryptXorTime = xorEncryptDecrypt(encryptXorTime) // 将aes+xor字符串再次xor
  const jsonObj = JSON.stringify({ tt_tt: xorEncryptXorTime, pp_pp: disturbedPwd }) // 转成json字符串
  const encryptedJson = aesEncryptCBC(jsonObj, PWD_OBJECT_AES_KEY) // json进行aes加密
  localCache.set('tp_tp_ss', encryptedJson) // 放到localCache
}
export const decryptPwd = () => {
  const encryptedJson = localCache.get('tp_tp_ss') // 通过localCache获取
  if (!encryptedJson) throw new Error('验证安全过程中发现异常');
  const jsonStr = aesDecryptCBC(encryptedJson, PWD_OBJECT_AES_KEY) // aes解密
  const jsonObj = JSON.parse(jsonStr) // 解密后转成json
  const xorEncryptXorTime = jsonObj.tt_tt // 拿到xor+aes
  if (!xorEncryptXorTime) throw new Error('验证安全过程中发现异常');
  const encryptXorTime = xorEncryptDecrypt(xorEncryptXorTime) // 接触第一层xor
  const xorTimeStr = aesDecryptCBC(encryptXorTime, PWD_TIMESTAMP_AES_KEY) // aes解密为xor
  const encryptXorTime1 = localCache.get(`_${xorTimeStr}_is_not_your_secret`) // 根据xor获取xor+aes
  if (!encryptXorTime1 || encryptXorTime != encryptXorTime1) throw new Error('验证安全过程中发现异常');
  const timeStrWithBits = xorEncryptDecrypt(xorTimeStr) // 获取时间戳str
  const bits = timeStrWithBits.split('_')[1].split('').map(Number) // 通过下划线后边获取bits数组
  const disturbedPwd = jsonObj.pp_pp // 拿到扰乱的pwd
  const answer_pwd = resetPwd(disturbedPwd, bits) // 拿到真实pwd
  return answer_pwd
}
export const clearPwd = () => {
  const encryptedJson = localCache.get('tp_tp_ss') // 通过localCache获取
  if (!encryptedJson) return;
  const jsonStr = aesDecryptCBC(encryptedJson, PWD_OBJECT_AES_KEY) // aes解密
  const jsonObj = JSON.parse(jsonStr) // 解密后转成json
  const xorEncryptXorTime = jsonObj.tt_tt // 拿到xor+aes
  if (!xorEncryptXorTime) return
  const encryptXorTime = xorEncryptDecrypt(xorEncryptXorTime) // 接触第一层xor
  const xorTimeStr = aesDecryptCBC(encryptXorTime, PWD_TIMESTAMP_AES_KEY) // aes解密为xor
  localCache.remove(`_${xorTimeStr}_is_not_your_secret`)
  localCache.remove('tp_tp_ss')
}
