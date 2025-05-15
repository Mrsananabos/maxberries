package servicesStorage

import (
	"gorm.io/gorm"
	"net/http"
	"orderService/configs"
	"orderService/http/rest/client/auth"
	"orderService/http/rest/client/product"
	"orderService/http/rest/client/rates"
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
	AuthHttpClient     auth.HttpClient
}

func NewServicesStorage(config configs.Config, db *gorm.DB, producer kafka.Producer) ServicesStorage {
	productHttpClient := product.NewHttpClient(config.Services)
	usdRateHttpClient := rates.NewHttpClient(config.Services)
	authHttpClient := auth.NewHttpClient(config.Services, http.DefaultClient)
	totalPriceServiceOrder := totalPriceServOrder.NewService(usdRateHttpClient)
	orderStatusService := orderStatusServ.NewService(orderStatusRepo.NewRepository(db))
	orderService := order.NewService(orderRepo.NewRepository(db), productHttpClient, orderStatusService, totalPriceServiceOrder, producer)
	orderItemsService := orderItemsServ.NewService(orderItemsRepo.NewRepository(db, totalPriceServiceOrder), orderStatusService, productHttpClient, producer)

	return ServicesStorage{
		OrderStatusService: orderStatusService,
		OrderService:       orderService,
		OrderItemsService:  orderItemsService,
		AuthHttpClient:     authHttpClient,
	}

}
