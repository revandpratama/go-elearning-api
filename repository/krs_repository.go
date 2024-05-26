package repository

import "gorm.io/gorm"

type krsRepository struct {
	db *gorm.DB
}

type KRSRepository interface {
}

func NewKRSRepository(db *gorm.DB) *krsRepository {
	return &krsRepository{
		db: db,
	}
}


