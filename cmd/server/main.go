package main

import (
	"api/internal/post"
	"api/internal/database"
	transportHttp "api/internal/transport/http"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// App -  Contain the application information
type App struct {
	Name    string
	Version string
}

// Run - handles the startup of our application
func (app *App) Run() error {
	//log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up Application")

	db, err := database.NewDatabase()

	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Error(err)
		log.Fatal(err)
		panic(err)
	}

	postService := post.NewService(db)

	handler := transportHttp.NewHandler(postService)
	handler.SetupRotues()

	if err := http.ListenAndServe(":4000", handler.Router); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Our main entrypoint for the application
func main() {

	app := App{
		Name:    "api-prod.dasith.works",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error(err)
		log.Fatal(err)
	}

}
