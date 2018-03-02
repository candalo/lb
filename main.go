package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/candalo/lb/service/drive"
)

var wg = &sync.WaitGroup{}

func callUploadService(filePath string) {
	defer wg.Done()

	fmt.Printf("Uploading file %s...\n", filePath)
	fileID, err := service.Upload(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
	fmt.Printf("File %s uploaded with id %s\n", filePath, fileID)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: lb [filename] or lb [folder]")
		return
	}

	filePath := os.Args[1]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File or folder %s does not exist\n", filePath)
		return
	}

	if fileInfo.IsDir() {
		filesInfo, err := ioutil.ReadDir(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on read dir %s\n", filePath)
		}
		for _, fileInfo := range filesInfo {
			wg.Add(1)
			go callUploadService(filepath.Join(filePath, fileInfo.Name()))
		}
	} else {
		wg.Add(1)
		go callUploadService(filePath)
	}
	wg.Wait()
}
