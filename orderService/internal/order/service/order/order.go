package order

import (
	"github.com/shopspring/decimal"
	"orderService/http/rest/client"
	"orderService/internal/order/model"
	"orderService/internal/order/repository"
	orderItem "orderService/internal/orderItem/model"
)

type Service struct {
	repo       repository.Repository
	httpClient client.HttpClient
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetAll() ([]model.Order, error) {
	return s.repo.GetAll()
}

func (s Service) GetById(id int64) (model.Order, error) {
	return s.repo.GetById(id)
}

func (s Service) Create(order model.Order, usdRate float64) error {
	totalPrice := generateTotalPrice(order.Items, usdRate)
	order.SetTotalPrice(totalPrice)
	order.SetStatus(model.CREATED)

	return s.repo.Create(order)
}

func (s Service) UpdateStatus(id int64, status model.UpdateStatusRequest) error {
	if err := status.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateStatus(id, status.Status)
}

func (s Service) Delete(id int64) error {
	return s.repo.Delete(id)
}

func generateTotalPrice(items []*orderItem.OrderItem, rate float64) decimal.Decimal {
	rsl := decimal.NewFromInt(0)
	rateDec := decimal.NewFromFloat(rate)
	for _, item := range items {
		count := decimal.NewFromFloat(item.UnitPrice).Mul(rateDec).Mul(decimal.NewFromInt32(item.Quantity))
		rsl = rsl.Add(count)
	}

	return rsl.Round(2)
}
