package config

import (
	"context"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ChatCollection *mongo.Collection

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

func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(context.Background(), clientOptions)
	ChatCollection = client.Database("chatdb").Collection("messages")
}
