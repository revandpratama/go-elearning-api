package repository

import (
	"errors"

	"github.com/revandpratama/go-elearning-api/model"
	"gorm.io/gorm"
)

type krsRepository struct {
	db *gorm.DB
}

type KRSRepository interface {
	Create(krs *model.KRS) error
	Update(id int, krs *model.KRS) error
	Delete(id int) error
	GetAll() (*[]model.KRS, error)
	GetById(id int) (*model.KRS, error)
	GetByUserID(id int) (*[]model.KRS, error)
}

func NewKRSRepository(db *gorm.DB) *krsRepository {
	return &krsRepository{
		db: db,
	}
}

func (r *krsRepository) Create(krs *model.KRS) error {
	err := r.db.Create(&krs).Error
	return err
}

func (r *krsRepository) Update(id int, krs *model.KRS) error {
	err := r.db.Where("id = ?", id).Updates(&krs).Error
	return err
}

func (r *krsRepository) Delete(id int) error {
	err := r.db.Delete(&model.KRS{}, id).Error
	return err
}

func (r *krsRepository) GetAll() (*[]model.KRS, error) {
	var krs []model.KRS
	err := r.db.Find(&krs).Error
	return &krs, err
}

func (r *krsRepository) GetById(id int) (*model.KRS, error) {
	var krs model.KRS
	err := r.db.Where("id = ?", id).First(&krs).Error
	return &krs, err
}

func (r *krsRepository) GetByUserID(id int) (*[]model.KRS, error) {
	var krs []model.KRS
	err := r.db.Find(&krs, "user_id = ?", id).Error
	if len(krs) < 1 {
		return nil, errors.New("record not found")
	}
	return &krs, err
}
