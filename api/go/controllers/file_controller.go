package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadFile handles file uploads
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload request"})
		return
	}

	err = c.SaveUploadedFile(file, filepath.Join("uploads", file.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

// DeleteFile handles file deletions
func DeleteFile(c *gin.Context) {
	filename := c.Param("filename")
	err := os.Remove(filepath.Join("uploads", filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

// ListFiles lists all files in the uploads directory
func ListFiles(c *gin.Context) {
	files, err := ioutil.ReadDir("uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list files"})
		return
	}

	var fileList []string
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}

	c.JSON(http.StatusOK, gin.H{"files": fileList})
}

// GetFileMetadataHandler retrieves metadata for a specific file
func GetFileMetadataHandler(c *gin.Context) {
	filename := c.Param("filename")
	fileInfo, err := os.Stat(filepath.Join("uploads", filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file metadata"})
		return
	}

	metadata := gin.H{
		"name":    fileInfo.Name(),
		"size":    fileInfo.Size(),
		"modTime": fileInfo.ModTime(),
	}

	c.JSON(http.StatusOK, gin.H{"metadata": metadata})
}
