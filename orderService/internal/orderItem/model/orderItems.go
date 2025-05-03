package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type OrderItem struct {
	Id        int64   `gorm:"primaryKey" json:"-"`
	OrderId   int64   `json:"-"`
	ProductId int64   `json:"product_id" valid:"required"`
	Quantity  int32   `json:"quantity" valid:"required"`
	UnitPrice float64 `json:"unit_price"`
}

type EditOrderItemsRequest struct {
	Items []*OrderItem `json:"items" valid:"required"`
}

func (o *OrderItem) SetPrice(price float64) {
	o.UnitPrice = price
}

func (o *OrderItem) SetOrderId(id int64) {
	o.OrderId = id
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

func (e *EditOrderItemsRequest) Validate() error {
	valid, err := govalidator.ValidateStruct(e)
	errorsStr := strings.Builder{}

	if valid {
		for i, item := range e.Items {
			err = item.Validate()
			if err != nil {
				errorsStr.WriteString(fmt.Sprintf("item %d: %s", i, err.Error()))
				errorsStr.WriteString("\n")
			}
		}

		if errorsStr.Len() != 0 {
			return fmt.Errorf(errorsStr.String())
		} else {
			return nil
		}
	}

	return err
}
