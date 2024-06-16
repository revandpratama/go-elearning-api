package service

import (
	"math"

	"github.com/revandpratama/go-elearning-api/dto"
	"github.com/revandpratama/go-elearning-api/repository"
)

type userService struct {
	userrepo  repository.UserRepository
	krsrepo   repository.KRSRepository
	scorerepo repository.ScoreRepository
}

type UserService interface {
	GetKRSByUserID(userid int, pagination *dto.Paginate) (*[]dto.KRSResponse, error)
	GetScoreByUserID(userid int, pagination *dto.Paginate) (*[]dto.ScoreResponse, error)
}

func NewUserService(userrepo repository.UserRepository, krsrepo repository.KRSRepository, scorerepo repository.ScoreRepository) *userService {
	return &userService{
		userrepo:  userrepo,
		krsrepo:   krsrepo,
		scorerepo: scorerepo,
	}
}

func (s *userService) GetKRSByUserID(userid int, pagination *dto.Paginate) (*[]dto.KRSResponse, error) {
	krs, err := s.krsrepo.GetByUserID(userid, pagination)
	if err != nil {
		return nil, err
	}


	totalData, err := s.krsrepo.GetTotalDataByUserId(userid)
	if err != nil {
		return nil, err
	}

	pagination.TotalData = int(totalData)
	totalPages := math.Ceil(float64(pagination.TotalData) / float64(pagination.DataPerPage))
	pagination.TotalPages = int(totalPages)
	var response []dto.KRSResponse

	for _, v := range *krs {
		r := dto.KRSResponse{
			ID:        v.ID,
			UserID:    v.UserID,
			SubjectID: v.SubjectID,
			Status:    v.Status,
		}
		response = append(response, r)
	}
	return &response, err
}

func (s *userService) GetScoreByUserID(userid int, pagination *dto.Paginate) (*[]dto.ScoreResponse, error) {
	score, err := s.scorerepo.GetByUserID(userid, pagination)
	if err != nil {
		return nil, err
	}
	totalData, err := s.scorerepo.GetTotalDataByUserId(userid)
	if err != nil {
		return nil, err
	}

	pagination.TotalData = int(totalData)
	totalPages := math.Ceil(float64(pagination.TotalData) / float64(pagination.DataPerPage))
	pagination.TotalPages = int(totalPages)

	var response []dto.ScoreResponse

	for _, v := range *score {
		r := dto.ScoreResponse{
			ID:     v.ID,
			UserID: v.UserID,
			KrsID:  v.KrsID,
			Score:  v.Score,
		}

		response = append(response, r)
	}

	return &response, err
}
