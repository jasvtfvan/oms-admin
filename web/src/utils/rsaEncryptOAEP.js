// 后端固定的公钥
const publicKeyPem = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7TQDqigwjCPWURt/BuAS
37NMQ3U2uBGpf5gBzNlyZjp96lkUvF93BqItt9UAAtNUVh4rjOgmIWX3xcCnu6ET
B8bho6kfbBy8qoTb9lTekz6Y2WHX3eydLozFxC8h1exZ6V3kYoiztUkChz6v+kwN
s6kGospt8r0rax81Eav3RgajJmpd4OXMOpVX4lNdRLOgzK9jUOmlaVQXJDHRbyzW
oC1ooNguKIOpGWQ5doKD+tznUwhisjTxmuW/tNbWDd2nPl/pTjYIY2RE9xigtOov
AHfY3AUY7NN+mGZn9UxTI2jjdzOQTB9un7izQvOB+otEF8Wb67dtHSFlDB7W8TbQ
VQIDAQAB
-----END PUBLIC KEY-----`;

// 将后端使用的标签 "oms" 转换为 Uint8Array  
const label = new TextEncoder().encode("oms");

// 将PEM格式的公钥转换为适合Web Crypto API使用的SPKI格式  
async function pemToSpki(pem) {
  const pemHeader = "-----BEGIN PUBLIC KEY-----";
  const pemFooter = "-----END PUBLIC KEY-----";
  const pemContents = pem.substring(pemHeader.length, pem.length - pemFooter.length);
  const binaryDerString = window.atob(pemContents);
  const binaryDer = new Uint8Array(binaryDerString.length);
  for (let i = 0; i < binaryDerString.length; i++) {
    binaryDer[i] = binaryDerString.charCodeAt(i);
  }
  return window.crypto.subtle.importKey(
    "spki",
    binaryDer,
    {
      name: "RSA-OAEP",
      hash: "SHA-256" // 使用与后端一致的哈希算法  
    },
    false,
    ["encrypt"]
  );
}

// 使用公钥加密数据  
async function getEncryptedBase64(publicKey, message) {
  const encoder = new TextEncoder();
  const data = encoder.encode(message);
  const encryptedData = await window.crypto.subtle.encrypt(
    {
      name: "RSA-OAEP",
      label: label // OAEP使用的标签，与后端保持一致  
    },
    publicKey,
    data
  );
  // 将ArrayBuffer转换为Array并返回  
  const encryptedArray = Array.from(new Uint8Array(encryptedData));
  // 如果你需要Base64格式的字符串，可以转换它  
  const encryptedBase64 = btoa(String.fromCharCode.apply(null, encryptedArray));
  return encryptedBase64;
}

// 调用函数进行加密  
async function encryptMessage(pem, message) {
  try {
    const publicKey = await pemToSpki(pem);
    const encrypted = await getEncryptedBase64(publicKey, message);
    return encrypted
  } catch (error) {
    console.error("Encryption failed:", error);
  }
}

export function rsaEncrypt(data, publicKeyStr) {
  return encryptMessage(publicKeyStr || publicKeyPem, data) || "";
}
