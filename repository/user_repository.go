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
	GetUserByUsername(username string) (*model.User, error)
	Create(user *model.User) error
	Update(id int, user *model.User) error
	Delete(id int) error
	GetAll() (*[]model.User, error)
	GetById(id int) (*model.User, error)
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
func (r *userRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User

	err := r.db.First(&user, "username = ?",  username).Error

	return &user, err
}

func (r *userRepository) Create(user *model.User) error {
	err := r.db.Create(&user).Error
	return err
}

func (r *userRepository) Update(id int, user *model.User) error {
	err := r.db.Where("id = ?", id).Updates(&user).Error
	return err
}

func (r *userRepository) Delete(id int) error {
	err := r.db.Delete(&model.User{}, id).Error
	return err
}

func (r *userRepository) GetAll() (*[]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return &users, err
}

func (r *userRepository) GetById(id int) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

// func (r *userRepository) UserAuthorized(id int) (error) {
// 	var user model.User
// 	err := r.db.First(&user, "id = ?", id).Error
// 	return &user, err
// }