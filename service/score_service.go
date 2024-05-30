package service

import (
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
	GetAll() (*[]model.Score, error)
	GetById(id int) (*model.Score, error)
}

func NewScoreService(r repository.ScoreRepository) *scoreService {
	return &scoreService{
		repository: r,
	}
}

func (s *scoreService) Create(req *dto.ScoreRequest) error {

	score := model.Score{
		UserID: req.UserID,
		KrsID: req.KrsID,
		Score: req.Score,
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
func (s *scoreService) GetAll() (*[]model.Score, error) {
	krs, err := s.repository.GetAll()

	return krs, err
}

func (s *scoreService) GetById(id int) (*model.Score, error) {
	return s.repository.GetById(id)
}
