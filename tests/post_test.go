package utils

import (
	"testing"

	"github.com/z9fr/blog-backend/internal/database"
	"github.com/z9fr/blog-backend/internal/post"
	"github.com/z9fr/blog-backend/internal/types"
	"gorm.io/gorm"
)

func DBConnection(t *testing.T) (*gorm.DB, error) {
	db, err := database.NewDatabase()

	if err != nil {
		t.Errorf("Database connection failed %s", err.Error())
		return nil, err
	}

	err = database.MigrateDB(db)
	if err != nil {
		t.Errorf("Database Migration Failed %s", err.Error())
		return nil, err
	}

	return db, nil

}

func PostService(t *testing.T) (post.Service, error) {
	db, err := DBConnection(t)

	service := post.NewService(db)

	if err != nil {
		return *service, err
	}

	return *service, nil
}

func TestCreatePost(t *testing.T) {
	postservice, err := PostService(t)

	if err != nil {
		t.Errorf("database level error %s", err.Error())
	}

	post, err := postservice.CreatePost(types.Post{
		Title:      "test",
		Descrption: "Descrption",
		Body:       "post body",
		IsPublic:   true,
	})

	if err != nil {
		t.Errorf("failed to create the post %s", err.Error())
	}

	t.Log(post.Slug)
}
