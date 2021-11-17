package message_queue

import (
	"fmt"

	"github.com/huypher/kit/rabbitmq"
)

const (
	priorityRange = 255
)

type Consumer interface {
	rabbitmq.Consumer
}

type consumer struct {
	consumer Consumer
}

func NewConsumer(cfg *Config, rabbitmqConsumer *RabbitmqConsumer) *consumer {
	c := new(consumer)

	c.consumer = rabbitmqConsumer.RabbitmqConn.CreateConsumer(
		rabbitmq.QueueName(cfg.QueueName),
		rabbitmq.PriorityQueue(priorityRange),
		rabbitmq.RegisterHandlerFunc(handler),
	)

	return c
}

func handler(data []byte) error {
	fmt.Println(string(data))
	return nil
}

func (c *consumer) Consume() {
	c.Consume()
}

func (c *consumer) Bind(exchangeName, routingKey string) error {
	return c.Bind(exchangeName, routingKey)
}
