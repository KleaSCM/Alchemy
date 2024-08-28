package controllers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// FileInfo represents the information of a file
type FileInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// ListFiles lists all files in the specified directory
func ListFiles(c *gin.Context) {
	directory := `C:\Users\Kliea\Documents\TestTESTtestALCHEMY`
	var files []FileInfo

	log.Printf("Starting to walk the directory: %s", directory)

	// Check if the directory exists
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Printf("Directory does not exist: %s", directory)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Directory does not exist"})
		return
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v", path, err)
			return err
		}

		if !info.IsDir() {
			log.Printf("Found file: %s", info.Name())
			files = append(files, FileInfo{Name: info.Name(), Size: info.Size()})
		} else {
			log.Printf("Skipping directory: %s", info.Name())
		}

		return nil
	})

	if err != nil {
		log.Printf("Error walking the directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list files"})
		return
	}

	if len(files) == 0 {
		log.Println("No files found in the directory")
	} else {
		log.Printf("Found %d files", len(files))
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
