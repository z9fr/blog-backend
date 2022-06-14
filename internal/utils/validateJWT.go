package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/z9fr/blog-backend/internal/types"
)

func VerifyToken(tokenString string, appsecret string) (types.TokenDetails, error) {
	// TokenDetails

	var hmacSampleSecret = []byte(appsecret)

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		return types.TokenDetails{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user types.TokenDetails
		mapstructure.Decode(claims["user"], &user)
		return user, nil
	} else {
		return types.TokenDetails{}, err
	}

}
