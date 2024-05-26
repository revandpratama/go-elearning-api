package service

import "github.com/revandpratama/go-elearning-api/repository"

type krsService struct {
	repository repository.KRSRepository
}

type KRSService interface {
}

func NewKRSService(r repository.KRSRepository) *krsService {
	return &krsService{
		repository: r,
	}
}