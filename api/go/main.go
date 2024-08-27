package main

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type Metadata struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	CreationTime time.Time `json:"creation_time"`
	ModTime      time.Time `json:"mod_time"`
}

type MetadataResponse struct {
	Files []Metadata `json:"files"`
}

func extractMetadata(directory string) ([]Metadata, error) {
	var metadataList []Metadata

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			metadata := Metadata{
				Name:         info.Name(),
				Size:         info.Size(),
				CreationTime: info.ModTime(), // Assuming ModTime as creation time as Go doesn't expose actual creation time on all platforms
				ModTime:      info.ModTime(),
			}
			metadataList.append(metadata)
		}
		return nil
	})

	return metadataList, err
}

func main() {
	router := gin.Default()

	// Metadata Extractor route
	router.GET("/metadata", func(c *gin.Context) {
		directory := c.Query("directory")
		if directory == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Directory path is required"})
			return
		}
		metadataList, err := extractMetadata(directory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract metadata"})
			return
		}
		c.JSON(http.StatusOK, MetadataResponse{Files: metadataList})
	})

	router.Run(":8080") // Start the server on port 8080
}
