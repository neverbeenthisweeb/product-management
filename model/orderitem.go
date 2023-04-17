package model

import "time"

type OrderItem struct {
	ID        int
	Name      string
	Price     float64
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (OrderItem) TableName() string {
	return "orders_items"
}
