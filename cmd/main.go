package main

import (
	"log"

	"github.com/AthulKrishna2501/zyra-msg-service/internals/broker"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Starting Message Service...")
	broker.InitRabbitMQ()

	broker.ConsumeOTP()
}
