package service

import (
	"fmt"
	"log"
	"orderService/http/rest/client"
	"orderService/internal/order/model"
	"orderService/internal/order/repository"
	orderStatusServ "orderService/internal/orderStatus/service"
	totalPriceOrderServ "orderService/internal/totalPriceOrder/service"
	"orderService/pkg/kafka"
	"orderService/pkg/kafka/message"
	"strings"
)

type Service struct {
	repo                   repository.Repository
	orderStatusService     orderStatusServ.Service
	totalPriceOrderService totalPriceOrderServ.Service
	httpClient             client.HttpClient
	kafkaProducer          kafka.Producer
}

func NewService(r repository.Repository, httpClient client.HttpClient, orderStatusServ orderStatusServ.Service,
	totalPriceOrderService totalPriceOrderServ.Service, kafka kafka.Producer) Service {
	return Service{
		repo:                   r,
		orderStatusService:     orderStatusServ,
		totalPriceOrderService: totalPriceOrderService,
		httpClient:             httpClient,
		kafkaProducer:          kafka,
	}
}

func (s Service) GetAll() ([]model.Order, error) {
	return s.repo.GetAll()
}

func (s Service) GetById(id int64) (model.Order, error) {
	return s.repo.GetById(id)
}

func (s Service) Create(order model.Order) (model.Order, error) {
	for _, item := range order.Items {
		price, err := s.httpClient.GetProductPrice(item.ProductId)
		if err != nil {
			return model.Order{}, fmt.Errorf("error getting price for product id = %d", item.ProductId)
		}

		item.SetPrice(price)
	}

	totalItemsPrice, err := s.totalPriceOrderService.GenerateTotalItemsPrice(order.Items, order.Currency)
	if err != nil {
		return model.Order{}, err
	}
	order.SetTotalItemsPrice(totalItemsPrice)

	statusCreated, err := s.orderStatusService.GetStatus("CREATED")
	if err != nil {
		return model.Order{}, err
	}

	order.SetStatus(statusCreated)
	createdOrder, err := s.repo.Create(order)
	if err != nil {
		return model.Order{}, err
	}

	kafkaMsg := message.OrderCreatedMsg{Event: message.ORDER_CREATED_EVENT, OrderID: createdOrder.ID, Currency: createdOrder.Currency,
		TotalItemsPrice: createdOrder.TotalItemsPrice.InexactFloat64(), Distance: createdOrder.Distance}
	_, err = s.kafkaProducer.SentMsg(kafka.ORDER_EVENTS_TOPIC, kafkaMsg)
	if err != nil {
		log.Printf("kafka: failed send order created msg %s", err.Error())
	}

	return createdOrder, nil
}

func (s Service) UpdateOrder(id int64, editOrderReq model.EditOrderRequest) (model.Order, error) {
	var updatedOrder model.Order
	updatedOrderFields := make(map[string]interface{})

	if editOrderReq.Status != nil {
		newStatus, err := s.orderStatusService.GetByName(strings.ToUpper(*editOrderReq.Status))
		if err != nil {
			return updatedOrder, err
		}

		updatedOrderFields["status_id"] = newStatus.ID
	}

	if editOrderReq.TotalPrice != nil {
		updatedOrderFields["total_price"] = editOrderReq.TotalPrice
	}

	if editOrderReq.TotalPrice != nil {
		updatedOrderFields["delivery_price"] = editOrderReq.DeliveryPrice
	}

	if editOrderReq.Distance != nil {
		updatedOrderFields["distance"] = editOrderReq.Distance
	}

	if len(updatedOrderFields) == 0 {
		return updatedOrder, fmt.Errorf("no fields for update")
	}

	_, ok := updatedOrderFields["distance"]

	if ok {
		updatedStatus, err := s.orderStatusService.GetByName("UPDATED")
		if err != nil {
			return updatedOrder, err
		}

		updatedOrderFields["status_id"] = updatedStatus.ID
	}

	updatedOrder, err := s.repo.Update(id, updatedOrderFields)
	if err == nil {
		if ok {
			kafkaMsg := message.OrderCreatedMsg{Event: message.ORDER_UPDATED_EVENT, OrderID: updatedOrder.ID, Currency: updatedOrder.Currency,
				TotalItemsPrice: updatedOrder.TotalItemsPrice.InexactFloat64(), Distance: updatedOrder.Distance}
			_, err = s.kafkaProducer.SentMsg(kafka.ORDER_EVENTS_TOPIC, kafkaMsg)
			if err != nil {
				log.Printf("kafka: failed send order updated msg %s", err.Error())
			}
		}
	}

	return updatedOrder, err
}

func (s Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
