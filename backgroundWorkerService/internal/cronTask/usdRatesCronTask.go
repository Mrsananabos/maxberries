package cronTask

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/http/client"
	"backgroundWorkerService/pkg/db"
	"context"
	"github.com/robfig/cron/v3"
	"time"
)

type CronConfig struct {
	CronExpression string
	TTL            time.Duration
	Config         configs.Config
}

func (c CronConfig) CreateCronTask(ctx context.Context) (*cron.Cron, error) {
	ct := cron.New()

	redis, err := db.Connect(c.Config.Redis)
	if err != nil {
		return ct, err
	}

	_, err = ct.AddFunc(c.CronExpression, func() {
		client.GetUSDRates(ctx, redis, c.Config.FixerAccessToken, c.TTL)
	})

	if err != nil {
		return ct, err
	}

	return ct, nil
}
