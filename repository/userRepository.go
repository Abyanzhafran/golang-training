package repository

import (
	"time"

	"golang-advance/entity"
	"golang-advance/service"
)

type UserRepository struct {
	db     []entity.User
	nextID int
}

func NewUserRepository(db []entity.User) service.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entity.User) entity.User {
	user.ID = r.nextID         
	r.nextID++                 
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.db = append(r.db, *user) 
	return *user
}

func (r *UserRepository) GetUserByID(id int) (entity.User, bool) {
	for _, user := range r.db {
		if user.ID == id {
			return user, true
		}
	}
	return entity.User{}, false
}

func (r *UserRepository) UpdateUser(id int, user entity.User) (entity.User, bool) {
	for i, u := range r.db {
		if u.ID == id {
			user.ID = id                 
			user.CreatedAt = u.CreatedAt 
			user.UpdatedAt = time.Now()  
			r.db[i] = user               
			return user, true            
		}
	}
	return entity.User{}, false 
}


func (r *UserRepository) DeleteUser(id int) bool {
	for i, user := range r.db {
		if user.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...) 
			return true                            
		}
	}
	return false 
}

func (r *UserRepository) GetAllUsers() []entity.User {
	return r.db 
}
