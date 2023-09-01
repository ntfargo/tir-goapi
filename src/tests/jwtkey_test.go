package tests

import (
	"testing"

	utils "github.com/ntfargo/tir-goapi/src/internal/utils"
)

func TestGenerateJWTSecretKey(t *testing.T) {
	t.Parallel()

	secretKey, err := utils.GenerateJWTSecretKey()
	if err != nil {
		t.Errorf("Error generating JWT secret key: %v", err)
	}

	if len(secretKey) == 0 {
		t.Error("Generated JWT secret key is empty")
	}
}
