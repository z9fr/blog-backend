package http

import (
	"fmt"
	"net/http"

	"encoding/xml"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
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

func (h *Handler) RSSFeed(w http.ResponseWriter, r *http.Request) {
	posts := h.PostService.GetAllPosts()
	data, err := xml.Marshal(&posts)

	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(http.StatusOK)

	if err != nil {
		logrus.Warn(err)
	}

	w.Write([]byte("<rss version='2.0'>"))
	w.Write(data)
	w.Write([]byte("</rss>"))
}
