package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type OrderItem struct {
	Id        int64   `gorm:"primaryKey" json:"id"`
	OrderId   int64   `json:"order_id"`
	ProductId int64   `json:"product_id" valid:"required"`
	Quantity  int32   `json:"quantity" valid:"required"`
	UnitPrice float64 `json:"unit_price"`
}

func (o *OrderItem) SetPrice(price float64) {
	o.UnitPrice = price
}

func (o *OrderItem) Validate() error {
	valid, err := govalidator.ValidateStruct(o)
	errorsStr := strings.Builder{}

	if !valid {
		var validationErrors govalidator.Errors
		if errors.As(err, &validationErrors) {

			for _, validationError := range validationErrors {
				errorsStr.WriteString(validationError.Error())
				errorsStr.WriteString("\n")
			}
		}

		return fmt.Errorf(errorsStr.String())
	}
	return nil
}
