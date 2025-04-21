package service

import (
	"github.com/shopspring/decimal"
	"orderService/internal/order/model"
	"orderService/internal/order/repository"
	orderItem "orderService/internal/orderItem/model"
	orderStatus "orderService/internal/orderStatus/model"
	orderStatusService "orderService/internal/orderStatus/service"
	"strings"
)

type Service struct {
	repo               repository.Repository
	orderStatusService orderStatusService.Service
}

func NewService(r repository.Repository, orderStatusService orderStatusService.Service) Service {
	return Service{
		repo:               r,
		orderStatusService: orderStatusService,
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
	statusCreated, err := s.orderStatusService.GetCreatedStatus()
	if err != nil {
		return err
	}

	order.SetStatus(statusCreated)
	return s.repo.Create(order)
}

func (s Service) UpdateStatus(id int64, status orderStatus.EditOrderStatusRequest) error {
	if err := status.Validate(); err != nil {
		return err
	}

	newStatus, err := s.orderStatusService.GetByName(strings.ToUpper(status.Status))
	if err != nil {
		return err
	}

	return s.repo.UpdateStatus(id, newStatus)
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
