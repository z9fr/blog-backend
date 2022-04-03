package http

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

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
