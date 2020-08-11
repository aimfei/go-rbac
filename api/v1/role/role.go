package role

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-rbac/common"
)

func List(ctx *gin.Context) {
	logrus.Info("132123")
	ctx.JSON(200, common.DataSuccess("123"))
}
