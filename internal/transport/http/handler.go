package http

import (
	"api/internal/comment"
	"encoding/json"
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

// Response - an object to store responses from our api
type Response struct {
	Message string
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
	h.Router.HandleFunc("/api/comment/update/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{
			Message: "Api is Running OK",
		}); err != nil {
			log.Panic(err)
		}
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		log.Panic(err)
	}
}

// GetAllComments - retriews all comments from the database
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Fprint(w, err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		log.Panic(err)
	}
}

// PostComment - add a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
  var comment comment.Comment
  // getting the comment from request body
  if err := json.NewDecoder(r.Body).Decode(&comment); err != nil{
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, "Failed to decode JSON body")
  }
  // saving the comment on the database
	comment, err := h.Service.PostComment(comment)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err != nil {
		fmt.Fprint(w, err)
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		log.Panic(err)
	}
}

// UpdateComment - Update comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)

  // parsting of the request body
  var newComment comment.Comment
  if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil{
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, "Failed to Decode JSON body")
  }
	
  comment, err := h.Service.UpdateComment(uint(commentId), newComment)

	if err != nil {
    w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error updating the comment")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		log.Panic(err)
	}
}

// DeleteComment - Delete a comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		fmt.Fprint(w, "Enable to parse UINT from ID ")
	}

	err = h.Service.DeleteComment(uint(commentId))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(Response{
			Message: "Failed to delete comment",
		}); err != nil {
			log.Panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{
		Message: "Successfully deleted the comment",
	}); err != nil {
		log.Panic(err)
	}
}
