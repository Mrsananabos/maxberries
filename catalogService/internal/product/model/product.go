package model

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type Product struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:text" json:"name" valid:"required"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float32   `json:"price" valid:"required"`
	CategoryId  int64     `json:"category_id" valid:"required"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (p Product) Validate() error {
	valid, err := govalidator.ValidateStruct(p)
	errorsStr := strings.Builder{}

	validPriceErrors := p.isValidPrice()
	if !valid || len(validPriceErrors) != 0 {
		var validationErrors govalidator.Errors
		if errors.As(err, &validationErrors) {

			for _, validationError := range validationErrors {
				errorsStr.WriteString(validationError.Error())
				errorsStr.WriteString("\n")
			}
		}

		for _, validationError := range validPriceErrors {
			errorsStr.WriteString(validationError.Error())
			errorsStr.WriteString("\n")
		}
		return fmt.Errorf(errorsStr.String())
	}
	return nil
}

func (p Product) isValidPrice() []error {
	rsl := make([]error, 0, 2)
	decimalPrice := decimal.NewFromFloat32(p.Price)

	if decimalPrice.LessThan(decimal.Zero) {
		rsl = append(rsl, fmt.Errorf("price: a positive number is required"))
	}

	countNumbersAfterDot := -decimalPrice.Exponent()
	if countNumbersAfterDot > 2 {
		rsl = append(rsl, fmt.Errorf("price: max 2 numbers after point"))
	}

	return rsl
}
