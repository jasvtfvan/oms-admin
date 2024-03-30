package internal

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// 加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(cipherText []byte, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}

// 生成密钥对
func GetRsaKeyPair() (string, string, error) {
	var privateKeyResult string
	var publicKeyResult string

	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		errStr := fmt.Sprintf("私钥生成失败: %s", err.Error())
		return "", "", errors.New(errStr)
	}
	// 将私钥编码为ASN.1 DER格式
	privateKeyASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
	// 使用PEM格式对私钥进行编码
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyASN1,
	}
	// 将私钥PEM编码写入缓冲区
	privateKeyBuffer := pem.EncodeToMemory(privateKeyPEM)
	// 私钥结果
	privateKeyResult = string(privateKeyBuffer)

	// 提取公钥
	publicKey := &privateKey.PublicKey
	// 将公钥编码为ASN.1 DER格式
	publicKeyASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		errStr := fmt.Sprintf("公钥生成失败: %s", err.Error())
		return "", "", errors.New(errStr)
	}
	// 使用PEM格式对公钥进行编码
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyASN1,
	}
	// 将公钥PEM编码写入缓冲区
	publicKeyBuffer := pem.EncodeToMemory(publicKeyPEM)
	// 公钥结果
	publicKeyResult = string(publicKeyBuffer)

	return privateKeyResult, publicKeyResult, nil
}
