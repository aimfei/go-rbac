package main

import (
	"github.com/gin-gonic/gin"
	"go-rbac/middleware"
	"go-rbac/router"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.LoggerToFile())
	router.InitRolePlatform(engine)
	engine.Run()
}
