package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	performancedb "github.com/z9fr/blog-backend/internal/performanceDb"
	"github.com/z9fr/blog-backend/internal/post"
	"github.com/z9fr/blog-backend/internal/user"

	"github.com/go-chi/httprate"
)

type Handler struct {
	// service and router
	Router              *chi.Mux
	PostService         *post.Service
	UserService         *user.Service
	PerformanceDatabase *performancedb.Service
	IsProd              bool
	StartTime           time.Time
	ApplicationSecret   string
}

// NewHandler -  construcutre to create and return a pointer to a handler
func NewHandler(postservice *post.Service, userservice *user.Service, dbstatus *performancedb.Service, isprod bool, starttime time.Time, secret string) *Handler {
	return &Handler{
		PostService:         postservice,
		UserService:         userservice,
		PerformanceDatabase: dbstatus,
		IsProd:              isprod,
		StartTime:           starttime,
		ApplicationSecret:   secret,
	}
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

	// RedirectSlashes is a middleware that will match request paths with a trailing slash
	// and redirect to the same path, less the trailing slash.
	h.Router.Use(middleware.RedirectSlashes)

	// automatically route undefined HEAD requests to GET handlers.
	h.Router.Use(middleware.GetHead)

	// Throttle is a middleware that limits number of currently processed requests at a time
	// across all users. Note: Throttle is not a rate-limiter per user, instead it just puts a
	// ceiling on the number of currentl in-flight requests being processed from the point
	// from where the Throttle middleware is mounted.
	h.Router.Use(middleware.Throttle(15))

	// ThrottleBacklog is a middleware that limits number of currently processed requests
	// at a time and provides a backlog for holding a finite number of pending requests
	h.Router.Use(middleware.ThrottleBacklog(10, 50, time.Second*10))

	// timeout middleware
	h.Router.Use(middleware.Timeout(time.Second * 60))

	// recovers from panics, logs the panic (and a backtrace),
	// returns a HTTP 500 (Internal Server Error) status if possible. Recoverer prints a request ID if one is provided.
	h.Router.Use(middleware.Recoverer)

	// RealIP is a middleware that sets a http.Request's RemoteAddr to the results of parsing either
	// the X-Real-IP header or the X-Forwarded-For header (in that order).
	h.Router.Use(middleware.RealIP)

	// Enable httprate request limiter of 100 requests per minute.
	//
	// rate-limiting is bound to the request IP address via the LimitByIP middleware handler.
	//
	// To have a single rate-limiter for all requests, use httprate.LimitAll(..).
	h.Router.Use(httprate.LimitByIP(100, 1*time.Minute))

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

		r.Get("/posts", h.FetchallPosts)
		r.Get("/post/{slug}", h.FetcheventbySlug)
		r.Get("/user/{username}", h.FetchuserbyUsername)
		r.Get("/health", h.GetApplicationHealth)

		r.Post("/login", h.AuthUser)

		// do not allow regisration on prod
		if !h.IsProd {
			r.Post("/user/create", h.CreateUser)
		}

		r.Route("/post/create", func(r chi.Router) {
			r.Use(h.JWTMiddlewhare)
			r.Post("/", h.WritePostHandler)
		})

		r.Route("/post/update", func(r chi.Router) {
			r.Use(h.JWTMiddlewhare)
			r.Put("/publish", h.PublishPost)
		})

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
