package repository

import (
	"github.com/revandpratama/go-elearning-api/model"
	"gorm.io/gorm"
)

type scoreRepository struct {
	db *gorm.DB
}

type ScoreRepository interface {
	Create(krs *model.Score) error
	Update(id int, score *model.Score) error
	Delete(id int) error
	GetAll() (*[]model.Score, error)
	GetById(id int) (*model.Score, error)
	GetByUserID(userid int) (*[]model.Score, error)
}

func NewScoreRepository(db *gorm.DB) *scoreRepository {
	return &scoreRepository{
		db: db,
	}
}

func (r *scoreRepository) Create(score *model.Score) error {
	err := r.db.Create(&score).Error
	return err
}

func (r *scoreRepository) Update(id int, score *model.Score) error {
	err := r.db.Where("id = ?", id).Updates(&score).Error
	return err
}

func (r *scoreRepository) Delete(id int) error {
	err := r.db.Delete(&model.Score{}, id).Error
	return err
}

func (r *scoreRepository) GetAll() (*[]model.Score, error) {
	var score []model.Score
	err := r.db.Find(&score).Error
	return &score, err
}

func (r *scoreRepository) GetById(id int) (*model.Score, error) {
	var score model.Score
	err := r.db.Where("id = ?", id).First(&score).Error
	return &score, err
}

func (r *scoreRepository) GetByUserID(userid int) (*[]model.Score, error) {
	var score []model.Score
	err := r.db.Find(&score, "user_id", userid).Error
	return &score, err
}
