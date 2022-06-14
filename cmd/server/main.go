package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/database"
	transportHttp "github.com/z9fr/blog-backend/internal/transport/http"
	"net/http"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up Application")

	// setup the database

	db, err := database.NewDatabase()

	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	// setup the routes and http handler
	handler := transportHttp.NewHandler()
	handler.SetupRotues()

	if err := http.ListenAndServe(":4000", handler.Router); err != nil {
		return err
	}

	return nil
}

func main() {
	app := App{
		Name:    "api.z9fr.xyz",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}
