package http

import (
	"encoding/json"
	"net/http"

	"github.com/z9fr/blog-backend/internal/helper"

	"github.com/z9fr/blog-backend/internal/types"
	"github.com/z9fr/blog-backend/internal/user"
)

func (h *Handler) WritePostHandler(w http.ResponseWriter, r *http.Request) {
	userdetails := r.Context().Value("user").(user.User)
	var postdetails types.PostCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&postdetails); err != nil {
		h.sendErrorResponse(w, "Failed to deocde JSON body", err, http.StatusBadRequest)
		return
	}

	// Createpostdetails
	postcontent, err := helper.Createpostdetails(postdetails, userdetails)

	if err != nil {
		h.sendErrorResponse(w, "Unable to create the post", err, http.StatusBadRequest)
		return
	}

	created_post, err := h.PostService.WritePost(postcontent)

	if err != nil {
		h.sendErrorResponse(w, "Unable to create the post", err, http.StatusInternalServerError)
		return
	}

	h.sendOkResponse(w, created_post)
	return
}

func (h *Handler) PublishPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	details := struct {
		Slug string `json:"slug"`
	}{}
	if err := decoder.Decode(&details); err != nil {
		h.sendErrorResponse(w, "Failed to deocde JSON body", err, http.StatusBadRequest)
		return
	}

	isPosted := h.PostService.PublishPost(details.Slug)

	h.sendOkResponse(w, struct {
		Success bool `success`
	}{
		Success: isPosted,
	})

	return
}
