package repository

import (
	"catalogService/internal/category/model"
	"catalogService/pkg/db"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	rsl := r.DB.Find(&categories)

	if rsl.Error != nil {
		return categories, db.HandleError(rsl.Error)
	}

	return categories, nil
}

func (r Repository) GetById(id int64) (model.Category, error) {
	var category model.Category
	rsl := r.DB.Take(&category, id)

	if rsl.Error != nil {
		return category, fmt.Errorf("not found category with id = %d", id)
	}

	return category, nil
}

func (r Repository) Create(category model.Category) error {
	rsl := r.DB.Create(&category)

	if rsl.Error != nil {
		return rsl.Error
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	rsl := r.DB.Delete(&model.Category{}, id)

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found category with id = %d", id)
	}

	return nil
}

func (r Repository) Update(category model.Category) (err error) {
	rsl := r.DB.Model(&model.Category{}).Where("id = ?", category.Id).Updates(category)

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found category with id = %d", category.Id)
	}

	return
}
