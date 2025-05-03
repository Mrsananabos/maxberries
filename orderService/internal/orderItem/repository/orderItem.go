package repository

import (
	"fmt"
	"gorm.io/gorm"
	order "orderService/internal/order/model"
	"orderService/internal/orderItem/model"
	"orderService/internal/totalPriceOrder/service"
)

type Repository struct {
	DB                     *gorm.DB
	totalPriceServiceOrder service.Service
}

func NewRepository(db *gorm.DB, totalPriceServiceOrder service.Service) Repository {
	return Repository{DB: db,
		totalPriceServiceOrder: totalPriceServiceOrder,
	}
}

func (r Repository) UpdateItems(orderId int64, items []*model.OrderItem, statusId int64) (order.Order, error) {
	tx := r.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var foundOrder order.Order
	if err := tx.Take(&foundOrder, orderId).Error; err != nil {
		tx.Rollback()
		return foundOrder, fmt.Errorf("not found order with id = %d", orderId)
	}

	totalItemsPrice, err := r.totalPriceServiceOrder.GenerateTotalItemsPrice(items, foundOrder.Currency)
	if err != nil {
		return foundOrder, err
	}

	if err = tx.Where("order_id = ?", orderId).Delete(&model.OrderItem{}).Error; err != nil {
		tx.Rollback()
		return foundOrder, fmt.Errorf("failed to clear existing items for order with id = %d: %v", orderId, err)
	}

	foundOrder.Items = items
	foundOrder.TotalItemsPrice = totalItemsPrice
	foundOrder.StatusID = statusId

	if err = tx.Save(&foundOrder).Error; err != nil {
		tx.Rollback()
		return foundOrder, fmt.Errorf("failed to update order with id = %d: %v", orderId, err)
	}

	if err = tx.Commit().Error; err != nil {
		return foundOrder, err
	}

	return foundOrder, nil
}
