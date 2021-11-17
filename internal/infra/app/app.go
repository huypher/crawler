package app

import (
	"fmt"
	"time"

	"github.com/huypher/crawler/internal/websocket"

	"github.com/huypher/crawler/internal/voz"

	"github.com/huypher/crawler/internal/cronjob"

	"github.com/huypher/crawler/internal/infra"

	"github.com/spf13/cobra"

	"github.com/google/wire"
)

type ApplicationContext struct {
	//cfg      *infra.Config
	//consumer message_queue.Consumer
	cronjob         *cronjob.Cronjob
	vozExec         voz.Executor
	websocketServer *websocket.Websocket
}

var ApplicationSet = wire.NewSet(
	infra.ProvideConfig,
	//infra.ProvideFrontier,
	infra.ProvideRabbitmqProducer,
	infra.ProvideRabbitmqConsumer,
	infra.ProvidePriorityProducer,
	infra.ProvideConsumer,
	infra.ProvideVozExecutor,
	infra.ProvideCache,
	infra.ProvideCronJob,
	infra.ProvideWebsocketService,
)

func (a *ApplicationContext) Run() {
	a.addCommands()
	Execute()
}

func (a *ApplicationContext) addCommands() {
	a.crawlCmd()
	a.wsCmd()
}

func (a *ApplicationContext) crawlCmd() {
	var crawlCmd = &cobra.Command{
		Use: "crawl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("websocket called")
			go a.websocketServer.Start()
			fmt.Println("crawl called")
			a.cronjob.SetFuncInterval(5*time.Second, a.vozExec.Do)
			a.cronjob.Start()
		},
	}

	rootCmd.AddCommand(crawlCmd)
}

func (a *ApplicationContext) wsCmd() {
	var crawlCmd = &cobra.Command{
		Use: "ws-serve",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("websocket called")
			a.websocketServer.Start()
		},
	}

	rootCmd.AddCommand(crawlCmd)
}
