package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-elearning-api/config"
	"github.com/revandpratama/go-elearning-api/routes"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()

	api := r.Group("/api")

	routes.UserRoutes(api)
	routes.AuthRoutes(api)
	routes.KRSRoutes(api)
	routes.ScoreRoutes(api)

	addressString := fmt.Sprintf("localhost:%v", config.ENV.PORT)
	r.Run(addressString)
}
