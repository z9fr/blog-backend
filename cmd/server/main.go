package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/z9fr/blog-backend/internal/database"
	performancedb "github.com/z9fr/blog-backend/internal/performanceDb"
	"github.com/z9fr/blog-backend/internal/post"
	transportHttp "github.com/z9fr/blog-backend/internal/transport/http"
	"github.com/z9fr/blog-backend/internal/user"
	"github.com/z9fr/blog-backend/internal/utils"
)

type App struct {
	Name    string
	Version string
	IsProd  bool
}

var startTime time.Time
var ApplicationSecret string

func init() {
	startTime = time.Now()
	secret := "secret"

	if os.Getenv("ENV") == "prod" {
		secret, _ = utils.SecretGenerator(100)
	}

	ApplicationSecret = secret
}

func (app *App) Run() error {
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
			"AppSecret":  ApplicationSecret,
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
	userservice := user.NewService(db)
	dbstatus := performancedb.NewService(db)

	// setup the routes and http handler
	handler := transportHttp.NewHandler(postservice, userservice, dbstatus, app.IsProd, startTime, ApplicationSecret)
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
