package api

import (
	"github.com/revandpratama/go-elearning-api/service"
)

type userAPI struct {
	service service.UserService
}

func NewUserAPI(s service.UserService) *userAPI {
	return &userAPI{service: s}
}
