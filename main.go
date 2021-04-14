package main

import (
	"github.com/minhd-vu/go-project/helpers"
	"github.com/minhd-vu/go-project/routes"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	// Gin router
	router := gin.Default()

	// Serve the front end static files
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	// Router group for the API
	api := router.Group("/api")
	{
		api.GET("/users", routes.GetUsers)
		api.POST("/users", routes.CreateUser)
	}

	config := helpers.GetConfiguration()
	router.Run(config.Port)
}
