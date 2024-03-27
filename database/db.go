package database

import (
	"fmt"
	"log"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "12345678"
	dbPort   = "5432"
	dbName   = "finalproject"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	fmt.Println("Connection success")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.Socialmedia{})
}

func GetDB() *gorm.DB {
	return db
}
