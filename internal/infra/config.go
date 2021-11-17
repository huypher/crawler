package infra

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
)

func ProvideConfig() (*Config, error) {
	cfg := &Config{}
	conf, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return nil, err
	}

	if err := conf.Scan(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

type Config struct {
	Rabbitmq Rabbitmq `json:"rabbitmq"`
	Redis    Redis    `json:"redis"`
}

type Rabbitmq struct {
	Addr string `json:"addr"`

	DelayExchangeName       string `json:"delay_exchange_name"`
	DelayExchangeType       string `json:"delay_exchange_type"`
	DelayExchangeRoutingKey string `json:"delay_exchange_routing_key"`

	QueueName string `json:"queue_name"`
}

type Redis struct {
	Addr string `json:"addr"`
	Pass string `json:"pass"`
	DB   int    `json:"db"`
}
