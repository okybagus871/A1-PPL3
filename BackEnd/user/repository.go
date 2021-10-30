package user

import "gorm.io/gorm"

type Repository interface {
	GetUser() ([]User, error)
	CreateUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error

	return user, err
}

func (r *repository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
