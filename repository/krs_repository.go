package repository

import (
	"errors"

	"github.com/revandpratama/go-elearning-api/dto"
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
	GetAll(paginate *dto.Paginate) (*[]model.KRS, error)
	GetById(id int) (*model.KRS, error)
	GetByUserID(userid int, paginate *dto.Paginate) (*[]model.KRS, error)
	GetTotalData() (int64, error)
	GetTotalDataByUserId(userid int) (int64, error)
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

func (r *krsRepository) GetAll(paginate *dto.Paginate) (*[]model.KRS, error) {
	var krs []model.KRS
	err := r.db.Offset((paginate.CurrentPage - 1) * 10).Limit(paginate.DataPerPage).Find(&krs).Error
	if len(krs) < 1 {
		return nil, errors.New("record not found")
	}

	return &krs, err
}

func (r *krsRepository) GetTotalData() (int64, error) {
	var count int64
	err := r.db.Model(&model.KRS{}).Count(&count).Error
	return count, err
}

func (r *krsRepository) GetById(id int) (*model.KRS, error) {
	var krs model.KRS
	err := r.db.Where("id = ?", id).First(&krs).Error
	return &krs, err
}

func (r *krsRepository) GetByUserID(userid int, paginate *dto.Paginate) (*[]model.KRS, error) {
	var krs []model.KRS
	err := r.db.Offset((paginate.CurrentPage - 1) * 10).Limit(paginate.DataPerPage).Find(&krs, "user_id", userid).Error
	if len(krs) < 1 {
		return nil, errors.New("record not found")
	}
	return &krs, err
}

func (r *krsRepository) GetTotalDataByUserId(userid int) (int64, error) {
	var count int64
	err := r.db.Model(&model.KRS{}).Where("user_id = ?", userid).Count(&count).Error
	return count, err
}
