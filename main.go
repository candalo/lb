package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/candalo/lb/service/drive"
)

var wg = &sync.WaitGroup{}

func callUploadService(filePath string, folderName string) {
	defer wg.Done()

	fileID, err := service.Upload(filePath, folderName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	fmt.Printf("File %s uploaded with id %s\n", filePath, fileID)
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: lb --filepath<filename or folder> [--folder=<name>]")
		return
	}

	filePath := flag.String("filepath", "", "Use this option to indicate the file which should be uploaded")
	folderName := flag.String("folder", "", "Use this option to indicate where uploaded file will be stored")
	flag.Parse()

	fileInfo, err := os.Stat(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File or folder %s does not exist\n", *filePath)
		return
	}

	if len(*folderName) == 0 {
		*folderName = "lb-default-folder"
	}

	if fileInfo.IsDir() {
		filesInfo, err := ioutil.ReadDir(*filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on read dir %s\n", *filePath)
		}
		for _, fileInfo := range filesInfo {
			wg.Add(1)
			go callUploadService(filepath.Join(*filePath, fileInfo.Name()), *folderName)
		}
	} else {
		wg.Add(1)
		go callUploadService(*filePath, *folderName)
	}
	wg.Wait()
}
