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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const AMOUNT_OF_DEMO_FILES = 5

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

func createDemoFiles() {
	for i := 0; i < AMOUNT_OF_DEMO_FILES; i++ {
		os.Create("../tests_assets/demo_file" + strconv.Itoa(i) + ".txt")
	}
}

func removeDemoFiles() {
	for i := 0; i < AMOUNT_OF_DEMO_FILES; i++ {
		os.Remove("../tests_assets/demo_file" + strconv.Itoa(i) + ".txt")
	}
}

func TestReadFilesAmount(t *testing.T) {
	createDemoFiles()
	defer removeDemoFiles()
	filesList := utils.ReadFiles()
	if len(filesList) != AMOUNT_OF_DEMO_FILES {
		t.Errorf("not enought files were read")
	}

}

func TestReadFilesNames(t *testing.T) {
	createDemoFiles()
	defer removeDemoFiles()

	filesList := utils.ReadFiles()

	for i, file := range filesList {
		expectedName := "demo_file" + strconv.Itoa(i) + ".txt"
		actualName := file.Name
		if expectedName != actualName {
			t.Errorf("file were not read correctly")
		}
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

	os.Create("../tests_assets/demo_file_exist.txt")
	defer os.Remove("../tests_assets/demo_file_exist.txt")

	tables := []struct {
		x           string
		expected    int
		description string
	}{
		{"demo_file.txt", 404, "trying to download file that doesn't exist"},
		{"demo_file_exist.txt", 200, "trying to download file that exist"},
		{"", 301, "trying to download file without a name"},
	}

	for _, table := range tables {
		router, w := setup()
		req := httptest.NewRequest("GET", "/files/"+table.x, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, table.expected, w.Code)
	}
}

func TestUpload(t *testing.T) {
	os.Create("../tests_assets/demo_file.txt")
	defer os.Remove("../tests_assets/demo_file.txt")

	router, w := setup()
	//Set up a pipe to avoid buffering
	pr, pw := io.Pipe()
	//This writers is going to transform
	//what we pass to it to multipart form data
	//and write it to our io.Pipe
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
