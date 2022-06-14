package database

import (
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(); err != nil {
		return err
	}
	return nil
}
