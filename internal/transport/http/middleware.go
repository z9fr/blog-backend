package http

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/context"
	"github.com/z9fr/blog-backend/internal/models"
)

// Auth Middleware - a middleware to validate user authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Authheader := r.Header["Authorization"]

		if len(Authheader) == 0 {
			sendErrorResponse(w, "Missing Header", fmt.Errorf("Authorization is required Header"))
			return
		}

		authToken := strings.Split(Authheader[0], " ")[1]

		var tokensignkey = []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return tokensignkey, nil

		})

		if err != nil {
			sendErrorResponse(w, "Auth Error", fmt.Errorf("Invalid Authencation Token Provided"))
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		var decodedToken models.AuthToken

		for key, val := range claims {
			val := fmt.Sprintf("%v", val)

			switch key {
			case "username":
				decodedToken.Username = val
			case "uuid":
				decodedToken.Uuid = val
			case "email":
				decodedToken.Email = val
			}
		}

		if !token.Valid {
			sendErrorResponse(w, "Auth Error", fmt.Errorf("Invalid Authencation Token Provided"))
			return
		}

		context.Set(r, "token", decodedToken)
		// token := context.Get(r, "token")
		next.ServeHTTP(w, r)
	})
}
