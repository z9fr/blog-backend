package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/z9fr/blog-backend/internal/user"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type UserResponse struct {
	UserName    string `json:"username"`
	Email       string `json:"email"`
	ID          string `json:"id"`
	Description string `json:"Description"`
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
	}); err != nil {
		log.Error(err)
	}
}
