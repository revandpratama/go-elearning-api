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

type krsAPI struct {
	service service.KRSService
}

func NewKRSAPI(s service.KRSService) *krsAPI {
	return &krsAPI{
		service: s,
	}
}

func (a *krsAPI) Create(g *gin.Context) {
	var req dto.KRSRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := a.service.Create(&req); err != nil {
		errorhandler.HandleError(g, &errorhandler.InternalServerError{Message: "failed creating krs"})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "krs created successfully",
	})

	g.JSON(http.StatusCreated, res)
}
func (a *krsAPI) Update(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	var req dto.KRSRequest

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
		Message:    "krs updated successfully",
	})

	g.JSON(http.StatusOK, res)

}
func (a *krsAPI) Delete(g *gin.Context) {
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
func (a *krsAPI) GetAll(g *gin.Context) {
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

	krs, err := a.service.GetAll(&pagination)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "krs retrieved successfully",
		Data:       krs,
	})

	g.JSON(http.StatusOK, res)
}
func (a *krsAPI) GetById(g *gin.Context) {

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}
	krs, err := a.service.GetById(id)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "krs retrieved successfully",
		Data:       krs,
	})

	g.JSON(http.StatusOK, res)
}
