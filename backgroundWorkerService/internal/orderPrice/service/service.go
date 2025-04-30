package service

import (
	"backgroundWorkerService/http/rest/client"
	"backgroundWorkerService/internal/deliveryTariffGrid/repository"
	"backgroundWorkerService/internal/deliveryTariffGrid/service"
	usdRates "backgroundWorkerService/internal/usdRates/service"
	"backgroundWorkerService/pkg/db/kafka/message"
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"log"
)

type Service struct {
	httpClient            client.HttpClient
	deliveryTariffService service.Service
	usdRatesService       usdRates.Service
}

func NewService(delTariffService service.Service, usdRatesService usdRates.Service, client client.HttpClient) Service {
	return Service{
		deliveryTariffService: delTariffService,
		usdRatesService:       usdRatesService,
		httpClient:            client,
	}
}

func (s Service) UpdateTotalOrderPrice(ctx context.Context, msg message.OrderMessage) error {
	var editOrderInfo client.OrderPriceInfo

	deliveryPrice, err := s.deliveryTariffService.GetPriceByDistance(msg.Distance)
	if err != nil {
		var notFoundErr repository.DeliveryTariffNotFoundError
		if errors.As(err, &notFoundErr) {
			editOrderInfo = client.OrderPriceInfo{
				Status: "OUTSIDE_DELIVERY",
			}
			log.Printf(notFoundErr.Error())
		} else {
			log.Printf("Failed to get delivery price for orderId %d: %v", msg.OrderID, err)
			return err
		}
	} else {
		rate, err := s.usdRatesService.GetUSDRate(ctx, msg.Currency)
		if err != nil {
			return err
		}

		deliveryPriceInRate := decimal.NewFromFloat(deliveryPrice).Mul(decimal.NewFromFloat(rate))
		totalPrice := decimal.NewFromFloat(msg.TotalItemsPrice).Add(deliveryPriceInRate)

		editOrderInfo = client.OrderPriceInfo{
			Status:        "READY_TO_PAYMENT",
			TotalPrice:    totalPrice.InexactFloat64(),
			DeliveryPrice: deliveryPriceInRate.InexactFloat64(),
		}
	}

	return s.httpClient.UpdateOrder(msg.OrderID, editOrderInfo)
}
