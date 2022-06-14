package user

import (
	"github.com/sirupsen/logrus"
)

func (s *Service) IsUsernameTaken(username string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&User{}).
		Select("count(*) > 0").Debug().
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		logrus.Warn(err)
	}

	return exists
}

func (s *Service) IsEmailTaken(email string) bool {
	var exists bool

	if err := s.DB.Debug().Model(&User{}).
		Select("count(*) > 0").Debug().
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		logrus.Warn(err)
	}

	return exists
}
