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
	encrypted := crypto.AesEncrypt("{\"initPwd\":\"Oms123Admin456\"}")
	fmt.Printf("%s 长度: %d\n", encrypted, len(encrypted))
}

func TestAesDecrypt(t *testing.T) {
	encrypted := "6794879eafffe1feea550af7c9fd7b7f9286e73ed23c634832c6c7246d5a9fb3"
	decrypted := crypto.AesDecrypt(encrypted)
	fmt.Println(decrypted)
}

func TestRsaEncrypt(t *testing.T) {
	encrypted := crypto.RsaEncrypt("{\"username\":\"oms_admin\",\"password\":\"Oms123Admin456\"}")
	fmt.Println(encrypted)
}

func TestRsaDecrypt(t *testing.T) {
	encrypted := "osRU4hYZ1pg4LUL3iC52ezD/HasiIzDnadzbY76nRlLaeHTYVVk12RlrlJewvR1h0MJiUql03p79OTEUw/93H4rM1hVyzdEe7bL0xFXCGKpVJrM6oKMlQjq7Qh3I3xtkzeGwMwCT70a9hx4/2rQypGjmGeKgo8n7UeY/g3futP6152KBmVcWvgmS38NG7LKcQbbEsDr7FegNP0uuUmRn+p6iIxmuWKh4oZa3adgzQwWfZm5cSCOtjdyCNDD9mkP6/KKQYuO4xAkj5icR4rwGgPjBTTwTc63ibjdZlRYXXWUE0HEZWTuilSaZtxPnEq0todFe95AxTg2noe/aNyg2Mw=="
	fmt.Println(crypto.RsaDecrypt(encrypted))
}
