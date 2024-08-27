package routes

import (
	"github.com/Jay-SCM/alchemy/api/go/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/files", controllers.ListFiles)
	router.POST("/files/upload", controllers.UploadFile)
	router.DELETE("/files/:name", controllers.DeleteFile)
}
