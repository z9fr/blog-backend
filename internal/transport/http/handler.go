package http

import (
	"api/internal/comment"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
)

// Handler - store pointer to our comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to store responses from our api
type Response struct {
	Message string
	Error   string
}
// NewHandler - return a pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// LoggingMiddleware - a handy middleware function that logs out incoming requests
func LogginMiddleware(next http.Handler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r* http.Request){
		log.WithFields(
			log.Fields{
				"Method":      r.Method,
				"Path":        r.URL.Path,
        "Host": r.RemoteAddr,
			}).
			Info("handled request")
    next.ServeHTTP(w,r)
  })
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRotues() {
	log.Info("Setting up routes")

	// initicate new gorilla mox router
	h.Router = mux.NewRouter()
  h.Router.Use(LogginMiddleware)

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
