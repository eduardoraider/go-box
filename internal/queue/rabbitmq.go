package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Time
}

type RabbitConnection struct {
	cfg  RabbitMQConfig
	conn *amqp.Connection
}

func (rc *RabbitConnection) Publish(msg []byte) error {

}

func (rc *RabbitConnection) Consume() error {
	
}
