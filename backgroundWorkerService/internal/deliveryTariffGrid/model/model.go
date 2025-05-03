package model

type DeliveryTariff struct {
	MaxDistance int64   `gorm:"not null;unique"`
	Price       float64 `gorm:"unique"`
}

func (s *DeliveryTariff) TableName() string {
	return "delivery_tariff"
}
