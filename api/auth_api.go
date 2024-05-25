package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/dto"
	"github.com/revandpratama/go-elearning-api/errorhandler"
	"github.com/revandpratama/go-elearning-api/helper"
	"github.com/revandpratama/go-elearning-api/service"
)

type authAPI struct {
	service service.AuthService
}

func NewAuthAPI(s service.AuthService) *authAPI {
	return &authAPI{service: s}
}

func (a *authAPI) Login(g *gin.Context) {
	var LoginRequest dto.LoginRequest

	if err := g.ShouldBindJSON(&LoginRequest); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	tokenString, err := a.service.Login(&LoginRequest); 
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}
	
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "login success",
		Data:       tokenString,
	})

	g.JSON(http.StatusOK, res)
}
