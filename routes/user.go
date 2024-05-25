package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/api"
	"github.com/revandpratama/go-elearning-api/config"
	"github.com/revandpratama/go-elearning-api/repository"
	"github.com/revandpratama/go-elearning-api/service"
)

func UserRouter(r *gin.RouterGroup) {
	repo := repository.NewUserRepository(config.DB)
	service := service.NewUserService(repo)
	handler := api.NewUserAPI(service)

	r.GET("/test", handler.Login)
}
