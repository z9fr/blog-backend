package main

import "fmt"

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run - handles the startup of our application
func (app *App) Run() error {
	fmt.Print("Server is Starting")
	return nil
}

// Our main entrypoint for the application
func main() {

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Print("Error running the REST api")
		fmt.Print(err)
	}

}
