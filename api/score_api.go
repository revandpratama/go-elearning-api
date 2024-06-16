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

type scoreAPI struct {
	service service.ScoreService
}

func NewScoreAPI(service service.ScoreService) *scoreAPI {
	return &scoreAPI{
		service: service,
	}
}

func (a *scoreAPI) Create(g *gin.Context) {
	var req dto.ScoreRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := a.service.Create(&req); err != nil {
		errorhandler.HandleError(g, &errorhandler.InternalServerError{Message: "failed creating score"})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "score created successfully",
	})

	g.JSON(http.StatusCreated, res)
}

func (a *scoreAPI) Update(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	var req dto.ScoreRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := a.service.Update(id, &req); err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "score updated successfully",
	})

	g.JSON(http.StatusOK, res)

}

func (a *scoreAPI) Delete(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := a.service.Delete(id); err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "krs deleted successfully",
	})

	g.JSON(http.StatusOK, res)
}

func (a *scoreAPI) GetAll(g *gin.Context) {
	page := g.Query("page")
	if page == "" {
		page = "1"
	}
	currentPage, err := strconv.Atoi(page)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	perPageString := g.Query("perPage")
	if perPageString == "" {
		perPageString= "10"
	}
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}
	var pagination = dto.Paginate{
		CurrentPage: currentPage,
		DataPerPage: perPage,
	}

	score, err := a.service.GetAll(&pagination)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	data := map[string]any{
		"score":score,
		"pagination": pagination,
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "score retrieved successfully",
		Data:       data,
	})

	g.JSON(http.StatusOK, res)
}

func (a *scoreAPI) GetById(g *gin.Context) {

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}
	score, err := a.service.GetById(id)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "score retrieved successfully",
		Data:       score,
	})

	g.JSON(http.StatusOK, res)
}
