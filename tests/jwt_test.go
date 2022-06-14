package utils

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/utils"
)

var ApplicationSecret string

func init() {
	secret, err := utils.SecretGenerator(100)

	if err != nil {
		logrus.Panic("Unable to generate the secret", err)
	}

	ApplicationSecret = secret
}

func TestGenerateJWT(t *testing.T) {

	token, err := utils.GenerateJWT("username", "email", ApplicationSecret)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(fmt.Sprintf("Token Secret -> %s\nJWT token -> %s", ApplicationSecret, token))
}
