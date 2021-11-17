package cronjob

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type Cronjob struct {
	cron *cron.Cron
}

func NewCronJob() *Cronjob {
	return &Cronjob{
		cron: cron.New(cron.WithLogger(cron.DefaultLogger)),
	}
}

func (c *Cronjob) SetFuncInterval(d time.Duration, f func()) {
	expression := fmt.Sprintf("@every %fs", d.Seconds())
	c.cron.AddFunc(expression, f)
}

func (c *Cronjob) Start() {
	c.cron.Start()
	for {
		continue
	}
}
