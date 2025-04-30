package service

import (
	"log"
	"orderService/internal/orderStatus/model"
	"orderService/internal/orderStatus/repository"
)

var cachedStatus = make(map[string]*model.OrderStatus)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetStatus(name string) (model.OrderStatus, error) {
	log.Println("statuses")
	for k, v := range cachedStatus {
		log.Printf("key = %s, value= %v", k, *v)
	}
	status, ok := cachedStatus[name]
	if ok {
		return *status, nil
	}

	createdStatus, err := s.GetByName(name)
	if err != nil {
		return model.OrderStatus{}, err
	}
	cachedStatus[name] = &createdStatus
	return createdStatus, nil
}

func (s Service) GetByName(name string) (model.OrderStatus, error) {
	return s.repo.GetByName(name)
}
