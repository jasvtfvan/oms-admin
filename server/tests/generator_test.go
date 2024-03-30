package tests

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid/v5"
	"github.com/jasvtfvan/oms-admin/server/utils"
)

func TestGenerateSigningKey(t *testing.T) {
	signingKey := uuid.Must(uuid.NewV4()).String()
	fmt.Println(signingKey)
}

func TestBcryptHash(t *testing.T) {
	passwordSrc := "Oms123Admin456"
	passwordTarget := utils.BcryptHash(passwordSrc)
	printStr := fmt.Sprintf("密码明文: %s\n密码密文: %s", passwordSrc, passwordTarget)
	fmt.Println(printStr)
}
