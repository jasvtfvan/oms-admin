package tests

import (
	"fmt"
	"testing"

	"github.com/jasvtfvan/oms-admin/server/utils/crypto"
)

func TestRsaGenerate(t *testing.T) {
	p1, p2, _ := crypto.GenerateKeyPair()
	fmt.Println(p1)
	fmt.Println(p2)
}

func TestAesEncrypt(t *testing.T) {
	encrypted := crypto.AesEncrypt("hello world ga ga")
	fmt.Printf("%s 长度: %d\n", encrypted, len(encrypted))
}

func TestAesDecrypt(t *testing.T) {
	encrypted := "_z-3AUtYLGL4yGQiB0LjTBBQWhzPaiNp-GLDoFxSuKs"
	decrypted := crypto.AesDecrypt(encrypted)
	fmt.Println(decrypted)
}

func TestRsaEncrypt(t *testing.T) {
	encrypted := crypto.RsaEncrypt("{\"username\":\"oms_admin\",\"password\":\"Oms123Admin456\"}")
	fmt.Println(encrypted)
}

func TestRsaDecrypt(t *testing.T) {
	encrypted := "J/y9n5fjvnOgRFGSMD22yF8ORuxYl1adejHJs7TZ2rx5ZPa1zMJ+PP3sD5+8q8MmUvQhrpyNp4cgYsk/72Mb2s7e8iVcqiMiZQXTTPjMaeGJez3GrlyoXDYXZvxk1YJT6ykXlRjjpQ5uenfXmgwRf2SXw5FeCUsQcCNaf09iC7riiH9PGByHuD0XXszmEpOgf3EdQTbMcuL2HRx2TdHZ2ioSLi7hvOcraTfHEoogpCUiLuPNNNoOZhvLPVJsEyTci4eLaA0HHNBYDFmqDWpNNJVphRdlKemBsaJkqtZp0Bte6rKVjmQcz2AsIzEZb8CBFJU9idVCxGPfRE5ZL3Wmzg=="
	fmt.Println(crypto.RsaDecrypt(encrypted))
}
