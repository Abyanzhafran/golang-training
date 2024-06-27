package repository

import (
	"context"

	"golang-assignment/entity"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetById(ctx context.Context, id int64) (*entity.User, error)
	Create(ctx context.Context, product *entity.User) error
}