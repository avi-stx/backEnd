package main

import (
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

const DIR_NAME = "/pseudoDB/"

func saveFileHandler(c *gin.Context) bool {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
		return false
	}

	filename := header.Filename

	fullFileName := getFullPathToFile(filename)
	out, err := os.Create(fullFileName)
	if err != nil {
		log.Fatal(err)
		return false

	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func isFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func getFullPathToFile(fileName string) string {
	relativePath := getRelativePath()
	dirPath := relativePath + DIR_NAME
	return dirPath + fileName
}

func removeFile(fileName string) bool {

	fullPath := getFullPathToFile(fileName)
	fileExist := isFileExist(fullPath)

	if !fileExist {
		return false
	}

	e := os.Remove(fullPath)

	if e != nil {
		log.Fatal(e)
		return false
	}

	return true
}

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
	for _, file := range files {
		fileStat, err := os.Stat(dirPath + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		uuid, _ := uuid.NewV4()
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
