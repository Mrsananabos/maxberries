package service

import (
	"catalogService/internal/category/model"
	"catalogService/internal/category/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}

func (s Service) GetById(id int64) (model.Category, error) {
	return s.repo.GetById(id)
}

func (s Service) Create(category model.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}

	return s.repo.Create(category)
}

func (s Service) Update(category model.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}

	return s.repo.Update(category)
}

func (s Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
