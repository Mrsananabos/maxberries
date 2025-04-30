package main

import (
	"backgroundWorkerService/configs"
	"backgroundWorkerService/http/rest"
	eventHandle "backgroundWorkerService/http/rest/eventHandler"
	"backgroundWorkerService/internal/servicesStorage"
	"backgroundWorkerService/internal/usdRates/cronTask"
	"backgroundWorkerService/pkg/db/kafka"
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

	config, err := configs.NewParsedConfig()
	if err != nil {
		panic(err)
	}

	services, err := servicesStorage.NewServicesStorage(config)
	if err != nil {
		panic(err)
	}

	cron, err := cronTask.CronConfig{
		CronExpression: "20 * * * *",
		Config:         config,
	}.CreateCronTask(ctx, services.InternalServices.USDRatesService)

	if err != nil {
		log.Fatalf("error creating cron task %v", err)
	}

	cron.Start()
	log.Println("Cron Background Worker started")
	defer cron.Stop()

	server, err := rest.NewServer(config, services.InternalServices)
	if err != nil {
		panic(err)
	}

	kafkaConsumer, err := kafka.CreateConsumer(config.Kafka)
	if err != nil {
		log.Printf("Kafka consumer failed: %s", err.Error())
	}

	if err = kafkaConsumer.SubscribeTopics([]string{kafka.ORDER_EVENTS_TOPIC}); err != nil {
		log.Printf("Kafka failed subscribe topics: %s", err.Error())
	}

	eventHandler := eventHandle.NewHandler(kafkaConsumer, services.InternalServices.OrderPriceService)

	go func() {
		eventHandler.Start(ctx)
		log.Println("Kafka event handler is starting")
	}()

	go func() {
		log.Println("Background Worker Service is starting")
		err = server.Run()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

	<-ctx.Done()
	log.Println("Background Worker Service stopped")
}
