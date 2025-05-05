package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
	amqp	"github.com/rabbitmq/amqp091-go"

)

type HealthCheckResponse struct {
	Status   string   `json:"status"`
	Services []string `json:"services"`
}

func checkRabbitMQ() string {
	conn, err := amqp.Dial("amqp://zyra:password123@rabbitmq:5672/")
	if err != nil {
		return "RabbitMQ is not healthy"
	}
	defer conn.Close()
	return "RabbitMQ is healthy"
}

func HealthCheckHandler(c *gin.Context) {
	services := []string{
		checkRabbitMQ(),
	}

	response := HealthCheckResponse{
		Status:   "MSG Service is healthy!",
		Services: services,
	}

	c.JSON(http.StatusOK, response)
}
