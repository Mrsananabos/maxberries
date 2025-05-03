package service

import (
	"authService/internal/permission/model"
	"authService/internal/permission/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetAll() ([]model.Permission, error) {
	return s.repo.GetAll()
}

func (s Service) GetById(id uint64) (model.Permission, error) {
	return s.repo.GetById(id)
}

func (s Service) Create(permission model.Permission) error {
	if err := permission.Validate(); err != nil {
		return err
	}

	return s.repo.Create(permission)
}

func (s Service) Update(permission model.Permission) error {
	if err := permission.Validate(); err != nil {
		return err
	}

	return s.repo.Update(permission)
}

func (s Service) Delete(id uint64) error {
	return s.repo.Delete(id)
}
