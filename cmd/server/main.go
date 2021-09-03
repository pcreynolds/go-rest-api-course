package main

import (
	"fmt"
	transportHTTP "github.com/pcreynolds/go-rest-api-course/internal/transport/http"
	"net/http"
)

// App - struct with default app config
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up our app")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}

}
