package archive

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

// Test Rename Zip archive to CBZ archive
func TestRenameZip(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Create a folder inside temp directory
	os.MkdirAll(filepath.Join(tempDir, "tmp"), 0755)

	// Create a zip file
	file1, _ := os.Create(filepath.Join(tempDir, "tmp", "hello.zip"))
	file1.Close()

	// Test Function
	err := RenameZip(filepath.Join(tempDir, "tmp", "hello.zip"))
	if err != nil {
		t.Error(err)
	}

	// Result Verify
	dest, openErr := os.Open(filepath.Join(tempDir, "tmp", "hello", "hello.cbz"))
	if openErr != nil {
		t.Error(openErr)
	}
	defer dest.Close()
}

// Test Create Zip from folder
func TestCreateZip(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Create a set of file
	file1, _ := os.Create(filepath.Join(tempDir, "image1.jpg"))
	file2, _ := os.Create(filepath.Join(tempDir, "image2.jpg"))
	file3, _ := os.Create(filepath.Join(tempDir, "image3.jpg"))
	file4, _ := os.Create(filepath.Join(tempDir, "test.xml"))
	defer file1.Close()
	defer file2.Close()
	defer file3.Close()
	defer file4.Close()

	// Start Testing Functions
	dest, err := CreateZip(tempDir)
	if err != nil {
		t.Error(err)
	}

	// Check Dest Filename
	destFileName := filepath.Base(tempDir)
	if dest != filepath.Join(tempDir, destFileName+".zip") {
		t.Errorf("Error Destination file: %v", dest)
	}

	// Check Zip Content
	reader, err := zip.OpenReader(dest)
	if err != nil {
		t.Error(err)
	}
	defer reader.Close()

	list := make(map[string]int, 0)
	for _, f := range reader.File {
		list[f.Name] = 1
	}

	_, exist1 := list["test.xml"]
	_, exist2 := list["image1.jpg"]
	_, exist3 := list["image2.jpg"]
	_, exist4 := list["image3.jpg"]

	if !exist1 {
		t.Error("Content 1 missing in zip")
	} else if !exist2 {
		t.Error("Content 2 missing in zip")
	} else if !exist3 {
		t.Error("Content 3 missing in zip")
	} else if !exist4 {
		t.Error("Content 4 missing in zip")
	}
}
