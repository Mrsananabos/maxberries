package repository

import (
	"fmt"
	"gorm.io/gorm"
	"orderService/internal/order/model"
	"strings"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	rsl := r.DB.Preload("Items").Find(&orders)

	if rsl.Error != nil {
		return orders, rsl.Error
	}

	return orders, nil
}

func (r Repository) GetById(id int64) (model.Order, error) {
	var order model.Order
	rsl := r.DB.Preload("Items").Take(&order, id)

	if rsl.Error != nil {
		return order, fmt.Errorf("not found order with id = %d", id)
	}

	return order, nil
}

func (r Repository) Create(order model.Order) error {
	rsl := r.DB.Create(&order)

	if rsl.Error != nil {
		return rsl.Error
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	rsl := r.DB.Delete(&model.Order{}, id)

	if rsl.Error != nil {
		return rsl.Error
	}

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found order with id = %d", id)
	}

	return nil
}

func (r Repository) UpdateStatus(id int64, status model.Status) (err error) {
	rsl := r.DB.Model(&model.Order{}).Where("id = ?", id).Update("status", strings.ToUpper(string(status)))

	if rsl.Error != nil {
		return rsl.Error
	}

	if rsl.RowsAffected != 1 {
		return fmt.Errorf("not found order with id = %d", id)
	}

	return
}
