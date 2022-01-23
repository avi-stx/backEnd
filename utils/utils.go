package utils

import (
	"Backend_side/config"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	uuid "github.com/nu7hatch/gouuid"
)

var DIR_NAME = config.AppConfig.TargetFolder // "pseudoDB"

func SaveFileHandler(file *io.Reader, filename string) bool {

	fullFileName := GetFullPathToFile(filename)
	out, err := os.Create(fullFileName)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer out.Close()
	_, err = io.Copy(out, *file)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func IsFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func GetFullPathToFile(fileName string) string {
	// relativePath := getRelativePath()
	dirPath := config.AppConfig.TargetFolder
	return dirPath + fileName
}

func RemoveFile(fileName string) bool {
	fullPath := GetFullPathToFile(fileName)

	if fileName == "" {
		return false
	}

	fileExist := IsFileExist(fullPath)

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

func getFilesInDir(dirName string) []fileInfo {

	relativePath := getRelativePath()
	dirPath := relativePath + "/" + dirName
	files := readDirContent(dirPath)
	filesInDir := generateFilesList(files, dirPath)

	return filesInDir
}

func ReadFiles() []fileInfo {
	filesList := getFilesInDir(config.AppConfig.TargetFolder)
	return filesList

}
