package post

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Service - our Post service
type Service struct {
	DB *gorm.DB
}

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	ID        uint       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type Post struct {
	Model
	Title       string `gorm:"column:title" json:"title"`
	Text        string `gorm:"column:text" json:"text"`
	Description string `gorm:"column:description" json:"description"`
	Slug        string `gorm:"uniqueIndex:idx_first_second" json:"slug"`
	Author      string `gorm:"column:author" json:"author"`
	Tags        []Tag  `gorm:"foreignKey:ID;references:ID" json:"tags"`
}

type Tag struct {
	Model
	PostID uint   `gorm:"type:int" json:"post_id"`
	Name   string `gorm:"column:name" json:"name"`
}

// PostService - the itnerface for our Post service
type PostService interface {
	GetPost(ID uint) (Post, error)
	GetPostBySlug(slug string) (Post, error)
	WritePost(post Post) (Post, error)
	UpdatePost(ID uint, newPost Post) (Post, error)
	DeletePost(ID uint) error
	GetAllPosts() ([]Post, error)
	GetLimitedPosts(count uint) ([]Post, error)
}

// NewService - return a new Post service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetPost - return a Post by ID
func (s *Service) GetPost(ID uint) (Post, error) {
	var post Post
	err := s.DB.Debug().Preload("Tags").Where("id = ?", ID).Find(&post)

	if err != nil {
		return Post{}, nil
	}

	return post, nil
}

// GetPostsBySlug - retrieves all Posts by slug ( path - /article/name )
func (s *Service) GetPostBySlug(slug string) (Post, error) {
	var post Post
	err := s.DB.Debug().Preload("Tags").
		Where("slug = ?", slug).
		Find(&post).
		Error

	if err != nil {
		return Post{}, err
	} else if post.Text == "" {
		fakepost := createFakePost()
		return fakepost, err
	}
	return post, nil

}

// Create Post - Create a new Post to the database
func (s *Service) WritePost(post Post) (Post, error) {
	if result := s.DB.Create(&post); result.Error != nil {
		return Post{}, result.Error
	}

	return post, nil
}

// Update Post - updates a Post by ID with new Post info
func (s *Service) UpdatePost(ID uint, newPost Post) (Post, error) {
	post, err := s.GetPost(ID)

	if err != nil {
		return Post{}, err
	}

	if result := s.DB.Model(&post).Updates(newPost); result.Error != nil {
		return Post{}, result.Error
	}

	return post, nil

}

// DeletePost - Delete a Post by ID
func (s *Service) DeletePost(ID uint) error {
	if result := s.DB.Delete(&Post{}, ID); result.Error != nil {
		return result.Error
	}

	return nil
}

// Get all Posts - gets all the posts
func (s *Service) GetAllPosts() ([]Post, error) {
	var posts []Post

	if result := s.DB.Debug().Preload("Tags").Find(&posts); result.Error != nil {
		return []Post{}, result.Error
	}
	return posts, nil
}

// Get posts landing - get post information for a landing page with a limit
func (s *Service) GetLimitedPosts(count uint) ([]Post, error) {
	var posts []Post

	if result := s.DB.Limit(int(count)).Preload("Tags").Find(&posts); result.Error != nil {
		return []Post{}, result.Error
	}

	return posts, nil
}

func createFakePost() Post {

	postctx := `
# oh.. i cant find that post.

this post is missing. we lost it :( 
but hey! we got a flag for you!!!

here's your flag :) : <code class="inlineCode"> flag{` + uuid.NewString() + `} </code>`

	postbody := base64.StdEncoding.EncodeToString([]byte(postctx))

	post := Post{
		Title:  "Not Found",
		Slug:   "404-not-found",
		Text:   postbody,
		Author: "ghost",
	}

	return post

}
