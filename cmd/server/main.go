package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/database"
	performancedb "github.com/z9fr/blog-backend/internal/performanceDb"
	"github.com/z9fr/blog-backend/internal/post"
	transportHttp "github.com/z9fr/blog-backend/internal/transport/http"
)

type App struct {
	Name    string
	Version string
	IsProd  bool
}

var startTime time.Time

func init() {
	startTime = time.Now()
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

	postservice := post.NewService(db)
	dbstatus := performancedb.NewService(db)

	// setup the routes and http handler
	handler := transportHttp.NewHandler(postservice, dbstatus, app.IsProd, startTime)
	handler.SetupRotues()

	if err := http.ListenAndServe(":4000", handler.Router); err != nil {
		return err
	}

	return nil
}

func main() {
	app := App{
		Name:    "api.z9fr.xyz",
		Version: "2.0.0",
		IsProd:  false,
	}

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}
