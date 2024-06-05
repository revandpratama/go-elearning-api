package helper

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserAuthorized(g *gin.Context, id int) error {
	authID := g.MustGet("userID").(*int)
	authRole := g.MustGet("userRole").(*string)
	fmt.Println(id)
	fmt.Println(authID)
	fmt.Println(authRole)
	if id == *authID || *authRole == "admin" {
		return nil
	}
	return errors.New("unauthorized request")
}
