package cronTask

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/internal/usdRates/service"
	"backgroundWorkerService/pkg/db"
	"context"
	"github.com/robfig/cron/v3"
	"log"
)

type CronConfig struct {
	CronExpression string
	Config         configs.Config
}

func (c CronConfig) CreateCronTask(ctx context.Context) (*cron.Cron, error) {
	cronTask := cron.New()

	redis, err := db.Connect(c.Config.Redis)
	if err != nil {
		return nil, err
	}

	serv := service.NewService(c.Config, redis)

	_, err = cronTask.AddFunc(c.CronExpression, func() {
		_, err = serv.GetUSDRates(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	})

	if err != nil {
		return cronTask, err
	}

	return cronTask, nil
}
