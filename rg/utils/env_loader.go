package utils

import (
	"github.com/joho/godotenv"
	"log"
)

// src/utils/dotenv_loader.go

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("---failed to load .env file---")
	}
}
