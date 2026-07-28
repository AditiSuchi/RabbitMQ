package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// // Publisher handles publishing messages to RabbitMQ exchanges
// type Publisher struct {
// 	// Add connection manager or channel references here
// }

// // NewPublisher creates a new Publisher instance
// func NewPublisher() *Publisher {
// 	return &Publisher{}
// }

func Publish(job ExportRequest) {

	body, _ := json.Marshal(job)

	channel.Publish(
		"",
		"report_queue",
		false,
		false,
		amqp.Publishing{
			Body: body,
		},
	)
}
