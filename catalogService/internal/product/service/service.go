package service

import (
	"catalogService/internal/category/service"
	"catalogService/internal/product/model"
	"catalogService/internal/product/repository"
)

type Service struct {
	repo            repository.Repository
	categoryService service.Service
}

func NewService(r repository.Repository, categoryService service.Service) Service {
	return Service{
		repo:            r,
		categoryService: categoryService,
	}
}

func (s Service) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s Service) GetById(id int64) (model.Product, error) {
	return s.repo.GetById(id)
}

func (s Service) Create(product model.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}

	_, err := s.categoryService.GetById(product.CategoryId)
	if err != nil {
		return err
	}

	return s.repo.Create(product)
}

func (s Service) Update(product model.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}

	_, err := s.categoryService.GetById(product.CategoryId)
	if err != nil {
		return err
	}

	return s.repo.Update(product)
}

func (s Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
