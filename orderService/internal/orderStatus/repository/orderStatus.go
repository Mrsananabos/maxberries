package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"orderService/internal/orderStatus/model"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetByName(name string) (model.OrderStatus, error) {
	var order model.OrderStatus
	rsl := r.DB.Model(&model.OrderStatus{}).Where("name = ?", name).First(&order)
	err := r.DB.Model(&model.OrderStatus{}).Where("name = ?", name).First(&order).Error
	log.Println(name)
	log.Printf("%v", rsl.Statement)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, fmt.Errorf("not found status with name = %s", name)
		}
		return order, fmt.Errorf("error retrieving status: %v", err)
	}

	return order, nil
}
