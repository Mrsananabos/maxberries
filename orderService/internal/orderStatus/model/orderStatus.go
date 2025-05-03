package model

type OrderStatus struct {
	ID   int64  `gorm:"primaryKey" json:"-"`
	Name string `gorm:"not null" json:"name"`
}

func (s *OrderStatus) TableName() string {
	return "statuses"
}
