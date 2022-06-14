package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) FetchuserbyUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user_exist := h.UserService.IsUsernameTaken(username)

	if !user_exist {
		h.sendErrorResponse(w, "404 user not found", fmt.Errorf("no user with that username"), 404)
		return
	}

	// @TODO
	// maybe fetch also the posts of the user
	user_details, err := h.UserService.GetUserbyUsername(username)

	if err != nil {
		h.sendErrorResponse(w, "Internal server error", err, 500)
	}

	h.sendOkResponse(w, user_details)
	return
}
