package main

import (
	"github.com/Jay-SCM/alchemy/api/go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect routes to their handler functions
	router.GET("/files", controllers.ListFiles)
	router.POST("/upload", controllers.UploadFile)
	router.DELETE("/delete/:filename", controllers.DeleteFile)
	router.GET("/metadata/:filename", controllers.GetFileMetadataHandler)

	router.Run(":8080")
}
