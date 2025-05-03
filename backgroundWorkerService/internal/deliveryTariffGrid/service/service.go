package service

import (
	"backgroundWorkerService/internal/deliveryTariffGrid/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetPriceByDistance(distance float64) (float64, error) {
	tariff, err := s.repo.GetPriceByDistance(distance)
	if err != nil {
		return 0, err
	}

	return tariff.Price, nil
}
