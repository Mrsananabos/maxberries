package repository

import (
	"authService/internal/user/model"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetAll() ([]model.User, error) {
	var users []model.User
	rsl := r.DB.Preload("Role").Preload("Role.Permissions").Find(&users)

	if rsl.Error != nil {
		return users, rsl.Error
	}

	return users, nil
}

func (r Repository) GetById(id uuid.UUID) (model.User, error) {
	var user model.User
	rsl := r.DB.Preload("Role").Preload("Role.Permissions").Take(&user, id)

	if rsl.Error != nil {
		return user, fmt.Errorf("not found user with id = %s", id)
	}

	return user, nil
}

func (r Repository) GetByEmail(email string) (model.User, error) {
	var user model.User
	rsl := r.DB.Preload("Role").Where("email=?", email).First(&user)

	if rsl.Error != nil {
		return user, fmt.Errorf("not found user with email = %s", email)
	}

	return user, nil
}

func (r Repository) Create(user model.User) error {
	rsl := r.DB.Create(&user)
	if rsl.Error != nil {
		return rsl.Error
	}

	return nil
}

func (r Repository) Delete(id uuid.UUID) error {
	rsl := r.DB.Delete(&model.User{}, id)

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found user with id = %s", id)
	}

	return nil
}

func (r Repository) Update(id uuid.UUID, user model.UserUpdateForm) (err error) {
	rsl := r.DB.Model(&model.User{}).Where("id = ?", id).Updates(user)

	if rsl.Error != nil {
		return rsl.Error
	}

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found user with id = %s", id)
	}

	return nil
}

func (r Repository) UpdateUserRole(id uuid.UUID, roleId uint64) (err error) {
	rsl := r.DB.Model(&model.User{}).Where("id = ?", id).Update("role_id", roleId)

	if rsl.Error != nil {
		return rsl.Error
	}

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found user with id = %s", id)
	}

	return nil
}
