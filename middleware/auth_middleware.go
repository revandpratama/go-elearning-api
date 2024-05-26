package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/errorhandler"
	"github.com/revandpratama/go-elearning-api/helper"
)

func Auth() gin.HandlerFunc {
	return func(g *gin.Context) {
		tokenString, err := g.Cookie("auth_token")

		if err != nil || tokenString == "" {
			errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: "Unauthorized"})
			g.Abort()
			return
		}

		id, role, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: err.Error()})
			g.Abort()
			return
		}

		g.Set("userID", id)
		g.Set("userRole", role)
		g.Next()
	}
}