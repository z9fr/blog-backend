package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"

	"github.com/z9fr/blog-backend/internal/models"
	"github.com/z9fr/blog-backend/internal/user"
	"github.com/z9fr/blog-backend/internal/utils"
)

type UserResponse struct {
	UserName    string    `json:"username"`
	Email       string    `json:"email"`
	ID          string    `json:"id"`
	Description string    `json:"Description"`
	CreatedAt   time.Time `json:"CreatedAt"`
}

type UserInputCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Get User by user name
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	u, err := h.ServiceUser.GetUser(username)

	if err != nil {
		sendErrorResponse(w, "Error Fetching User", err)
		return
	}

	if err := sendOkResponse(w, UserResponse{
		UserName:    u.UserName,
		Email:       u.Email,
		ID:          u.ID,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
	}); err != nil {
		log.Error(err)
	}
}

// create user  - add a user
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	/*
	   using the user service validate if the user is already availible
	*/
	var userinput user.User
	if err := json.NewDecoder(r.Body).Decode(&userinput); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	// check if username is taken
	userCheck, err := h.ServiceUser.GetUser(userinput.UserName)

	if userCheck.UserName == userinput.UserName {
		sendErrorResponse(w, "Failed to Create the User", fmt.Errorf("User Name Alreay Taken"))
		return
	}

	createdUser, err := h.ServiceUser.CreateUser(userinput)

	if err != nil {
		sendErrorResponse(w, "Failed to Create the User", err)
		return
	}

	if err := sendOkResponse(w, UserResponse{
		UserName:    createdUser.UserName,
		Email:       createdUser.Email,
		ID:          createdUser.ID,
		Description: createdUser.Description,
		CreatedAt:   createdUser.CreatedAt,
	}); err != nil {
		log.Error(err)
	}
}

// Authenticate User
func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var creds UserInputCreds

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	actualUser, err := h.ServiceUser.GetUser(creds.Username)

	if err != nil {
		sendErrorResponse(w, "Error Fetching User", err)
		return
	}

	math := utils.CheckPasswordHash(creds.Password, actualUser.Password)

	if !math {
		sendErrorResponse(w, "Authentication Failure", fmt.Errorf("Invalid username or password"))
		return
	}

	authtoken, err := utils.GenerateJWT(actualUser.UserName, actualUser.Email, actualUser.ID)

	if err != nil {
		sendErrorResponse(w, "Internal Server Error", fmt.Errorf("Failed to generate auth token %w", err))
		return
	}

	w.Header().Set("Authorization", "bearer "+authtoken)

	if err := sendOkResponse(w, authtoken); err != nil {
		log.Error(err)
	}

}

// Current user - display information about the current /logged in user

func (h *Handler) CurrentUser(w http.ResponseWriter, r *http.Request) {

	token := context.Get(r, "token")
	username := token.(models.AuthToken).Username

	u, err := h.ServiceUser.GetUser(username)

	if err != nil {
		sendErrorResponse(w, "Failed to Fetch User", err)
		return
	}

	if err := sendOkResponse(w, UserResponse{
		UserName:    u.UserName,
		Email:       u.Email,
		ID:          u.ID,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
	}); err != nil {
		log.Error(err)
	}
}
