package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/dto"
	"github.com/revandpratama/go-elearning-api/errorhandler"
	"github.com/revandpratama/go-elearning-api/helper"
	"github.com/revandpratama/go-elearning-api/model"
	"github.com/revandpratama/go-elearning-api/service"
)

type userAPI struct {
	service service.UserService
}

func NewUserAPI(s service.UserService) *userAPI {
	return &userAPI{service: s}
}

func (a *userAPI) Login(g *gin.Context) {
	var LoginRequest dto.LoginRequest
	var userDB *model.User
	var err error
	if err := g.ShouldBindJSON(&LoginRequest); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}	

	if strings.Contains(LoginRequest.UsernameOrEmail, "@") {
		userDB, err = a.service.GetUserByEmail(LoginRequest.UsernameOrEmail)
		if err != nil {
			errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
			return
		}
	}else {
		userDB, err = a.service.GetUserByUsername(LoginRequest.UsernameOrEmail)
		if err != nil {
			errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
			return
		}
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login Success",
		Data: token,
	})

	g.JSON(http.StatusOK, res)
}
