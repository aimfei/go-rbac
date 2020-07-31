package main

import (
	"github.com/gin-gonic/gin"
	"go-rbac/controller/v1"
	"go-rbac/middleware"
)

func main() {
	router := gin.Default()
	router.Use(middleware.LoggerToFile())

	platformController := router.Group("/platform/v1")
	{
		platformController.POST("/save", v1.Save)
	}
	router.Run()
}
