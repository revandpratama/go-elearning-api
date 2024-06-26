package service

import (
	"math"

	"github.com/revandpratama/go-elearning-api/dto"
	"github.com/revandpratama/go-elearning-api/model"
	"github.com/revandpratama/go-elearning-api/repository"
)

type scoreService struct {
	repository repository.ScoreRepository
}

type ScoreService interface {
	Create(req *dto.ScoreRequest) error
	Update(id int, req *dto.ScoreRequest) error
	Delete(id int) error
	GetAll(pagination *dto.Paginate) (*[]dto.ScoreResponse, error)
	GetById(id int) (*dto.ScoreResponse, error)
}

func NewScoreService(r repository.ScoreRepository) *scoreService {
	return &scoreService{
		repository: r,
	}
}

func (s *scoreService) Create(req *dto.ScoreRequest) error {

	score := model.Score{
		UserID: req.UserID,
		KrsID:  req.KrsID,
		Score:  req.Score,
	}

	if err := s.repository.Create(&score); err != nil {
		return err
	}

	return nil
}

func (s *scoreService) Update(id int, req *dto.ScoreRequest) error {
	score := model.Score{
		KrsID: req.KrsID,
		Score: req.Score,
	}

	if err := s.repository.Update(id, &score); err != nil {
		return err
	}

	return nil
}
func (s *scoreService) Delete(id int) error {
	err := s.repository.Delete(id)

	return err
}
func (s *scoreService) GetAll(pagination *dto.Paginate) (*[]dto.ScoreResponse, error) {
	score, err := s.repository.GetAll(pagination)
	if err != nil {
		return nil, err
	}

	totalData, err := s.repository.GetTotalData()
	if err != nil {
		return nil, err
	}

	pagination.TotalData = int(totalData)
	totalPages := math.Ceil(float64(pagination.TotalData) / float64(pagination.DataPerPage))
	pagination.TotalPages = int(totalPages)

	var response []dto.ScoreResponse
	for _, v := range *score {
		s := dto.ScoreResponse{
			ID:         v.ID,
			UserID:     v.UserID,
			KrsID:      v.KrsID,
			Score:      v.Score,
		}

		response = append(response, s)

	}

	return &response, err
}

func (s *scoreService) GetById(id int) (*dto.ScoreResponse, error) {
	score, err := s.repository.GetById(id)
	response := dto.ScoreResponse{
		ID:         score.ID,
		UserID:     score.UserID,
		KrsID:      score.KrsID,
		Score:      score.Score,
	}

	return &response, err
}
