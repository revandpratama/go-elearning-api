package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/api"
	"github.com/revandpratama/go-elearning-api/config"
	"github.com/revandpratama/go-elearning-api/repository"
	"github.com/revandpratama/go-elearning-api/service"
)

func AuthRoutes(r *gin.RouterGroup) {
	repo := repository.NewUserRepository(config.DB)
	service := service.NewAuthService(repo)
	handler := api.NewAuthAPI(service)


	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
}
