package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/errorhandler"
)

func Guest() gin.HandlerFunc {
	return func(g *gin.Context) {
		tokenString, err := g.Cookie("auth_token")

		if err != nil || tokenString == "" {
			g.Next()
			return
		}

		errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: "already logged in!"})
		g.Abort()
	}
}
