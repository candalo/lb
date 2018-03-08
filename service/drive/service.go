package service

import (
	"fmt"
	"os"
	"sync"

	"github.com/candalo/lb/config/drive"
	"github.com/candalo/lb/storage"
	"github.com/peterbourgon/diskv"
	"google.golang.org/api/drive/v3"
)

var mux sync.Mutex

type upload struct {
	driveService *drive.Service
	storage      *diskv.Diskv
	fileName     string
	folderName   string
}

func (u *upload) createFolder() (string, error) {
	mux.Lock()
	defer mux.Unlock()
	folderID, err := u.storage.Read("folder-" + u.folderName)
	if err != nil {
		fmt.Printf("Creating folder %s...\n", u.folderName)
		driveFolder, err := u.driveService.Files.Create(&drive.File{Name: u.folderName, MimeType: "application/vnd.google-apps.folder"}).Do()
		if err != nil {
			return "", err
		}
		u.storage.Write("folder-"+u.folderName, []byte(driveFolder.Id))
		return driveFolder.Id, nil
	}
	return string(folderID), nil
}

func (u *upload) createFile(file *os.File, folderID string) (string, error) {
	fmt.Printf("Uploading file %s...\n", file.Name())
	driveFileID, err := u.storage.Read(u.fileName)
	if err != nil {
		// File has not yet been uploaded
		driveFile, err := u.driveService.Files.Create(&drive.File{Name: u.fileName, Parents: []string{folderID}}).Media(file).Do()
		if err != nil {
			return "", err
		}
		err = u.storage.Write(u.fileName, []byte(driveFile.Id))
		if err != nil {
			return "", err
		}
		return driveFile.Id, nil
	}
	// File has already been uploaded
	driveFile, err := u.driveService.Files.Update(string(driveFileID), &drive.File{Name: u.fileName}).Media(file).Do()
	if err != nil {
		return "", err
	}
	return driveFile.Id, nil
}

// Upload uploads the file to Google Drive
//
// Returns file id if uploaded successfully and an error otherwise
func Upload(filePath string, folderName string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileStat, err := file.Stat()
	if err != nil {
		return "", err
	}
	defer file.Close()

	driveService := config.GetDriveService()

	storage, err := storage.GetStorage()
	if err != nil {
		return "", err
	}

	upload := upload{driveService: driveService, storage: storage, fileName: fileStat.Name(), folderName: folderName}
	folderID, err := upload.createFolder()
	if err != nil {
		return "", err
	}
	fileID, err := upload.createFile(file, folderID)
	if err != nil {
		return "", err
	}
	return fileID, nil
}
