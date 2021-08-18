package infra

import (
	"github.com/huypher/crawler/internal/components"
	"github.com/huypher/crawler/internal/components/frontier"
	"github.com/huypher/crawler/internal/message_queue"
)

func ProvideFrontier() components.Frontier {
	return frontier.NewFrontier()
}

func ProvideRabbitmqProducer(cfg *Config) *message_queue.RabbitmqProducer {
	return message_queue.NewRabbitmqProducer(cfg)
}

func ProvideRabbitmqConsumer(cfg *Config) *message_queue.RabbitmqConsumer {
	return message_queue.NewRabbitmqConsumer(cfg)
}

func ProvidePriorityProducer(cfg *Config, rabbitmqProducer *message_queue.RabbitmqProducer) message_queue.PriorityProducer {
	return message_queue.NewPriorityProducer(cfg, rabbitmqProducer)
}

func ProvideConsumer(cfg *Config, rabbitmqConsumer *message_queue.RabbitmqConsumer) message_queue.Consumer {
	return message_queue.NewConsumer(cfg, rabbitmqConsumer)
}
