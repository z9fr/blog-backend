package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
	ID       string
}

// userSerive - the itnerface for our User Service
type UserService interface {
	GetUser(username string) (User, error)
	CreateUser(user User) (User, error)
	DeleteUser(ID string) error
}

// NewService - return a new Post service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetUser - return a user using username
func (s *Service) GetUser(username string) (User, error) {
	var user User

	if result := s.DB.First(&user, username); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (s *Service) CreateUser(user User) (User, error) {
	return User{}, fmt.Errorf("not implemented")
}

func (s *Service) DeleteUser(ID string) error {
	return fmt.Errorf("not implemented")
}
