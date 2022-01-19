package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	for _, file := range files {
		fileStat, err := os.Stat(dirPath + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		filesInDir = append(filesInDir, fileInfo{1, fileStat.Name(), fileStat.Name(), fileStat.Size(), fileStat.ModTime()})
		fmt.Println(file.Name())
	}

	return filesInDir
}

func readFiles() {

	files := getFilesInDir()
	fmt.Println(files)

}
