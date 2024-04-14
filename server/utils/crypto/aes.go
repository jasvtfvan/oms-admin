package crypto

import "github.com/jasvtfvan/oms-admin/server/utils/crypto/internal"

var aesKey string = "12345678123456781234567812345678"

func AesEncrypt(orig string) string {
	encrypted, _ := internal.AesEncryptCBC(orig, aesKey)
	return encrypted
}

func AesDecrypt(encrypted string) string {
	decrypted, _ := internal.AesDecryptCBC(encrypted, aesKey)
	return decrypted
}
