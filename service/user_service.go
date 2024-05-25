package service

import (
	"github.com/revandpratama/go-elearning-api/model"
	"github.com/revandpratama/go-elearning-api/repository"
)

type userService struct {
	repo repository.UserRepository
}

type UserService interface {
	GetUserByEmail(email string) (*model.User, error)
}


func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		repo: r,
	}
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.GetUserByEmail(email)
}

