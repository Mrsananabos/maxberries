package servicesStorage

import (
	"gorm.io/gorm"
	"orderService/configs"
	"orderService/http/rest/client"
	orderRepo "orderService/internal/order/repository"
	order "orderService/internal/order/service"
	orderItemsRepo "orderService/internal/orderItem/repository"
	orderItemsServ "orderService/internal/orderItem/service"
	orderStatusRepo "orderService/internal/orderStatus/repository"
	orderStatusServ "orderService/internal/orderStatus/service"
	totalPriceServOrder "orderService/internal/totalPriceOrder/service"
	"orderService/pkg/kafka"
)

type ServicesStorage struct {
	OrderStatusService orderStatusServ.Service
	OrderService       order.Service
	OrderItemsService  orderItemsServ.Service
	HttpClient         client.HttpClient
}

func NewServicesStorage(config configs.Config, db *gorm.DB, producer kafka.Producer) ServicesStorage {
	httpClient := client.NewHttpClient(config.Services)
	totalPriceServiceOrder := totalPriceServOrder.NewService(httpClient)
	orderStatusService := orderStatusServ.NewService(orderStatusRepo.NewRepository(db))
	orderService := order.NewService(orderRepo.NewRepository(db), httpClient, orderStatusService, totalPriceServiceOrder, producer)
	orderItemsService := orderItemsServ.NewService(orderItemsRepo.NewRepository(db, totalPriceServiceOrder), orderStatusService, httpClient, producer)

	return ServicesStorage{
		OrderStatusService: orderStatusService,
		OrderService:       orderService,
		OrderItemsService:  orderItemsService,
		HttpClient:         httpClient,
	}

}
