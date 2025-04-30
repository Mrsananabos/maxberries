package kafka

import (
	"backgroundWorkerService/configs"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

const (
	ORDER_EVENTS_TOPIC = "order_events"
)

type Consumer struct {
	KafkaConsumer *kafka.Consumer
}

func CreateConsumer(cnf configs.Kafka) (Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s%s", cnf.Host, cnf.Port),
		"group.id":          "1",
	})
	if err != nil {
		return Consumer{}, fmt.Errorf("failed to create kafka consumer: %w", err)
	}
	return Consumer{
		KafkaConsumer: consumer,
	}, nil
}

func (c Consumer) SubscribeTopics(topics []string) error {
	err := c.KafkaConsumer.SubscribeTopics(topics, nil)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic %v: %w", topics, err)
	}

	log.Println("Consumer started")
	return nil
}
