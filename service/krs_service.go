package service

import (
	"github.com/revandpratama/go-elearning-api/dto"
	"github.com/revandpratama/go-elearning-api/model"
	"github.com/revandpratama/go-elearning-api/repository"
)

type krsService struct {
	repository repository.KRSRepository
}

type KRSService interface {
	Create(req *dto.KRSRequest) error
	Update(id int, req *dto.KRSRequest) error
	Delete(id int) error
	GetAll() (*[]model.KRS, error)
	GetById(id int) (*model.KRS, error)
}

func NewKRSService(r repository.KRSRepository) *krsService {
	return &krsService{
		repository: r,
	}
}

func (s *krsService) Create(req *dto.KRSRequest) error {

	krs := model.KRS{
		UserID:    req.UserID,
		SubjectID: req.SubjectID,
		Status:    req.Status,
	}

	if err := s.repository.Create(&krs); err != nil {
		return err
	}

	return nil
}

func (s *krsService) Update(id int, req *dto.KRSRequest) error {
	krs := model.KRS{
		UserID:    req.UserID,
		SubjectID: req.SubjectID,
		Status:    req.Status,
	}

	if err := s.repository.Update(id, &krs); err != nil {
		return err
	}

	return nil
}
func (s *krsService) Delete(id int) error {
	err := s.repository.Delete(id)

	return err
}
func (s *krsService) GetAll() (*[]model.KRS, error) {
	krs, err := s.repository.GetAll()

	return krs, err
}

func (s *krsService) GetById(id int) (*model.KRS, error) {
	return s.repository.GetById(id)
}
