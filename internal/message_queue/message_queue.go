package message_queue

import (
	"github.com/huypher/kit/rabbitmq"
)

type Config struct {
	Addr string

	DelayExchangeName       string
	DelayExchangeType       string
	DelayExchangeRoutingKey string

	QueueName string
}

type RabbitmqProducer struct {
	RabbitmqConn rabbitmq.RabbitmqProducer
}

func NewRabbitmqProducer(cfg *Config) *RabbitmqProducer {
	return &RabbitmqProducer{
		RabbitmqConn: rabbitmq.NewRabbitmq(rabbitmq.Address(cfg.Addr)),
	}
}

// ======================================================================================

type RabbitmqConsumer struct {
	RabbitmqConn rabbitmq.RabbitmqConsumer
}

func NewRabbitmqConsumer(cfg *Config) *RabbitmqConsumer {
	return &RabbitmqConsumer{
		RabbitmqConn: rabbitmq.NewRabbitmq(rabbitmq.Address(cfg.Addr)),
	}
}
