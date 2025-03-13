package main

import (
	"log"

	"github.com/AthulKrishna2501/zyra-msg-service/internals/broker"
	"github.com/AthulKrishna2501/zyra-msg-service/internals/cache"
)

func main() {
	log.Println("Starting Message Service...")
	cache.InitRedis()
	broker.InitRabbitMQ()

	broker.ConsumeOTP()
}
