package service

import (
	"os"

	"github.com/candalo/lb/config/drive"
	"google.golang.org/api/drive/v3"
)

// Upload uploads the file to Google Drive
//
// Returns file id if uploaded successfully and an error otherwise
func Upload(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	driveService := config.GetDriveService()
	driveFile, err := driveService.Files.Create(&drive.File{Name: file.Name()}).Media(file).Do()
	if err != nil {
		return "", err
	}
	return driveFile.Id, nil
}
