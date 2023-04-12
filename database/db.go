package database

import (
	"book-simple-rest-api/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: don't forget to change the DB if needed
var (
	host     = "localhost"
	user     = "postgres"
	password = ""
	dbName   = "db-book"
	port     = "5432"
	sslmode  = "disable"
	db       *gorm.DB
	err      error
)

func StartDB() {
	log.Println("Connecting to database...")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	log.Println("Successfully connected to database")

	db.AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
