package app

import (
	"crawler/internal/infra"

	"github.com/google/wire"
	"github.com/urfave/cli"
)

type ApplicationContext struct {
}

var ApplicationSet = wire.NewSet(
	infra.ProvideFrontier,
)

func (a *ApplicationContext) Commands() *cli.App {
	app := cli.NewApp()
	app.Commands = []cli.Command{}

	return app
}
