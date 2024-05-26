package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/errorhandler"
)

func IsAdmin() gin.HandlerFunc {
	return func(g *gin.Context) {
		role := g.GetString("userRole")

		if role != "admin" {
			errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: "Unauthorized | only admin"})
			g.Abort()
			return
		}

		g.Next()
	}
}
