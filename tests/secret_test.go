package utils

import (
	"fmt"
	"testing"

	"github.com/z9fr/blog-backend/internal/utils"
)

func TestSecretGeneration(t *testing.T) {
	secret, err := utils.SecretGenerator(100)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(fmt.Sprintf("Generated secret -> %s", secret))

}
