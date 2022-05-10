package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/minhd-vu/react-mon-go/api"
	"github.com/minhd-vu/react-mon-go/util"
)

func main() {
	// Gin router
	router := gin.Default()

	// Serve the front end static files
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	// Router group for the API
	route := router.Group("/api")
	{
		route.GET("/users", api.GetUsers)
		route.POST("/users", api.CreateUser)
	}

	config := util.GetConfiguration()
	router.Run(config.Port)
}
