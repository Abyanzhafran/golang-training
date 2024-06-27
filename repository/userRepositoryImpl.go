package repository

import (
	"context"

	"golang-assignment/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.User, error) {
	var user *entity.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}	

	return user, nil
}