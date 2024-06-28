package repository

import (
	"context"

	"golang-assignment/entity"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetById(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context, product *entity.User) error
	Update(ctx context.Context, product *entity.User) error
	Delete(ctx context.Context, id int) error
}
