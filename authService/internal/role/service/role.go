package service

import (
	"authService/internal/role/model"
	"authService/internal/role/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetAll() ([]model.Role, error) {
	return s.repo.GetAll()
}

func (s Service) GetById(id uint64) (model.Role, error) {
	return s.repo.GetById(id)
}

func (s Service) Create(role model.Role) error {
	if err := role.Validate(); err != nil {
		return err
	}

	return s.repo.Create(role)
}

func (s Service) Update(role model.Role) error {
	if err := role.Validate(); err != nil {
		return err
	}

	return s.repo.Update(role)
}

func (s Service) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s Service) GetRolePermCodesById(id uint64) ([]string, error) {
	permission, err := s.GetById(id)
	if err != nil {
		return []string{}, err
	}

	codes := make([]string, 0, len(permission.Permissions))
	for _, perm := range permission.Permissions {
		codes = append(codes, perm.Code)
	}

	return codes, nil
}
