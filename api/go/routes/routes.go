package routes

import (
	"github.com/Jay-SCM/alchemy-backend/api/go/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// File routes
	router.GET("/files", controllers.ListFiles)
	router.POST("/files/upload", controllers.UploadFile)
	router.DELETE("/files/:name", controllers.DeleteFile)

	// Deduplication route
	router.POST("/deduplicate", controllers.Deduplicate)
}
