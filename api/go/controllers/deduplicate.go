package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type DeduplicationResponse struct {
	Duplicates []string `json:"duplicates"`
}

// Calculate the SHA-256 checksum of a file
func calculateChecksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Find duplicate files in the directory by comparing checksums
func findDuplicates(directory string) ([]string, error) {
	hashMap := make(map[string]string)
	var duplicates []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			checksum, err := calculateChecksum(path)
			if err != nil {
				return err
			}
			if originalPath, exists := hashMap[checksum]; exists {
				duplicates = append(duplicates, path)
				duplicates = append(duplicates, originalPath)
			} else {
				hashMap[checksum] = path
			}
		}
		return nil
	})

	return duplicates, err
}

// Handle the deduplication request
func Deduplicate(c *gin.Context) {
	directory := c.PostForm("directory")
	if directory == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Directory path is required"})
		return
	}
	duplicates, err := findDuplicates(directory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find duplicates"})
		return
	}
	c.JSON(http.StatusOK, DeduplicationResponse{Duplicates: duplicates})
}
