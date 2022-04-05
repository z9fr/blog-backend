package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/z9fr/blog-backend/internal/post"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Getpost - Retriew post by ID
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	// get id from user parameter
	vars := mux.Vars(r)
	id := vars["id"] // this is a string but id is uint so we need to convert

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Enable to parse UINT from ID ", err)
		return
	}

	post, err := h.ServicePost.GetPost(uint(i))

	if err != nil {
		sendErrorResponse(w, "Error Fetching post", err)
		return
	}

	if err := sendOkResponse(w, post); err != nil {
		log.Error(err)
	}
}

// Getpost - Retriew Based on Slug
func (h *Handler) GetPostBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	post, err := h.ServicePost.GetPostBySlug(slug)
	if err != nil {
		sendErrorResponse(w, "Error Fetching post", err)
		return
	}

	if err := sendOkResponse(w, post); err != nil {
		log.Error(err)
	}
}

// GetAllposts - retriews all posts from the database
func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.ServicePost.GetAllPosts()

	if err != nil {
		sendErrorResponse(w, "Failed to Fetch posts", err)
		return
	}

	if err := sendOkResponse(w, posts); err != nil {
		log.Error(err)
	}
}

// create post - add a new post
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		log.Debug("Options requests")
		sendOkResponse(w, "options")
		return
	}

	var post post.Post
	// getting the post from request body
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	// saving the post on the database
	post, err := h.ServicePost.WritePost(post)

	if err != nil {
		sendErrorResponse(w, "Failed to Post the post", err)
		return
	}

	if err := sendOkResponse(w, post); err != nil {
		log.Error(err)
	}
}

// Updatepost - Update post
func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	postId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Error while Parsing id", err)
		return
	}

	// parsting of the request body
	var newpost post.Post
	if err := json.NewDecoder(r.Body).Decode(&newpost); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	post, err := h.ServicePost.UpdatePost(uint(postId), newpost)

	if err != nil {
		sendErrorResponse(w, "Error updating the post", err)
		return
	}

	if err := sendOkResponse(w, post); err != nil {
		log.Error(err)
	}
}

// Deletepost - Delete a post by id
func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	postId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Enable to parse UINT from ID ", err)
		return
	}

	err = h.ServicePost.DeletePost(uint(postId))

	if err != nil {
		sendErrorResponse(w, "Failed to delete the post", err)
	}

	if err := sendOkResponse(w, Response{
		Message: "Successfully deleted the post",
	}); err != nil {
		log.Error(err)
	}
}
