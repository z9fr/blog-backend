package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Handler struct {
	// service and router
	Router *chi.Mux
}

// NewHandler -  construcutre to create and return a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRotues() {
	h.Router = chi.NewRouter()

	// logs the start and end of each request, along with some useful data about what was requested,
	// what the response status was, and how long it took to return. When standard output is a TTY,
	// Logger will print in color, otherwise it will print in black and white. Logger prints a request ID if one is provided.
	h.Router.Use(middleware.Logger)

	// clean out double slash mistakes from a user's request path.
	// For example, if a user requests /users//1 or //users////1 will both be treated as: /users/1
	h.Router.Use(middleware.CleanPath)

	// automatically route undefined HEAD requests to GET handlers.
	h.Router.Use(middleware.GetHead)

	// recovers from panics, logs the panic (and a backtrace),
	// returns a HTTP 500 (Internal Server Error) status if possible. Recoverer prints a request ID if one is provided.
	h.Router.Use(middleware.Recoverer)

	h.Router.Route("/api/v2", func(r chi.Router) {

		r.Use(cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		r.Get("/", TestRoute)

		/* handle errors */

		h.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "route not found"})
		})

		h.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "method is not valid"})
		})
	})
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("flag{flag}"))
	return
}
