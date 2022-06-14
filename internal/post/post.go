package post

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	pageSize = 10
)

type Service struct {
	DB *gorm.DB
}

type PostService interface {
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
