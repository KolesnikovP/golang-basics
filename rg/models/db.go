package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB

func SetupDataBase() *gorm.DB {
	var err error

	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Africa/Douala", host, username, password, databaseName, port)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("🚀🚀🚀---ASCENDE SUPERIUS---🚀🚀🚀")
	}

	return Database
}

func AutoMigrateModels() {
	Database.AutoMigrate(&User{})
}
