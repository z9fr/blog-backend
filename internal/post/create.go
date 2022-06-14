package post

import (
	"errors"

	"github.com/z9fr/blog-backend/internal/types"
	"github.com/z9fr/blog-backend/internal/utils"
)

func (s *Service) CreatePost(post types.Post) (types.Post, error) {
	if s.IsTitleTaken(post.Title) {
		return types.Post{}, errors.New("post with the same title already exist")
	}

	generatedSlug := func(postTitle string) string {

		postTitle = utils.GenerateSlug(postTitle, false)

		if !s.IsSlugTaken(postTitle) {
			return postTitle
		}

		return utils.GenerateSlug(postTitle, true)
	}

	post.Slug = generatedSlug(post.Title)

	return post, nil
}
