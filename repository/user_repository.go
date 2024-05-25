package repository

import (
	"github.com/revandpratama/go-elearning-api/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := r.db.First(&user, "email = ?",  email).Error

	return &user, err
}