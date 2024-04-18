import CryptoJS from 'crypto-js';

const AES_KEY = "12345678123456781234567812345678"

// 加密函数
export function aesEncryptCBC(plaintText, key = AES_KEY) {
  const len = key.length
  if (len != 16 && len != 24 && len != 32) {
    console.error('key 长度必须 16/24/32长度')
    return ''
  }
  const parseKey = CryptoJS.enc.Utf8.parse(key);
  const encryptedData = CryptoJS.AES.encrypt(plaintText, parseKey, {
    iv: parseKey,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  });
  return encryptedData.ciphertext.toString();
}

// 解密函数
export function aesDecryptCBC(encrypted, key = AES_KEY) {
  const len = key.length
  if (len != 16 && len != 24 && len != 32) {
    console.error('key 长度必须 16/24/32长度')
    return ''
  }
  const parseKey = CryptoJS.enc.Utf8.parse(key);
  const encryptedHexStr = CryptoJS.enc.Hex.parse(encrypted);
  const encryptedBase64Str = CryptoJS.enc.Base64.stringify(encryptedHexStr);
  const decryptedData = CryptoJS.AES.decrypt(encryptedBase64Str, parseKey, {
    iv: parseKey,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  });
  return decryptedData.toString(CryptoJS.enc.Utf8);
}
