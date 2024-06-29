package database

import (
	"log"
	"os"

	"github.com/irksome0/pigeonTracker/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file!")
	}

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open database connection!")
	}

	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Pigeon{},
	)
}
