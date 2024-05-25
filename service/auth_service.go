package service

import (
	"strings"

	"github.com/revandpratama/go-elearning-api/dto"
	"github.com/revandpratama/go-elearning-api/errorhandler"
	"github.com/revandpratama/go-elearning-api/helper"
	"github.com/revandpratama/go-elearning-api/model"
	"github.com/revandpratama/go-elearning-api/repository"
)

type authService struct {
	repository repository.UserRepository
}

type AuthService interface {
	Login(LoginRequest *dto.LoginRequest) (*string, error)
	Register(RegisterRequest *dto.RegisterRequest) error
}

func NewAuthService(r repository.UserRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Login(LoginRequest *dto.LoginRequest) (*string, error) {
	var userDB *model.User
	var err error
	var token *string
	if strings.Contains(LoginRequest.UsernameOrEmail, "@") {
		userDB, err = s.repository.GetUserByEmail(LoginRequest.UsernameOrEmail)
		if err != nil {
			return nil, err
		}
	} else {
		userDB, err = s.repository.GetUserByUsername(LoginRequest.UsernameOrEmail)
		if err != nil {
			return nil, err
		}
	}

	if err := helper.VerifyPassword(userDB.Password, LoginRequest.Password); err != nil {
		return nil, err
	}

	token, err = helper.GenerateToken(userDB)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *authService) Register(RegisterRequest *dto.RegisterRequest) error {

	if RegisterRequest.Password != RegisterRequest.PasswordConfirm {
		return &errorhandler.BadRequestError{Message: "password comfirmation not match"}
	}

	HashedPassword, err := helper.HashPassword(RegisterRequest.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Role:     RegisterRequest.Role,
		Name:     RegisterRequest.Name,
		Username: RegisterRequest.Username,
		Email:    RegisterRequest.Email,
		Phone:    RegisterRequest.Phone,
		Password: HashedPassword,
	}

	if err := s.repository.Create(&user); err != nil {
		return err
	}

	return nil
}
