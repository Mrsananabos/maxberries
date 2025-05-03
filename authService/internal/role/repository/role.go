package repository

import (
	"authService/internal/role/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetAll() ([]model.Role, error) {
	var roles []model.Role
	rsl := r.DB.Preload("Permissions").Find(&roles)

	if rsl.Error != nil {
		return roles, rsl.Error
	}

	return roles, nil
}

func (r Repository) GetById(id uint64) (model.Role, error) {
	var role model.Role
	rsl := r.DB.Preload("Permissions").Take(&role, id)

	if errors.Is(gorm.ErrRecordNotFound, rsl.Error) {
		return role, fmt.Errorf("not found role with id = %d", id)
	}

	if rsl.Error != nil {
		return role, rsl.Error
	}

	return role, nil
}

func (r Repository) Create(role model.Role) error {
	rsl := r.DB.Create(&role)

	if rsl.Error != nil {
		return rsl.Error
	}

	return nil
}

func (r Repository) Delete(id uint64) error {
	rsl := r.DB.Delete(&model.Role{}, id)

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found role with id = %d", id)
	}

	return nil
}

func (r Repository) Update(role model.Role) (err error) {
	rsl := r.DB.Model(&model.Role{}).Where("id = ?", role.ID).Updates(role)

	if rsl.RowsAffected == 0 {
		return fmt.Errorf("not found role with id = %d", role.ID)
	}

	return
}
