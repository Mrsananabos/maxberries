package main

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/http/rest"
	"backgroundWorkerService/internal/usdRates/cronTask"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal)
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
		Config:         config,
	}.CreateCronTask(ctx)

	if err != nil {
		log.Fatalf("error creating cron task %v", err)
	}

	cron.Start()
	log.Println("Cron Background Worker started")
	defer cron.Stop()

	server, err := rest.NewServer()
	if err != nil {
		panic(err)
	}
	log.Println("Background Worker Service started")

	err = server.Run()
	if err != nil {
		panic(err)
	}

	<-ctx.Done()
	log.Println("Background Worker Service stopped")
}
