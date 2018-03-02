package storage

import (
	"os/user"
	"path/filepath"

	"github.com/peterbourgon/diskv"
)

const folder = ".lb"

// GetStorage returns the storage instance that will be used to save and retrieve data
func GetStorage() (*diskv.Diskv, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	folderpath := filepath.Join(usr.HomeDir, folder)

	flatTransform := func(s string) []string { return []string{} }
	d := diskv.New(diskv.Options{
		BasePath:     folderpath,
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
	return d, nil
}
