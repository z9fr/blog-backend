package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//   https://www.bacancytechnology.com/blog/golang-jwt

func GenerateJWT(username string, email string, uuid string) (string, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["email"] = email
	claims["uuid"] = uuid
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
