package main

import (
	"log"

	"github.com/AthulKrishna2501/zyra-msg-service/internals/broker"
	"github.com/AthulKrishna2501/zyra-msg-service/internals/config"
	"github.com/AthulKrishna2501/zyra-msg-service/internals/healthcheck"
	"github.com/AthulKrishna2501/zyra-msg-service/internals/websocket"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Starting Message Service...")
	broker.InitRabbitMQ()

	go broker.ConsumeOTP()
	config.ConnectMongoDB()

	router.GET("/ws", websocket.WebSocketHandler)
	router.GET("/health", healthcheck.HealthCheckHandler)

	router.Run(":8082")

}
