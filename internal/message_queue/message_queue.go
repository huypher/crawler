package message_queue

import (
	"github.com/huypher/crawler/internal/infra"
	"github.com/huypher/kit/rabbitmq"
)

type RabbitmqProducer struct {
	RabbitmqConn rabbitmq.RabbitmqProducer
}

func NewRabbitmqProducer(cfg *infra.Config) *RabbitmqProducer {
	return &RabbitmqProducer{
		RabbitmqConn: rabbitmq.NewRabbitmq(rabbitmq.Address(cfg.Rabbitmq.Addr)),
	}
}

// ======================================================================================

type RabbitmqConsumer struct {
	RabbitmqConn rabbitmq.RabbitmqConsumer
}

func NewRabbitmqConsumer(cfg *infra.Config) *RabbitmqConsumer {
	return &RabbitmqConsumer{
		RabbitmqConn: rabbitmq.NewRabbitmq(rabbitmq.Address(cfg.Rabbitmq.Addr)),
	}
}
