package http

import (
	"encoding/json"
	"net/http"

	"github.com/z9fr/blog-backend/internal/post"

	user "github.com/z9fr/blog-backend/internal/user"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler - store pointer to our comment service
type Handler struct {
	Router      *mux.Router
	ServicePost *post.Service
	ServiceUser *user.Service
}

// Response - an object to store responses from our api
type Response struct {
	Message string
	Error   string
}

// NewHandler - return a pointer to a handler
func NewHandler(postservice *post.Service, userservice *user.Service) *Handler {
	return &Handler{
		ServicePost: postservice,
		ServiceUser: userservice,
	}
}

// LoggingMiddleware - a handy middleware function that logs out incoming requests
func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
				"Host":   r.RemoteAddr,
			}).
			Info("handled request")
		next.ServeHTTP(w, r)
	})
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRotues() {
	log.Info("Setting up routes")

	// initicate new gorilla mox router
	h.Router = mux.NewRouter()
	h.Router.Use(LogginMiddleware)

	// posts
	h.Router.HandleFunc("/api/v1/posts", h.GetAllPosts).Methods("GET")
	h.Router.HandleFunc("/api/v1/post/{id}", h.GetPost).Methods("GET")
	h.Router.HandleFunc("/api/v1/post/create", h.CreatePost).Methods("POST")
	h.Router.HandleFunc("/api/v1/post/delete/{id}", h.DeletePost).Methods("DELETE")
	h.Router.HandleFunc("/api/v1/post/update/{id}", h.UpdatePost).Methods("PUT")

	// users
	h.Router.HandleFunc("/api/v1/user/{username}", h.GetUser).Methods("GET")

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

// handle ok responses
func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

// handle error responses
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
