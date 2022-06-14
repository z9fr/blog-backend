package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z9fr/blog-backend/internal/user"
	"github.com/z9fr/blog-backend/internal/utils"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) FetchuserbyUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user_exist := h.UserService.IsUsernameTaken(username)

	if !user_exist {
		h.sendErrorResponse(w, "404 user not found", fmt.Errorf("no user with that username"), http.StatusNotFound)
		return
	}

	// @TODO
	// maybe fetch also the posts of the user
	user_details, err := h.UserService.GetUserbyUsername(username)

	if err != nil {
		h.sendErrorResponse(w, "Internal server error", err, http.StatusInternalServerError)
	}

	h.sendOkResponse(w, user_details)
	return
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput user.User

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		h.sendErrorResponse(w, "Failed to decode the JSON body", err, http.StatusInternalServerError)
		return
	}

	// check if the username or password taken

	userMailExist := h.UserService.IsEmailTaken(userInput.Email)
	userUsernameExist := h.UserService.IsEmailTaken(userInput.UserName)

	if userMailExist {
		h.sendErrorResponse(w, "Email is already taken.", fmt.Errorf("Email is taken."), http.StatusConflict)
		return
	}

	if userUsernameExist {
		h.sendErrorResponse(w, "Username is already taken.", fmt.Errorf("username is taken."), http.StatusConflict)
		return
	}

	// save the user on db
	createdUser, err := h.UserService.CreateUser(userInput)

	if err != nil {
		h.sendErrorResponse(w, "Unable to create a user please try again", err, http.StatusInternalServerError)
		return
	}

	h.sendOkResponse(w, struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}{
		Username: createdUser.UserName,
		Email:    createdUser.Email,
	})

}

// Authenticate User
func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var creds LoginReq

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		h.sendErrorResponse(w, "Failed to decode JSON body", err, http.StatusInternalServerError)
		return
	}

	actualUser, err := h.UserService.GetUserbyEmail(creds.Email)

	if err != nil {
		h.sendErrorResponse(w, "Login Failed", err, http.StatusUnauthorized)
		return
	}

	isPasswordCorrect := utils.CheckPasswordHash(creds.Password, actualUser.Password)

	if !isPasswordCorrect {
		h.sendErrorResponse(w, "Authenticate Failed", fmt.Errorf("Invalid email or password"), http.StatusUnauthorized)
		return
	}

	authtoken, err := utils.GenerateJWT(actualUser.UserName, actualUser.Email, h.ApplicationSecret)

	if err != nil {
		h.sendErrorResponse(w, "Inernal server error", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", authtoken))

	h.sendOkResponse(w, struct {
		AuthToken string `json:"auth_token"`
	}{
		AuthToken: authtoken,
	})
}
