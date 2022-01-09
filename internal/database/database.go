package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

//NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {
	// fmt.Println("Set up new database connection")
	log.Info("Setting up new database connection")

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Panic("Error loading environment", err)
	// }

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbTable, dbPassword, sslMode)
	fmt.Println("debuging maziar", connectionString)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	// Test connection to the database
	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
