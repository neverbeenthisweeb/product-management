package gormrepo

import (
	"context"
	"fmt"
	"productmanagement/model"
	"productmanagement/service"

	"gorm.io/gorm"
)

type userImpl struct {
	db *gorm.DB
}

func NewUserImpl(db *gorm.DB) *userImpl {
	return &userImpl{
		db: db,
	}
}

func (r *userImpl) Get(ctx context.Context, ID string) (model.User, error) {
	u := model.User{}

	if err := r.db.First(&u, ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, service.NewErrNotFound(fmt.Sprintf("user ID %s not found", ID))
		}

		return model.User{}, err
	}

	return u, nil
}
