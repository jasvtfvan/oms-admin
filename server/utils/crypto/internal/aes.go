package internal

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

func AesEncryptCBC(orig string, key string) (string, error) {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		errStr := fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error())
		return "", errors.New(errStr)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	encrypted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(encrypted, origData)
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	return base64.RawURLEncoding.EncodeToString(encrypted), nil
}

func AesDecryptCBC(encrypted string, key string) (string, error) {
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	encryptedByte, _ := base64.RawURLEncoding.DecodeString(encrypted)
	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		errStr := fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error())
		return "", errors.New(errStr)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(encryptedByte))
	// 解密
	blockMode.CryptBlocks(orig, encryptedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig), nil
}

// 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
