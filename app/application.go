package app

import (
	"github.com/gin-gonic/gin"
	"github.com/khalil-farashiani/microservice_users-api/logger"
)

var (
	router = gin.Default()
)

// StartApplication start server
func StartApplication() {
	mapUrls()
	logger.Info("about to start the application ...")
	router.Run(":8080")
}
