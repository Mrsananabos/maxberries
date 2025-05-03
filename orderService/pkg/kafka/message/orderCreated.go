package message

import (
	"encoding/json"
	"fmt"
)

const ORDER_CREATED_EVENT = "ORDER_CREATED"
const ORDER_UPDATED_EVENT = "ORDER_UPDATED"

type OrderCreatedMsg struct {
	Event           string  `json:"event" valid:"required"`
	OrderID         int64   `json:"order_id" valid:"required"`
	Currency        string  `json:"currency" valid:"required"`
	TotalItemsPrice float64 `json:"total_items_price" valid:"required"`
	Distance        float64 `json:"distance" valid:"required"`
}

func (m OrderCreatedMsg) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return []byte{}, fmt.Errorf("failed unmarshal to JSON %v: %w", m, err)
	}

	return jsonData, nil
}
