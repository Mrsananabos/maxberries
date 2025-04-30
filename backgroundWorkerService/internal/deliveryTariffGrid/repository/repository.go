package repository

import (
	"backgroundWorkerService/internal/deliveryTariffGrid/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type DeliveryTariffNotFoundError struct {
	Distance float64
}

func (e DeliveryTariffNotFoundError) Error() string {
	return fmt.Sprintf("delivery tariff for %f distance not found", e.Distance)
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetPriceByDistance(distance float64) (model.DeliveryTariff, error) {
	var delTariff model.DeliveryTariff
	result := r.DB.Where("max_distance >= ?", distance).Order("max_distance").First(&delTariff)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.DeliveryTariff{}, DeliveryTariffNotFoundError{Distance: distance}
		}
		return model.DeliveryTariff{}, result.Error
	}

	return delTariff, nil
}
