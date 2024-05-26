package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/api"
	"github.com/revandpratama/go-elearning-api/config"
	"github.com/revandpratama/go-elearning-api/middleware"
	"github.com/revandpratama/go-elearning-api/repository"
	"github.com/revandpratama/go-elearning-api/service"
)

func KRSRoutes(r *gin.RouterGroup) {
	repo := repository.NewKRSRepository(config.DB)
	service := service.NewKRSService(repo)
	handler := api.NewKRSAPI(service)

	// r.Use(middleware.Auth())
	krs := r.Group("/krs")

	krs.Use(middleware.Auth())

	krs.GET("/", middleware.IsAdmin(), handler.GetAll)
	krs.GET("/:id", handler.GetById)
	krs.POST("/", handler.Create)
	krs.PUT(":id", handler.Update)
	krs.DELETE(":id", handler.Delete)
}
