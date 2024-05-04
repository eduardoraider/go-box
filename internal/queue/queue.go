package queue

import (
	"fmt"
	"log"
)

const (
	RabbitMQ AppQueueType = iota
)

type AppQueueType int

func New(qt AppQueueType, cfg any) *Queue {
	q := new(Queue)
	switch qt {
	case RabbitMQ:
		fmt.Println("New RabbitMQ")
	default:
		log.Fatal("Unsupported queue type")
	}
	return q
}

type AppQueueConnection interface {
	Publish([]byte) error
	Consume() error
}

type Queue struct {
	cfg any
	qc  AppQueueConnection
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume() error {
	return q.qc.Consume()
}
