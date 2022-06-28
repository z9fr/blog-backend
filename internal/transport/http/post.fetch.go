package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) FetchallPosts(w http.ResponseWriter, r *http.Request) {
	posts := h.PostService.GetAllPosts()
	h.sendOkResponse(w, posts)
	return
}

func (h *Handler) FetcheventbySlug(w http.ResponseWriter, r *http.Request) {
	post_slug := chi.URLParam(r, "slug")

	post_exist := h.PostService.IsSlugTaken(post_slug)

	if !post_exist {
		h.sendErrorResponse(w, "404 post not found", fmt.Errorf("post not found"), http.StatusNotFound)
		return
	}

	post := h.PostService.GetPostsBySlug(post_slug)

	h.sendOkResponse(w, post)
	return
}

func (h *Handler) FetchUnpublishedAllEvents(w http.ResponseWriter, r *http.Request) {
	// not implemented
	return
}

func (h *Handler) UnpublishedPosts(w http.ResponseWriter, r *http.Request) {
	posts := h.PostService.GetUnpublishedPosts()
	h.sendOkResponse(w, posts)
	return
}
