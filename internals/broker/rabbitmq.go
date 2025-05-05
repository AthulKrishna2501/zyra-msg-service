package broker

import (
	"encoding/json"
	"log"

	"github.com/AthulKrishna2501/zyra-msg-service/internals/email"
	"github.com/rabbitmq/amqp091-go"
)

var RabbitMQConn *amqp091.Connection

func InitRabbitMQ() {
	var err error
	RabbitMQConn, err = amqp091.Dial("amqp://zyra:password123@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	log.Println("Connected to RabbitMQ!")
}

func ConsumeOTP() {
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("otp_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	log.Println("Waiting for OTP messages...")

	for msg := range msgs {
		var data map[string]string
		json.Unmarshal(msg.Body, &data)

		UserEmail := data["email"]
		UserOtp := data["otp"]

		email.SendOTPEmail(UserEmail, UserOtp)

		log.Printf("Sent OTP %s to %s", UserOtp, UserEmail)
	}
}
