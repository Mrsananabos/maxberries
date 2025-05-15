package message

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const REVIEW_CREATED_EVENT = "REVIEW_CREATED"

type ReviewCreatedMsg struct {
	Event     string             `json:"event"`
	ID        primitive.ObjectID `json:"id"`
	ProductID int64              `json:"product_id"`
	UserID    string             `json:"user_id"`
	Rating    *int               `json:"rating"`
	Text      string             `json:"text"`
}

func (m ReviewCreatedMsg) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return []byte{}, fmt.Errorf("failed unmarshal to JSON %v: %w", m, err)
	}

	return jsonData, nil
}
