package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestRemoveFile(t *testing.T) {

	dir, _ := os.Getwd()

	err := ioutil.WriteFile(dir+"/pseudoDB/existFile.txt", []byte("Hello"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	tables := []struct {
		x        string
		expected bool
	}{
		{"existFile.txt", true},
		{"nonExist.txt", false},
		{"", false},
	}

	for _, table := range tables {
		result := removeFile(table.x)
		if result != table.expected {
			t.Errorf("file named: %s returned %v , expected %v.", table.x, result, table.expected)
		}
	}
}
