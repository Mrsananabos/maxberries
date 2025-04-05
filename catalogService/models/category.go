package models

import (
	"time"
)

type Category struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:text" json:"name" validate:"required,min=1"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
