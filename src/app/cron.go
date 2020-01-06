package app

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"

	// "github.com/owlsn/apis/src/services/task"
)

func StartSchedule() {
	// c := cron.New()
	// addCronFunc(c, "@every 30m", task.doJob)
	// c.Start()
}

func AddCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}
