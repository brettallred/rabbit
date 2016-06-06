package rabbit

import (
	"github.com/streadway/amqp"
	"os"
)

func connect() *amqp.Connection {
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	failOnError(err, "Failed to connect to RabbitMQ")
	return connection
}
