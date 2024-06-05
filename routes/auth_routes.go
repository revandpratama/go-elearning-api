package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/api"
	"github.com/revandpratama/go-elearning-api/config"
	"github.com/revandpratama/go-elearning-api/middleware"
	"github.com/revandpratama/go-elearning-api/repository"
	"github.com/revandpratama/go-elearning-api/service"
)

func AuthRoutes(r *gin.RouterGroup) {
	repo := repository.NewUserRepository(config.DB)
	service := service.NewAuthService(repo)
	handler := api.NewAuthAPI(service)

	guestRoutes := r.Group("/")
	guestRoutes.Use(middleware.Guest())
	guestRoutes.POST("/login", handler.Login)
	guestRoutes.POST("/register", handler.Register)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.Auth())
	authRoutes.POST("/logout", handler.Logout)
}
