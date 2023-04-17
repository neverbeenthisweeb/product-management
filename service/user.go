package service

import (
	"context"
	"productmanagement/config"
	"productmanagement/model"
	"productmanagement/repository"
)

type User interface {
	// Create()
	Get(ctx context.Context, ID string) (model.User, error)
	// List()
	// Update()
	// Delete()
}

type userImpl struct {
	cfg      config.Config
	repoUser repository.User
}

func NewUserImpl(cfg config.Config, repo *repository.Repository) *userImpl {
	return &userImpl{
		cfg:      cfg,
		repoUser: repo.User,
	}
}

func (s *userImpl) Get(ctx context.Context, ID string) (model.User, error) {
	return s.repoUser.Get(ctx, ID)
}
