package cronTask

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/internal/usdRates/service"
	"context"
	"github.com/robfig/cron/v3"
	"log"
)

type CronConfig struct {
	CronExpression string
	Config         configs.Config
}

func (c CronConfig) CreateCronTask(ctx context.Context, usdRatesService service.Service) (*cron.Cron, error) {
	cronTask := cron.New()

	_, err := cronTask.AddFunc(c.CronExpression, func() {
		_, err := usdRatesService.GetUSDRates(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	})

	if err != nil {
		return cronTask, err
	}

	return cronTask, nil
}
