package tests

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid/v5"
)

func TestGenerateSigningKey(t *testing.T) {
	signingKey := uuid.Must(uuid.NewV4()).String()
	fmt.Println(signingKey)
}
