package queue

import (
	"fmt"
	"log"
	"reflect"
)

const (
	RabbitMQ AppQueueType = iota
)

type AppQueueType int

func New(qt AppQueueType, cfg any) (q *Queue, err error) {
	rt := reflect.TypeOf(cfg)

	switch qt {
	case RabbitMQ:
		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("configuration must be of type RabbitMQConfig")
		}
		fmt.Println("RabbitMQ not implemented")
	default:
		log.Fatal("Unsupported queue type")
	}
	return
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
