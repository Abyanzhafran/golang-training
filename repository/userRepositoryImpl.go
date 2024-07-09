package repository

import (
	"context"
	"log"

	"golang-assignment/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) (entity.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.User{}, err
	}
	return *user, nil
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetById(ctx context.Context, id int) (*entity.User, error) {
	var user *entity.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	return r.DB.Model(&entity.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id int) error {
	return r.DB.Where("id = ?", id).Delete(&entity.User{}).Error
}
