package post

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/types"
	"github.com/z9fr/blog-backend/internal/utils"
)

func (s *Service) WritePost(post types.Post) (types.Post, error) {
	if s.IsTitleTaken(post.Title) {
		return types.Post{}, errors.New("post with the same title already exist")
	}

	// generate a slug based on the post title
	generatedSlug := func(postTitle string) string {

		postTitle = utils.GenerateSlug(postTitle, false)

		if !s.IsSlugTaken(postTitle) {
			return postTitle
		}

		return utils.GenerateSlug(postTitle, true)
	}

	post.Slug = generatedSlug(post.Title)

	if result := s.DB.Debug().Save(&post); result.Error != nil {
		logrus.Warn(result.Error)
		return types.Post{}, result.Error
	}

	return post, nil
}

func (s *Service) PublishPost(postSlug string) bool {
	isExist := s.IsSlugTaken(postSlug)

	if isExist {
		postdetails := s.GetPostsBySlug(postSlug)
		postdetails.IsPublic = true
		s.DB.Debug().Save(postdetails)
		return true
	}
	return false
}
