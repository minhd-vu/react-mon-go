package api

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Serve the front end static files
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	// Router group for the API
	route := router.Group("/api")
	{
		route.GET("/users", GetUsers)
		route.POST("/users", CreateUser)
	}

	return router
}
