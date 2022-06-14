package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/z9fr/blog-backend/internal/types"
)

//   https://www.bacancytechnology.com/blog/golang-jwt

func GenerateJWT(username string, email string, appsecret string) (string, error) {

	var mySigningKey = []byte(appsecret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = types.TokenDetails{
		UserName: username,
		Email:    email,
	}
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
