package database

import (
	"github.com/z9fr/blog-backend/internal/types"
	"github.com/z9fr/blog-backend/internal/user"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(types.Post{}, user.User{}); err != nil {
		return err
	}
	return nil
}
