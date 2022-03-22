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

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/create", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/delete/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/update", h.UpdateComment).Methods("PUT")
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

// GetAllComments - retriews all comments from the database
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, comments)
}

// PostComment - add a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	// Just for testing
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/test",
	})

	if err != nil {
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, comment)
}

// UpdateComment - Update comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/updated",
	})

	if err != nil {
		fmt.Fprint(w, "Error updating the comment")
	}
	fmt.Fprint(w, "updated", comment)
}

// DeleteComment - Delete a comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Fprint(w, "Enable to parse UINT from ID ")
	}

	err = h.Service.DeleteComment(uint(commentId))

	if err != nil {
		fmt.Fprint(w, "Failed to delete comment")
	}

	fmt.Fprint(w, "Successfully deleted the comment")

}
