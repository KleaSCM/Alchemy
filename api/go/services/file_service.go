package services

import (
	"os"
	"time"
)

type FileMetadata struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"mod_time"`
}

// GetFileMetadata retrieves metadata for a file
func GetFileMetadata(filePath string) (FileMetadata, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return FileMetadata{}, err
	}

	metadata := FileMetadata{
		Name:    fileInfo.Name(),
		Size:    fileInfo.Size(),
		ModTime: fileInfo.ModTime(),
	}

	return metadata, nil
}
