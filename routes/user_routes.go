package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/api"
	"github.com/revandpratama/go-elearning-api/config"
	"github.com/revandpratama/go-elearning-api/middleware"
	"github.com/revandpratama/go-elearning-api/repository"
	"github.com/revandpratama/go-elearning-api/service"
)

func UserRoutes(r *gin.RouterGroup) {
	userrepo := repository.NewUserRepository(config.DB)
	krsrepo := repository.NewKRSRepository(config.DB)
	scorerepo := repository.NewScoreRepository(config.DB)

	service := service.NewUserService(userrepo, krsrepo, scorerepo)

	handler := api.NewUserAPI(service)

	// r.GET("/test", handler.Login)

	user := r.Group("/users")

	user.Use(middleware.Auth())

	user.GET("/:userid/krs", handler.GetKRS)
	user.GET("/:userid/score", handler.GetScore)

}
