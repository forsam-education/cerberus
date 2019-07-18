package utils

import (
	"log"
	"os"
	"path/filepath"
)

// DirExists checks if a directory exists
func DirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	var _, err = os.Stat(path)

	return !os.IsNotExist(err)
}

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}

// CreateFile create a file and the required directories if necessary
func CreateFile(path string) {
	ensureDir(path)
	if !FileExists(path) {
		var file, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}
