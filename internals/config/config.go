package config

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	rootPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("Failed to get root directory:", err)
	}

	envPath := filepath.Join(rootPath, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		log.Printf("Error loading .env file from: %s, trying /app/.env instead...", envPath)

		err = godotenv.Load("/app/.env")
		if err != nil {
			log.Fatal("Error loading .env file from /app/.env")
		}
	}

	log.Println(".env file loaded successfully!")
	return nil
}
