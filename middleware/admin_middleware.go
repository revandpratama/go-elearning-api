package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/errorhandler"
	"github.com/revandpratama/go-elearning-api/helper"
)

func IsAdmin() gin.HandlerFunc {
	return func(g *gin.Context) {
		tokenString, err := g.Cookie("auth_token")

		if err != nil || tokenString == "" {
			errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: "Unauthorized"})
			g.Abort()
			return
		}

		_, role, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: err.Error()})
			g.Abort()
			return
		}

		if *role != "admin" {
			errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: "Unauthorized | only admin"})
			g.Abort()
			return
		}

		g.Next()
	}
}
