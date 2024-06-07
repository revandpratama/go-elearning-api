package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/api"
	"github.com/revandpratama/go-elearning-api/config"
	"github.com/revandpratama/go-elearning-api/middleware"
	"github.com/revandpratama/go-elearning-api/repository"
	"github.com/revandpratama/go-elearning-api/service"
)

func ScoreRoutes(r *gin.RouterGroup) {
	repo := repository.NewScoreRepository(config.DB)
	service := service.NewScoreService(repo)
	handler := api.NewScoreAPI(service)

	//user path
	score := r.Group("/scores")
	score.Use(middleware.Auth())

	score.GET("/:id", handler.GetById)


	// Admin only path
	scoreAdmin := score.Group("/")
	scoreAdmin.Use(middleware.IsAdmin())

	scoreAdmin.GET("/", handler.GetAll)
	scoreAdmin.POST("/", handler.Create)
	scoreAdmin.PUT(":id", handler.Update)
	scoreAdmin.DELETE(":id", handler.Delete)
}
