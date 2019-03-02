package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

type FileTest struct {
	Filename string
	Content  []byte
	// The field below serves as an extra check to see if this FileTest should
	// be correct or not.
	ShouldBeCorrect bool
}

var (
	filesTable []FileTest = []FileTest{
		{
			Filename: "test_1.txt",
			Content: []byte("{ 'pk': '10', 'score', '120'}\n" +
				"{ 'pk': '15', 'score', '4'}\n" +
				"{ 'pk': '25', 'score', '52561'}\n"),
			ShouldBeCorrect: true,
		},
		{
			Filename: "test_2.txt",
			Content: []byte("{ 'pk': '12', 'score', '1'}\n" +
				"{ 'pk': '', 'score', '2'}\n" +
				"{ 'pk': '256', 'score', '12937'}\n"),
			ShouldBeCorrect: true,
		},
		{
			Filename: "test_3.json",
			Content: []byte("{ 'pk': '29', 'score', '231'}\n" +
				"{ 'pk': '44', 'score', '256'}\n" +
				"{ 'pk': '16', 'score', '13'}\n"),
			ShouldBeCorrect: true,
		},
		{
			Filename: "test_4.go",
			Content: []byte("{ 'pk': '29', 'score', '231'}\n" +
				"{ 'pk': '44', 'score', '256'}\n" +
				"{ 'pk': '16', 'score', '13'}\n"),
			ShouldBeCorrect: false,
		},
	}

	filesTestFolder = "fileTest/"
)

// Tests whether the file created with the filesTable has the same contents as
// it is loaded by the load function.
func TestLoad(t *testing.T) {

	tempDir := os.TempDir()

	if !strings.HasSuffix(tempDir, "/") {
		tempDir += "/"
	}

	path := tempDir + filesTestFolder
	os.Mkdir(path, 0777)

	// Create every file within the filesTable
	for i := range filesTable {
		file, err := os.Create(path + filesTable[i].Filename)

		if err != nil {
			t.Fatalf("[!] Couldn't create the file. Cause %v", err.Error())
		}

		defer file.Close()
		_, err = file.Write(filesTable[i].Content)

		if err != nil {
			t.Fatalf("[!] Couldn't write content to the file. Cause %v",
				err.Error())
		}
	}

	files, err := Load(path, ".txt,.json")

	if err != nil {
		t.Fatalf("[!] Couldn't read the directory %s. Cause %s",
			path, err.Error())
	}

	for i := range files {
		filename := reflect.DeepEqual(files[i].Filename, filesTable[i].Filename)
		content := reflect.DeepEqual(files[i].Content, filesTable[i].Content)

		if !(filename && content) && filesTable[i].ShouldBeCorrect {
			t.Fatalf("[!] Values does not correspond, got %v wants %v",
				files[i], filesTable[i])
		}
	}

}

// Tests whether an invalid extension will return an error
func TestLoad2(t *testing.T) {

	tempDir := os.TempDir()

	if !strings.HasSuffix(tempDir, "/") {
		tempDir += "/"
	}

	path := tempDir + filesTestFolder

	// Passing the wrong extension to see if a new error is generated
	_, err := Load(path, ".txt,json")

	if err == nil {
		t.Fatalf("[!] Expected an error for wrong extension, got %v", err)
	}
}
