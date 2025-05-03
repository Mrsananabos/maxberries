package model

import (
	"encoding/json"
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
	ID              int64              `gorm:"primaryKey" json:"id"`
	UserId          uuid.UUID          `json:"user_id" valid:"required"`
	TotalPrice      decimal.Decimal    `json:"total_price"`
	TotalItemsPrice decimal.Decimal    `json:"total_items_price"`
	DeliveryPrice   decimal.Decimal    `json:"delivery_price"`
	Currency        string             `json:"currency" valid:"required"`
	StatusID        int64              `json:"status_id"`
	Status          status.OrderStatus `gorm:"foreignKey:StatusID;references:ID" json:"status"`
	Items           []*model.OrderItem `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE;" json:"items" valid:"required"`
	Distance        float64            `json:"distance" valid:"required"`
	CreatedAt       time.Time          `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

type EditOrderRequest struct {
	TotalPrice    *float64 `json:"total_price"`
	Status        *string  `json:"status"`
	DeliveryPrice *float64 `json:"delivery_price"`
	Distance      *float64 `json:"distance"`
}

func (o *Order) SetTotalItemsPrice(price decimal.Decimal) {
	o.TotalItemsPrice = price
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

func (o Order) MarshalJSON() ([]byte, error) {
	type OrderFields Order
	return json.Marshal(&struct {
		*OrderFields
		TotalPrice      float64 `json:"total_price"`
		TotalItemsPrice float64 `json:"total_items_price"`
		DeliveryPrice   float64 `json:"delivery_price"`
	}{
		OrderFields:     (*OrderFields)(&o),
		TotalPrice:      o.TotalPrice.InexactFloat64(),
		TotalItemsPrice: o.TotalItemsPrice.InexactFloat64(),
		DeliveryPrice:   o.DeliveryPrice.InexactFloat64(),
	})
}
