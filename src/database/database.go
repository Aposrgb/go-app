package database

import (
	"app-aposrgb/src/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	connect()
	err := autoMigrate()
	if err != nil {
		fmt.Println("Couldn't migrate database")
	}
}

func autoMigrate() error {
	return DB.AutoMigrate(models.User{})
}

func connect() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
