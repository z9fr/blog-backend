package main

import (
	"api/internal/comment"
	"api/internal/database"
	transportHttp "api/internal/transport/http"
	"log"
	"net/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run - handles the startup of our application
func (app *App) Run() error {

	log.Printf("Setting up the API on http://localhost:4000")

  db, err := database.NewDatabase()

	if err != nil {
		return err
	}

  commentService := comment.NewService(db)

	handler := transportHttp.NewHandler(commentService)
	handler.SetupRotues()

	if err := http.ListenAndServe(":4000", handler.Router); err != nil {
		return err
	}

	return nil
}

// Our main entrypoint for the application
func main() {

	app := App{}

	if err := app.Run(); err != nil {
		log.Panicf("error running the API")
		log.Panic(err)
	}

}
