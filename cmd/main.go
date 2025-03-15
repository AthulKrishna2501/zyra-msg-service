package main

import (
	"log"

	"github.com/AthulKrishna2501/zyra-msg-service/internals/broker"
)

func main() {
	log.Println("Starting Message Service...")
	broker.InitRabbitMQ()

	broker.ConsumeOTP()
}
