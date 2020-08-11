package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-rbac/config"
	"os"
	"path"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.Log_FILE_PATH
	logFileName := config.LOG_FILE_NAME
	fileName := path.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return func(c *gin.Context) {
		startTime := time.Now()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
