package main

import (
	"app-aposrgb/src/config"
	"app-aposrgb/src/database"
	"app-aposrgb/src/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	database.Init()
	r := gin.Default()
	routes.InitRoutes(r)
	err := r.Run(config.GetPort())
	if err != nil {
		fmt.Println("Not valid port: " + config.GetPort(), err)
	}
}