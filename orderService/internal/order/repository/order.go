package repository

import (
	"fmt"
	"gorm.io/gorm"
	"orderService/internal/order/model"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	rsl := r.DB.Preload("Items").Preload("Status").Find(&orders)

	if rsl.Error != nil {
		return orders, rsl.Error
	}

	return orders, nil
}

func (r Repository) GetById(id int64) (model.Order, error) {
	var order model.Order
	rsl := r.DB.Preload("Items").Preload("Status").Take(&order, id)

	if rsl.Error != nil {
		return order, fmt.Errorf("not found order with id = %d", id)
	}

	return order, nil
}

func (r Repository) Create(order model.Order) (model.Order, error) {
	rsl := r.DB.Create(&order)

	if rsl.Error != nil {
		return model.Order{}, rsl.Error
	}

	return order, nil
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

func (r Repository) Update(id int64, updatedFields map[string]interface{}) (model.Order, error) {
	var updatedOrder model.Order
	rsl := r.DB.Model(&updatedOrder).Where("id = ?", id).Updates(updatedFields).Scan(&updatedOrder)

	if rsl.Error != nil {
		return model.Order{}, rsl.Error
	}

	if rsl.RowsAffected != 1 {
		return model.Order{}, fmt.Errorf("not found order with id = %d", id)
	}

	return updatedOrder, nil
}
