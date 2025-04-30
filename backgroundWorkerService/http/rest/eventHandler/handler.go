package eventHandler

import (
	"backgroundWorkerService/internal/orderPrice/service"
	"backgroundWorkerService/pkg/db/kafka"
	"backgroundWorkerService/pkg/db/kafka/message"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type Handler struct {
	orderPriceService service.Service
	consumer          kafka.Consumer
}

func NewHandler(consumer kafka.Consumer, service service.Service) Handler {
	return Handler{
		orderPriceService: service,
		consumer:          consumer,
	}
}

func (h Handler) Start(ctx context.Context) {
	defer h.consumer.KafkaConsumer.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("Kafka event handler shutting down")
			return
		default:
			msg, err := h.consumer.KafkaConsumer.ReadMessage(-1)
			if err == nil {
				log.Printf("Received message on %s: %s\n", *msg.TopicPartition.Topic, string(msg.Value))

				err = h.handleMsg(ctx, msg.Value)
				if err != nil {
					log.Println(err.Error())
				}
			} else {
				log.Printf("Consumer error: %v\n", err)
			}
		}
	}
}

func (h Handler) handleMsg(ctx context.Context, msg []byte) error {
	var orderMsg message.OrderMessage
	if err := json.Unmarshal(msg, &orderMsg); err != nil {
		log.Printf("failed unmarshal msg %s: %w", string(msg), err)
		return fmt.Errorf("failed unmarshal msg %s: %w", string(msg), err)
	}

	if err := orderMsg.Validate(); err != nil {
		return err
	}

	return h.orderPriceService.UpdateTotalOrderPrice(ctx, orderMsg)
}
