package main

import (
	"fmt"
	"net/http"

	"github.com/sojoudian/goRestfulAPI/internal/comment"
	"github.com/sojoudian/goRestfulAPI/internal/database"
	transportHttp "github.com/sojoudian/goRestfulAPI/internal/transport/http"
)

// App - contain application information
type App struct {
	Name    string
	Version string
}

//
//Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	var err error
	//the following line should be like this db, err = database.NewDatabase()
	//but because we dont need db var in this line yet we will define db later

	//now we have our database implemented we can change _ to db
	//_, err = database.NewDatabase()
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	commentService := comment.NewService(db)

	handler := transportHttp.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}
func main() {
	fmt.Println("Go RestfulAPI")
	app := App{
		Name: "Commenting Service",
		Version: "0.0.1"
	}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting Go RestfulAPI")
		fmt.Println(err)
	}
}
