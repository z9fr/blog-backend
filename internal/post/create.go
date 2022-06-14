package post

import (
	"errors"

	"github.com/z9fr/blog-backend/internal/types"
)

func (s *Service) CreatePost(post types.Post) (types.Post, error) {
	if s.IsTitleTaken(post.Title) {
		return types.Post{}, errors.New("post with the same title already exist")
	}

	return types.Post{}, nil
}
