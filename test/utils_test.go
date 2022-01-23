package test

import (
	"Backend_side/config"
	"Backend_side/utils"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFile(t *testing.T) {
	err := ioutil.WriteFile(config.AppConfig.TargetFolder+"existFile.txt", []byte("Hello"), 0755)

	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	tables := []struct {
		x           string
		expected    bool
		description string
	}{
		{"existFile.txt", true, "deleting an existing file"},
		{"nonExist.txt", false, "deleting unexisting file"},
		{"", false, "deleting empty string file name"},
	}

	for _, table := range tables {
		result := utils.RemoveFile(table.x)
		if result != table.expected {
			t.Errorf("test description: %s file named: %s returned %v , expected %v.", table.description, table.x, result, table.expected)
		}
	}
}

func TestReadFilesAmount(t *testing.T) {
	filesList := utils.ReadFiles()
	if len(filesList) != 2 {
		t.Errorf("not enought files were read")
	}
}

func TestReadFilesNames(t *testing.T) {
	filesList := utils.ReadFiles()
	name1, name2 := filesList[0].Name, filesList[1].Name
	if name1 != "demo_file.txt" || name2 != "demo_file2.txt" {
		t.Errorf("file were not read correctly")

	}
}

func TestIsExist(t *testing.T) {

	fileNotExist := "FakeFile.txt"

	if utils.IsFileExist(fileNotExist) == true {
		t.Errorf("file doesn't exist, but true returned, meaning file found")
	}

}

func TestGetFullPathToFile(t *testing.T) {

	actual := utils.GetFullPathToFile("demo_file.txt")
	expected := "../tests_assets/demo_file.txt"

	if expected != actual {
		t.Errorf("path %s not equal to %s", actual, expected)
	}

}

func TestGetFiles(t *testing.T) {
	router, w := setup()
	req := httptest.NewRequest("GET", "/files", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDownload(t *testing.T) {
	router, w := setup()
	req := httptest.NewRequest("GET", "/files/demo_file.txt", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpload(t *testing.T) {
	router, w := setup()
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	go func() {
		defer writer.Close()
		part, err := writer.CreateFormFile("file", "test_demo_file.txt")
		assert.NoError(t, err)
		file, _ := os.Open("../tests_assets/demo_file.txt")
		fmt.Println("file is: ", file)
		defer file.Close()
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error(err)
		}
	}()

	request := httptest.NewRequest("POST", "/files", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	router.ServeHTTP(w, request)
	assert.Equal(t, 200, w.Code)
	os.Remove("../tests_assets/test_demo_file.txt")
}
