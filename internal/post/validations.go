package post

import (
	"github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/types"
)

func (s *Service) IsTitleTaken(title string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&types.Post{}).
		Select("count(*) > 0").Debug().
		Where("title = ?", title).
		Find(&exists).
		Error; err != nil {
		logrus.Warn(err)
	}

	return exists
}

func (s *Service) IsSlugTaken(slug string) bool {
	var exists bool

	if err := s.DB.Debug().Model(&types.Post{}).
		Select("count(*) > 0").Debug().
		Where("slug = ?", slug).
		Find(&exists).
		Error; err != nil {
		logrus.Warn(err)
	}

	// just for testing
	if slug == "test" {
		exists = true
	}

	return exists
}
