package models

import (
	"time"
)

type Product struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:text" json:"name" validate:"required,min=1"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float32   `gorm:"type:decimal(10,2)" json:"price" validate:"required,min=0""`
	CategoryId  int64     `json:"category_id" validate:"required"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
