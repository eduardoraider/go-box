package queue

import (
	"context"
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
	c, err := rc.conn.Channel()
	if err != nil {
		return err
	}
	mp := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         msg,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	return c.PublishWithContext(ctx,
		"",
		rc.cfg.TopicName,
		false,
		false,
		mp)
}

func (rc *RabbitConnection) Consume() error {
	ch, err := rc.conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(rc.cfg.TopicName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for d := range msgs {
		
	}
}
