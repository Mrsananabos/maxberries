package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"orderService/configs"
	"orderService/pkg/kafka/message"
)

const (
	ORDER_EVENTS_TOPIC = "order_events"
)

type Producer struct {
	KafkaProducer *kafka.Producer
}

func CreateProducer(cnf configs.Kafka) (Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s%s", cnf.Host, cnf.Port)})
	if err != nil {
		return Producer{}, fmt.Errorf("failed to create kafka producer: %w", err)
	}
	return Producer{
		KafkaProducer: producer,
	}, nil
}

func (p Producer) SentMsg(topic string, msg message.OrderCreatedMsg) (offset int64, err error) {
	jsonData, err := msg.ToJson()
	if err != nil {
		return 0, err
	}

	deliveryChan := make(chan kafka.Event)
	err = p.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          jsonData,
	}, deliveryChan)

	if err != nil {
		return 0, fmt.Errorf("failed to send message: %w", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		err = fmt.Errorf("error delivering message %s: %w", msg, m.TopicPartition.Error)
		return 0, m.TopicPartition.Error
	}

	log.Printf("message delivering in topic %v [%d] with offset %v\n",
		*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)

	return int64(m.TopicPartition.Offset), nil
}
