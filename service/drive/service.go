package service

import (
	"os"

	"github.com/candalo/lb/config/drive"
	"github.com/candalo/lb/storage"
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
	fileStat, err := file.Stat()
	if err != nil {
		return "", nil
	}
	defer file.Close()

	driveService := config.GetDriveService()

	storage, err := storage.GetStorage()
	if err != nil {
		return "", err
	}
	driveFileID, err := storage.Read(fileStat.Name())
	if err != nil {
		// File has not yet been uploaded
		driveFile, err := driveService.Files.Create(&drive.File{Name: file.Name()}).Media(file).Do()
		if err != nil {
			return "", err
		}
		err = storage.Write(fileStat.Name(), []byte(driveFile.Id))
		if err != nil {
			return "", err
		}
		return driveFile.Id, nil
	}
	// File has already been uploaded
	driveFile, err := driveService.Files.Update(string(driveFileID), &drive.File{Name: file.Name()}).Media(file).Do()
	if err != nil {
		return "", nil
	}
	return driveFile.Id, nil
}
