package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
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
	log.SetFormatter(&log.JSONFormatter{})
	// fmt.Println("Setting up our APP")
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up application")

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
		log.Error("Failed to set up server")
		return err
	}

	return nil
}
func main() {
	// fmt.Println("Go RestfulAPI")
	app := App{
		Name:    "Commenting Service",
		Version: "0.0.1",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting Go RestfulAPI")
		log.Fatal(err)
	}
}
