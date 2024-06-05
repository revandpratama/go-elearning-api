package service

import (
	"github.com/revandpratama/go-elearning-api/model"
	"github.com/revandpratama/go-elearning-api/repository"
)

type userService struct {
	userrepo  repository.UserRepository
	krsrepo   repository.KRSRepository
	scorerepo repository.ScoreRepository
}

type UserService interface {
	GetKRSByUserID(userid int) (*[]model.KRS, error)
	GetScoreByUserID(userid int) (*[]model.Score, error)
}

func NewUserService(userrepo repository.UserRepository, krsrepo repository.KRSRepository, scorerepo repository.ScoreRepository) *userService {
	return &userService{
		userrepo:  userrepo,
		krsrepo:   krsrepo,
		scorerepo: scorerepo,
	}
}

func (s *userService) GetKRSByUserID(userid int) (*[]model.KRS, error) {
	krs, err := s.krsrepo.GetByUserID(userid)
	if err != nil {
		return nil, err
	}
	return krs, err
}

func (s *userService) GetScoreByUserID(userid int) (*[]model.Score, error) {
	return s.scorerepo.GetByUserID(userid)
}
