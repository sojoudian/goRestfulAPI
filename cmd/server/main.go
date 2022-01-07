package main

import (
	"fmt"
	"net/http"

	"github.com/sojoudian/goRestfulAPI/internal/database"
	transportHttp "github.com/sojoudian/goRestfulAPI/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

//
//Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	var err error
	//the following line should be like this db, err = database.NewDatabase()
	//but because we dont need db var in this line yet we will define db later
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}
func main() {
	fmt.Println("Go RestfulAPI")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting Go RestfulAPI")
		fmt.Println(err)
	}
}
