package model

import (
	"github.com/asaskevich/govalidator"
	_ "github.com/asaskevich/govalidator"
)

type OrderStatus struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type EditOrderStatusRequest struct {
	Status string `valid:"required" json:"status"`
}

func (s *OrderStatus) TableName() string {
	return "statuses"
}

func (s EditOrderStatusRequest) Validate() error {
	valid, err := govalidator.ValidateStruct(s)
	if !valid {
		return err
	}
	return nil
}
