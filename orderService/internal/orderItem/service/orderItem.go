package service

import (
	"fmt"
	"log"
	"orderService/http/rest/client"
	"orderService/internal/orderItem/model"
	"orderService/internal/orderItem/repository"
	"orderService/internal/orderStatus/service"
	"orderService/pkg/kafka"
	"orderService/pkg/kafka/message"
)

type Service struct {
	repo               repository.Repository
	orderStatusService service.Service
	httpClient         client.HttpClient
	kafkaProducer      kafka.Producer
}

func NewService(r repository.Repository, orderStatusService service.Service, httpClient client.HttpClient, kafkaProducer kafka.Producer) Service {
	return Service{
		repo:               r,
		orderStatusService: orderStatusService,
		httpClient:         httpClient,
		kafkaProducer:      kafkaProducer,
	}
}

func (s Service) UpdateItems(orderId int64, newItems model.EditOrderItemsRequest) error {
	for _, item := range newItems.Items {
		price, err := s.httpClient.GetProductPrice(item.ProductId)

		if err != nil {
			return fmt.Errorf("error getting price for product id = %d", item.ProductId)
		}

		item.SetPrice(price)
		item.SetOrderId(orderId)
	}

	updatedStatus, err := s.orderStatusService.GetByName("UPDATED")
	if err != nil {
		return err
	}

	updatedOrder, err := s.repo.UpdateItems(orderId, newItems.Items, updatedStatus.ID)
	if err == nil {
		kafkaMsg := message.OrderCreatedMsg{Event: message.ORDER_UPDATED_EVENT, OrderID: updatedOrder.ID, Currency: updatedOrder.Currency,
			TotalItemsPrice: updatedOrder.TotalItemsPrice.InexactFloat64(), Distance: updatedOrder.Distance}
		_, err = s.kafkaProducer.SentMsg(kafka.ORDER_EVENTS_TOPIC, kafkaMsg)
		if err != nil {
			log.Printf("kafka: failed send order updated msg %s", err.Error())
		}
	}

	return err
}
