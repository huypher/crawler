package infra

import (
	"github.com/huypher/crawler/internal/websocket"

	"github.com/huypher/crawler/internal/cronjob"

	"github.com/huypher/crawler/internal/cache"
	"github.com/huypher/crawler/internal/components"
	"github.com/huypher/crawler/internal/components/frontier"
	"github.com/huypher/crawler/internal/message_queue"
	"github.com/huypher/crawler/internal/voz"
)

func ProvideFrontier() components.Frontier {
	return frontier.NewFrontier()
}

func ProvideRabbitmqProducer(cfg *Config) *message_queue.RabbitmqProducer {
	config := &message_queue.Config{
		Addr:                    cfg.Rabbitmq.Addr,
		DelayExchangeName:       cfg.Rabbitmq.DelayExchangeName,
		DelayExchangeType:       cfg.Rabbitmq.DelayExchangeType,
		DelayExchangeRoutingKey: cfg.Rabbitmq.DelayExchangeRoutingKey,
		QueueName:               cfg.Rabbitmq.QueueName,
	}
	return message_queue.NewRabbitmqProducer(config)
}

func ProvideRabbitmqConsumer(cfg *Config) *message_queue.RabbitmqConsumer {
	config := &message_queue.Config{
		Addr:                    cfg.Rabbitmq.Addr,
		DelayExchangeName:       cfg.Rabbitmq.DelayExchangeName,
		DelayExchangeType:       cfg.Rabbitmq.DelayExchangeType,
		DelayExchangeRoutingKey: cfg.Rabbitmq.DelayExchangeRoutingKey,
		QueueName:               cfg.Rabbitmq.QueueName,
	}
	return message_queue.NewRabbitmqConsumer(config)
}

func ProvidePriorityProducer(cfg *Config, rabbitmqProducer *message_queue.RabbitmqProducer) message_queue.PriorityProducer {
	config := &message_queue.Config{
		Addr:                    cfg.Rabbitmq.Addr,
		DelayExchangeName:       cfg.Rabbitmq.DelayExchangeName,
		DelayExchangeType:       cfg.Rabbitmq.DelayExchangeType,
		DelayExchangeRoutingKey: cfg.Rabbitmq.DelayExchangeRoutingKey,
		QueueName:               cfg.Rabbitmq.QueueName,
	}
	return message_queue.NewPriorityProducer(config, rabbitmqProducer)
}

func ProvideConsumer(cfg *Config, rabbitmqConsumer *message_queue.RabbitmqConsumer) message_queue.Consumer {
	config := &message_queue.Config{
		Addr:                    cfg.Rabbitmq.Addr,
		DelayExchangeName:       cfg.Rabbitmq.DelayExchangeName,
		DelayExchangeType:       cfg.Rabbitmq.DelayExchangeType,
		DelayExchangeRoutingKey: cfg.Rabbitmq.DelayExchangeRoutingKey,
		QueueName:               cfg.Rabbitmq.QueueName,
	}
	return message_queue.NewConsumer(config, rabbitmqConsumer)
}

func ProvideVozExecutor(cache cache.Cache, websocket *websocket.Websocket) voz.Executor {
	return voz.NewVozExecutor(cache, websocket)
}

func ProvideCache(cfg *Config) (cache.Cache, func(), error) {
	return cache.NewCache(&cache.Config{
		Addr: cfg.Redis.Addr,
		Pass: cfg.Redis.Pass,
		DB:   cfg.Redis.DB,
	})
}

func ProvideCronJob() *cronjob.Cronjob {
	return cronjob.NewCronJob()
}

func ProvideWebsocketService() *websocket.Websocket {
	return websocket.NewWebsocketService()
}
