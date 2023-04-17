package repository

import (
	"context"
	"productmanagement/model"
)

type Repository struct {
	User
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SetUser(u User) {
	r.User = u
}

type User interface {
	Get(ctx context.Context, ID string) (model.User, error)
}
