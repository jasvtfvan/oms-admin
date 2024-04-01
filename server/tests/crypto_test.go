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
	encrypted := "CyI8+wQr8OJyoWTpK3y9WzbTgFFSPIzi5We4cULB0gKGoNYz4wdEQ7Py/GzCwP435ZJ2OB0joOL3TBvdfFr5RsA7WB+7zewYeopiuieZTY1T6ouq0bwn1V/F8axa6CzBwqYByr5hkjXob7Ek0huYooGX1P4cM0N/Kp+Mq75UvlFmeukZ6znlshNwTPGfw0GCzc4TUBdA7zZwkG7iAvwvsp44H8KA6yQ8YGfwzlyyAFQ1c4W1OuLKDsU8mdWrJml1xfuIiTD5dMZv5HPMaYjoewTiVWhIDEDe7x9LawLvXWNX+5ZN5cwwBjRVAzZQWIukILAXU0fHRvINkqEzjr98Yg=="
	fmt.Println(crypto.RsaDecrypt(encrypted))
}
