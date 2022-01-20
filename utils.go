package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	uuid "github.com/nu7hatch/gouuid"
)

const DIR_NAME = "/pseudoDB/"

func getRelativePath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func readDirContent(path string) []fs.FileInfo {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func generateFilesList(files []fs.FileInfo, dirPath string) []fileInfo {
	filesInDir := []fileInfo{}

	uuid, _ := uuid.NewV4()

	for _, file := range files {
		fileStat, err := os.Stat(dirPath + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		filesInDir = append(filesInDir,
			fileInfo{uuid.String(),
				fileStat.Name(),
				fileStat.Name(),
				fileStat.Size(),
				fileStat.ModTime()})
	}

	return filesInDir
}

func getFilesInDir() []fileInfo {

	relativePath := getRelativePath()
	dirPath := relativePath + DIR_NAME
	files := readDirContent(dirPath)
	filesInDir := generateFilesList(files, dirPath)

	return filesInDir
}

func readFiles() []fileInfo {

	filesList := getFilesInDir()
	return filesList

}
