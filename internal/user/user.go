package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/z9fr/blog-backend/internal/utils"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	UserName    string `gorm:"column:username;uniqueIndex:idx_username" json:"username"`
	Description string `gorm:"column:description" json:"description"`
	Email       string `gorm:"column:email;uniqueIndex:idx_email" json:"email"`
	Password    string `gorm:"column:password" json:"password"`
}

// userSerive - the itnerface for our User Service
type UserService interface {
	GetUserbyEmail(email string) (User, error)
	GetUserbyUsername(username string) (User, error)
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
func (s *Service) GetUserbyUsername(username string) (User, error) {
	var user User

	if result := s.DB.First(&user, "username = ?", username); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

// GetUser - return a user based on email
func (s *Service) GetUserbyEmail(email string) (User, error) {
	var user User

	if result := s.DB.First(&user, "email = ?", email); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

// Create a new User
func (s *Service) CreateUser(user User) (User, error) {

	hashedPassword, err := utils.HashPassword(user.Password)
	user.ID = uuid.NewString()

	if err != nil {
		return User{}, fmt.Errorf("Unable to hash user password %w", err)
	}

	user.Password = hashedPassword

	if result := s.DB.Save(&user); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

// delete user - deletes a user using given id
func (s *Service) DeleteUser(ID string) error {
	return fmt.Errorf("not implemented")
}
