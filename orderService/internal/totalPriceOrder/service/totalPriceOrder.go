package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"orderService/http/rest/client/rates"
	"orderService/internal/orderItem/model"
)

type Service struct {
	usdRateHttpClient rates.HttpClient
}

func NewService(httpClient rates.HttpClient) Service {
	return Service{
		usdRateHttpClient: httpClient,
	}
}

func (s Service) GenerateTotalItemsPrice(items []*model.OrderItem, currency string) (decimal.Decimal, error) {
	usdRate, err := s.usdRateHttpClient.GetUsdRate(currency)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed get rate for %s: %w", currency, err)
	}

	totalItemsPrice := decimal.NewFromInt(0)
	rate := decimal.NewFromFloat(usdRate)
	for _, item := range items {
		itemsPrice := decimal.NewFromFloat(item.UnitPrice).Mul(rate).Mul(decimal.NewFromInt32(item.Quantity))
		totalItemsPrice = totalItemsPrice.Add(itemsPrice)
	}

	return totalItemsPrice.Round(2), nil
}
