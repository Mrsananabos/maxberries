package main

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/internal/cronTask"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-stop
		cancel()
	}()

	config := configs.NewParsedConfig()
	cron, err := cronTask.CronConfig{
		CronExpression: "20 * * * *",
		TTL:            time.Minute * 30,
		Config:         config,
	}.CreateCronTask(ctx)

	if err != nil {
		log.Fatalf("error creating cron task %v", err)
	}

	cron.Start()
	log.Println("Background Worker Service started")
	defer cron.Stop()

	<-ctx.Done()
	log.Println("Background Worker Service stopped")
}
