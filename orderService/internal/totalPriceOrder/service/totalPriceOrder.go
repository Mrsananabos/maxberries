package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"orderService/http/rest/client"
	"orderService/internal/orderItem/model"
)

type Service struct {
	httpClient client.HttpClient
}

func NewService(httpClient client.HttpClient) Service {
	return Service{
		httpClient: httpClient,
	}
}

func (s Service) GenerateTotalItemsPrice(items []*model.OrderItem, currency string) (decimal.Decimal, error) {
	usdRate, err := s.httpClient.GetUsdRate(currency)
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
