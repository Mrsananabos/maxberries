package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"orderService/internal/orderItem/model"
	status "orderService/internal/orderStatus/model"
	"strings"
	"time"
)

type Order struct {
	ID            int64              `gorm:"primaryKey" json:"id"`
	UserId        uuid.UUID          `json:"user_id" valid:"required"`
	TotalPrice    decimal.Decimal    `json:"total_price"`
	DeliveryPrice float64            `json:"delivery_price"`
	Currency      string             `json:"currency" valid:"required"`
	StatusID      int64              `json:"status_id"`
	Status        status.OrderStatus `gorm:"foreignKey:StatusID;references:ID" json:"status"`
	Items         []*model.OrderItem `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE;" json:"items" valid:"required"`
	CreatedAt     time.Time          `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (o *Order) SetTotalPrice(price decimal.Decimal) {
	o.TotalPrice = price
}

func (o *Order) SetStatus(status status.OrderStatus) {
	o.Status = status
}

func (o *Order) Validate() error {
	valid, err := govalidator.ValidateStruct(o)
	if !valid {
		var validationErrors govalidator.Errors
		if errors.As(err, &validationErrors) {
			errorsStr := strings.Builder{}
			for _, validationError := range validationErrors {
				errorsStr.WriteString(validationError.Error())
				errorsStr.WriteString("\n")
			}
			return fmt.Errorf(errorsStr.String())
		}
	}

	for _, item := range o.Items {
		err = item.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
