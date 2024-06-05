package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/dto"
	"github.com/revandpratama/go-elearning-api/errorhandler"
	"github.com/revandpratama/go-elearning-api/helper"
	"github.com/revandpratama/go-elearning-api/service"
)

type userAPI struct {
	service service.UserService
}

func NewUserAPI(s service.UserService) *userAPI {
	return &userAPI{service: s}
}

func (a *userAPI) GetKRS(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("userid"))
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: "invalid parameter"})
		return
	}

	if err := helper.UserAuthorized(g, id); err != nil {
		errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: "unauthorized request"})
		return
	}

	krs, err := a.service.GetKRSByUserID(id)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "krs by user",
		Data:       krs,
	})

	g.JSON(http.StatusOK, res)
}
func (a *userAPI) GetScore(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("userid"))
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: "invalid parameter"})
		return
	}

	if err := helper.UserAuthorized(g, id); err != nil {
		errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: "unauthorized request"})
		return
	}

	score, err := a.service.GetScoreByUserID(id)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success retrieve score by user",
		Data:       score,
	})

	g.JSON(http.StatusOK, res)
}
