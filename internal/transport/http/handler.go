package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - store pointer to our comment service
type Handler struct {
	Router *mux.Router
}

// NewHandler - return a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
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
