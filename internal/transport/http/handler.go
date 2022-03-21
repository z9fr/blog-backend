package http

import (
	"api/internal/comment"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler - store pointer to our comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// NewHandler - return a pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRotues() {
	log.Printf("Setting up routes")

	// initicate new gorilla mox router
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API running ok ")
	})

}

// GetComment - Retriew comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	// get id from user parameter
	vars := mux.Vars(r)
	id := vars["id"] // this is a string but id is uint so we need to convert

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Fprint(w, "Enable to parse UINT from ID ")
	}

	comment, err := h.Service.GetComment(uint(i))

	if err != nil {
		fmt.Fprint(w, "Error Fetching Comment")
	}

	fmt.Fprint(w, comment)
}
