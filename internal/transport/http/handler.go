package http

import (
	"api/internal/comment"
	"encoding/json"
	log "github.com/sirupsen/logrus"

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
	Error   string
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRotues() {
	log.Info("Setting up routes")

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
			log.Fatal(err)
			panic(err)
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
		sendErrorResponse(w, "Enable to parse UINT from ID ", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))

	if err != nil {
		sendErrorResponse(w, "Error Fetching Comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		log.Error(err)
	}
}

// GetAllComments - retriews all comments from the database
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()

	if err != nil {
		sendErrorResponse(w, "Failed to Fetch Comments", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		sendErrorResponse(w, "JSON encoder error", err)
		return
	}

	if err := sendOkResponse(w, comments); err != nil {
		log.Error(err)
	}
}

// PostComment - add a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	// getting the comment from request body
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	// saving the comment on the database
	comment, err := h.Service.PostComment(comment)

	if err != nil {
		sendErrorResponse(w, "Failed to Post the Comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		log.Error(err)
	}
}

// UpdateComment - Update comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Error while Parsing id", err)
		return
	}

	// parsting of the request body
	var newComment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	comment, err := h.Service.UpdateComment(uint(commentId), newComment)

	if err != nil {
		sendErrorResponse(w, "Error updating the comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		log.Error(err)
	}
}

// DeleteComment - Delete a comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Enable to parse UINT from ID ", err)
		return
	}

	err = h.Service.DeleteComment(uint(commentId))

	if err != nil {
		sendErrorResponse(w, "Failed to delete the comment", err)
	}

	if err := sendOkResponse(w, Response{
		Message: "Successfully deleted the comment",
	}); err != nil {
		log.Error(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)

	if err := json.NewEncoder(w).Encode(Response{
		Message: message,
		Error:   err.Error(),
	}); err != nil {
		panic(err)
	}
}
