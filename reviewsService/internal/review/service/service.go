package service

import (
	"reviewsService/internal/review/model"
	"reviewsService/internal/review/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetByProductId(id string) ([]model.Review, error) {
	reviews, err := s.repo.GetByProductId(id)
	if err != nil {
		return []model.Review{}, err
	}

	return reviews, nil
}

func (s Service) Create(review model.Review) error {
	//if err := category.Validate(); err != nil {
	//	return err
	//}

	return s.repo.Create(review)
}

func (s Service) Update(review model.Review) error {
	//if err := category.Validate(); err != nil {
	//	return err
	//}

	return s.repo.Update(review)
}

func (s Service) DeleteByProductId(id string) error {
	return s.repo.DeleteByProductId(id)
}

func (s Service) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}
