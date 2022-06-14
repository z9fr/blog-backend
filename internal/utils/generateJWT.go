package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//   https://www.bacancytechnology.com/blog/golang-jwt

func GenerateJWT(username string, email string, appsecret string) (string, error) {

	var mySigningKey = []byte(appsecret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
