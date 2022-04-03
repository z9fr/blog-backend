package http

import (
	"encoding/json"
	"net/http"

	"github.com/z9fr/blog-backend/internal/user"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Get User by user name
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := h.ServiceUser.GetUser(username)

	if err != nil {
		sendErrorResponse(w, "Error Fetching User", err)
		return
	}

	if err := sendOkResponse(w, user); err != nil {
		log.Error(err)
	}
}

// create user  - add a user
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	/*
	   using the user service validate if the user is already availible
	*/
	var user user.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	createdUser, err := h.ServiceUser.CreateUser(user)

	if err != nil {
		sendErrorResponse(w, "Failed to Create the User", err)
		return
	}

	if err := sendOkResponse(w, createdUser); err != nil {
		log.Error(err)
	}
}
