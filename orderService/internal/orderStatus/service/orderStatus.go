package service

import (
	"orderService/internal/orderStatus/model"
	"orderService/internal/orderStatus/repository"
)

const CREATED_STATUS_NAME = "CREATED"

var cachedStatusCreated *model.OrderStatus

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetCreatedStatus() (model.OrderStatus, error) {
	if cachedStatusCreated != nil {
		return *cachedStatusCreated, nil
	}

	createdStatus, err := s.GetByName(CREATED_STATUS_NAME)
	cachedStatusCreated = &createdStatus
	if err != nil {
		return model.OrderStatus{}, err
	}

	return createdStatus, nil
}

func (s Service) GetByName(name string) (model.OrderStatus, error) {
	return s.repo.GetByName(name)
}
