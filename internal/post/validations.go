package post

import (
	"github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/types"
)

func (s *Service) IsTitleTaken(title string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&types.Post{}).
		Select("count(*) > 0").
		Where("title = ?", title).
		Find(&exists).
		Error; err != nil {
		logrus.Warn(err)
	}

	return exists
}
