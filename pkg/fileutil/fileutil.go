package fileutil

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func GetFilesOnDir(directory string) []string {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}
