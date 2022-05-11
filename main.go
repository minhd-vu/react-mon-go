package main

import (
	"github.com/minhd-vu/react-mon-go/api"
	"github.com/minhd-vu/react-mon-go/util"
)

func main() {
	// Gin router
	router := api.SetupRouter()
	router.Run(util.GetConfiguration().Port)
}
