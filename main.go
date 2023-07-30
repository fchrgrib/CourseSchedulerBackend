package main

import (
	"docker_training.com/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"*"}

	router.Use(cors.New(config))
	routers.CMDRouters(router)
	routers.CSRouters(router)

	if err := router.Run(); err != nil {
		return
	}
}
