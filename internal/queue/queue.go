package queue

import (
	"fmt"
	"log"
	"reflect"
)

const (
	RabbitMQ AppQueueType = iota
	Mock
)

type AppQueueType int

func New(qt AppQueueType, cfg any) (q *Queue, err error) {
	q = new(Queue)
	rt := reflect.TypeOf(cfg)

	switch qt {
	case RabbitMQ:
		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("configuration must be of type RabbitMQConfig")
		}
		conn, err := newRabbitConn(cfg.(RabbitMQConfig))
		if err != nil {
			return nil, err
		}
		q.qc = conn
	case Mock:
		q.qc = &MockQueue{
			make([]*AppQueueDto, 0),
		}
	default:
		log.Fatal("Unsupported queue type")
	}
	return
}

type AppQueueConnection interface {
	Publish([]byte) error
	Consume(chan<- AppQueueDto) error
}

type Queue struct {
	qc AppQueueConnection
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume(cdto chan<- AppQueueDto) error {
	return q.qc.Consume(cdto)
}
