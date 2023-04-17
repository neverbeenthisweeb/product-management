package model

import "time"

type User struct {
	ID         int
	FullName   string
	FirstOrder string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func (User) TableName() string {
	return "users"
}
