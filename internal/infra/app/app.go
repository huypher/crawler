package app

import (
	"github.com/huypher/crawler/internal/infra"
	"github.com/huypher/crawler/internal/message_queue"

	"github.com/google/wire"
	"github.com/urfave/cli"
)

type ApplicationContext struct {
	cfg      *infra.Config
	consumer message_queue.Consumer
}

var ApplicationSet = wire.NewSet(
	infra.ProvideConfig,
	//infra.ProvideFrontier,
	infra.ProvideRabbitmqProducer,
	infra.ProvideRabbitmqConsumer,
	infra.ProvidePriorityProducer,
	infra.ProvideConsumer,
)

func (a *ApplicationContext) Commands() *cli.App {
	app := cli.NewApp()
	app.Commands = []cli.Command{}

	return app
}
