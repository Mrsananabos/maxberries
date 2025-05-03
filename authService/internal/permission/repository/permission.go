package repository

import (
	"authService/internal/permission/model"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetAll() ([]model.Permission, error) {
	var permissions []model.Permission
	rsl := r.DB.Find(&permissions)

	if rsl.Error != nil {
		return permissions, rsl.Error
	}

	return permissions, nil
}

func (r Repository) GetById(id uint64) (model.Permission, error) {
	var permission model.Permission
	rsl := r.DB.Take(&permission, id)

	if rsl.Error != nil {
		return permission, rsl.Error
	}

	return permission, nil
}

func (r Repository) Create(permission model.Permission) error {
	rsl := r.DB.Create(&permission)

	if rsl.Error != nil {
		return rsl.Error
	}

	return nil
}

func (r Repository) Delete(id uint64) error {
	rsl := r.DB.Delete(&model.Permission{}, id)

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found permission with id = %d", id)
	}

	return nil
}

func (r Repository) Update(permission model.Permission) (err error) {
	rsl := r.DB.Model(&model.Permission{}).Where("id = ?", permission.ID).Updates(permission)

	if rsl.RowsAffected == 0 {
		return fmt.Errorf("not found permission with id = %d", permission.ID)
	}

	return
}
