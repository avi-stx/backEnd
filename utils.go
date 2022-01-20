package main

import (
	"io/ioutil"
	"log"
	"os"

	uuid "github.com/nu7hatch/gouuid"
)

func getRelativePath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func getFilesInDir() []fileInfo {
	relativePath := getRelativePath()
	dirPath := relativePath + "/pseudoDB/"
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	filesInDir := []fileInfo{}

	uuid, _ := uuid.NewV4()

	for _, file := range files {
		fileStat, err := os.Stat(dirPath + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		filesInDir = append(filesInDir, fileInfo{uuid.String(), fileStat.Name(), fileStat.Name(), fileStat.Size(), fileStat.ModTime()})
	}

	return filesInDir
}

func readFiles() []fileInfo {

	filesList := getFilesInDir()
	return filesList

}
