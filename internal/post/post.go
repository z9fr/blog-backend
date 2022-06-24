package post

import (
	"github.com/z9fr/blog-backend/internal/types"
	"gorm.io/gorm"
)

const (
	pageSize = 10
)

type Service struct {
	DB *gorm.DB
}

type PostService interface {
	// create
	WritePost(post types.Post) (types.Post, error)

	// fetch
	GetAllPosts() []*types.Post
	GetPostsBySlug(slug string) types.Post
	GetAllUnPublishedPosts() []*types.Post

	// validations
	IsTitleTaken(title string) bool
	IsSlugTaken(slug string) bool
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
