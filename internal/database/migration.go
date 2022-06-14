package database

import (
	"github.com/z9fr/blog-backend/internal/types"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(types.Post{}); err != nil {
		return err
	}
	return nil
}
