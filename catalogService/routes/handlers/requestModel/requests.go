package requestModel

import (
	"github.com/shopspring/decimal"
)

type SaveProductRequest struct {
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price" validate:"required"`
	CategoryId  int64           `json:"category_id"`
}
