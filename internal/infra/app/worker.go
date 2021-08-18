package app

import "github.com/urfave/cli"

func (a *ApplicationContext) Consume() cli.Command {
	return cli.Command{
		Name:  "consume",
		Usage: "Start Service",
		Action: func(c *cli.Context) error {
			err := a.consumer.Bind(a.cfg.Rabbitmq.DelayExchangeName, a.cfg.Rabbitmq.DelayExchangeRoutingKey)
			if err != nil {
				panic(err.Error())
			}

			a.consumer.Consume()
			return nil
		},
	}
}
