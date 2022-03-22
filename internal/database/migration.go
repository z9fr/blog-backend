package database

import (
	"api/internal/comment"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	// MigrateDB - migrates our database and create our comment tables
	if err := db.AutoMigrate(&comment.Comment{}); err != nil {
		return err
	}
	return nil
}
