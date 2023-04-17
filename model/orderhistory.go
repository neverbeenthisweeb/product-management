package model

import "time"

type OrderHistory struct {
	ID           int
	UserID       int
	OrderItemID  int
	Descriptions string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func (OrderHistory) TableName() string {
	return "orders_histories"
}
