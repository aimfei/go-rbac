package api_platform_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-rbac/common"
)

func Save(ctx *gin.Context) {
	logrus.Error("4353")
}
func List(ctx *gin.Context) {
	logrus.Info("12312")
	ctx.JSON(200, common.DataSuccess("123"))
}
