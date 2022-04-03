package database

import (
	"github.com/z9fr/blog-backend/internal/post"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	// MigrateDB - migrates our database and create our comment tables
	if err := db.AutoMigrate(&post.Post{}); err != nil {
		return err
	}
	return nil
}
