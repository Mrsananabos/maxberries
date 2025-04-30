package message

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type OrderMessage struct {
	Event           string  `json:"event" valid:"required"`
	OrderID         int64   `json:"order_id" valid:"required"`
	Currency        string  `json:"currency" valid:"required"`
	TotalItemsPrice float64 `json:"total_items_price" valid:"required"`
	Distance        float64 `json:"distance" valid:"required"`
}

func (c OrderMessage) Validate() error {
	valid, err := govalidator.ValidateStruct(c)
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
	return nil
}
