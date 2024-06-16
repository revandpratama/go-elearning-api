package service

import (
	"math"

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
	GetAll(pagination *dto.Paginate) (*[]dto.KRSResponse, error)
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
func (s *krsService) GetAll(pagination *dto.Paginate) (*[]dto.KRSResponse, error) {
	krs, err := s.repository.GetAll(pagination)
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

	var response []dto.KRSResponse

	for _, v := range *krs {
		r := dto.KRSResponse{
			ID:         v.ID,
			UserID:     v.UserID,
			SubjectID:  v.SubjectID,
			Status:     v.Status,
			Pagination: *pagination,
		}

		response = append(response, r)
	}

	return &response, err
}

func (s *krsService) GetById(id int) (*model.KRS, error) {
	return s.repository.GetById(id)
}
