package controllers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadFile handles file uploads
func UploadFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}
	defer file.Close()

	// Define the upload directory
	directory := `C:\Users\Kliea\Documents\TestTESTtestALCHEMY`

	// Create the file path
	filePath := filepath.Join(directory, "uploaded_file") // Example filename; you might want to generate a unique name

	// Save the uploaded file
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

// DeleteFile handles file deletions
func DeleteFile(c *gin.Context) {
	// Get the file name from the URL parameter
	fileName := c.Param("name")

	// Construct the full file path
	filePath := filepath.Join(`C:\Users\Kliea\Documents\TestTESTtestALCHEMY`, fileName)

	// Attempt to remove the file
	err := os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File deletion failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
