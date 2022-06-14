package utils

import (
	"testing"

	"github.com/z9fr/blog-backend/internal/database"
)

func TestDbConnection(t *testing.T) {
	db, err := database.NewDatabase()

	if err != nil {
		t.Errorf("Database connection failed %s", err.Error())
	}

	err = database.MigrateDB(db)
	if err != nil {
		t.Errorf("Database Migration Failed %s", err.Error())
	}
}
