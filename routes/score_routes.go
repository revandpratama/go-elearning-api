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

	score := r.Group("/score")

	score.Use(middleware.Auth())

	score.GET("/", middleware.IsAdmin(), handler.GetAll)
	score.GET("/:id", handler.GetById)
	score.POST("/", middleware.IsAdmin(), handler.Create)
	score.PUT(":id", middleware.IsAdmin(), handler.Update)
	score.DELETE(":id", middleware.IsAdmin(), handler.Delete)
}
