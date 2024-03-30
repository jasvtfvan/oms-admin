package tests

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/jasvtfvan/oms-admin/server/utils/crypto"
)

// func TestRsa(t *testing.T) {
// 	p1, p2, _ := crypto.GenerateKeyPair()
// 	fmt.Println(p1)
// 	fmt.Println(p2)
// }

func TestAes(t *testing.T) {
	encrypted := crypto.AesEncrypt("hello world ga ga")
	fmt.Printf("%s 长度: %d\n", encrypted, len(encrypted))
	decrypted := crypto.AesDecrypt(encrypted)
	fmt.Println(decrypted)
}

func TestRsa(t *testing.T) {
	encrypted := crypto.RsaEncrypt("{\"username\":\"oms_admin\",\"password\":\"Oms123Admin456\"}")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(encrypted)))
	decrypted := crypto.RsaDecrypt(encrypted)
	fmt.Println(decrypted)
}
