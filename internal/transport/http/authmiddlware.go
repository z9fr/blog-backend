package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/utils"
)

func (h *Handler) JWTMiddlewhare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authheader := r.Header["Authorization"]

		if len(authheader) == 0 {
			h.sendErrorResponse(w, "Missing Authorization Header", fmt.Errorf("Authorization is required Header"), http.StatusUnauthorized)
			return
		}

		authToken := strings.Split(authheader[0], " ")[1]
		AuthValeus, err := utils.VerifyToken(authToken, h.ApplicationSecret)

		user, err := h.UserService.GetUserbyEmail(AuthValeus.Email)

		if err != nil {
			logrus.Warn(r, err)
			h.sendErrorResponse(w, "Unable to Verify JWT Token", err, http.StatusUnauthorized)
			return
		}

		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
