package router

import (
	"github.com/gin-gonic/gin"
	apiplatformv1 "go-rbac/api/v1/platform"
)

func InitRolePlatform(route *gin.Engine) {
	platformRouter := route.Group("/platform/v1")
	{
		platformRouter.POST("/save", apiplatformv1.Save) //保存
		platformRouter.GET("/list", apiplatformv1.List)  //查询列表
	}
}
