package service

import (
	"authService/internal/role/service"
	"authService/internal/user/model"
	"authService/internal/user/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo        repository.Repository
	roleService service.Service
}

func NewService(r repository.Repository, roleService service.Service) Service {
	return Service{
		repo:        r,
		roleService: roleService,
	}
}

func (s Service) ValidateUserPassword(user model.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (s Service) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s Service) GetById(id uuid.UUID) (model.User, error) {
	return s.repo.GetById(id)
}

func (s Service) GetByEmail(email string) (model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s Service) Create(user model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	return s.repo.Create(user)
}

func (s Service) Update(id uuid.UUID, user model.UserUpdateForm) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hash)
	}

	return s.repo.Update(id, user)
}

func (s Service) UpdateUserRole(id uuid.UUID, roleId uint64) error {
	if _, err := s.roleService.GetById(roleId); err != nil {
		return err
	}

	return s.repo.UpdateUserRole(id, roleId)
}

func (s Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
