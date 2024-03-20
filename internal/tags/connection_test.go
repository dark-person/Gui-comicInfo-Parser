package tags

import (
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

// Prepare testing. This function should be called in 1st line of every test.
func prepareTest() {
	// Set Log Level & Output
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
}

func TestLocalTags_connectTo(t *testing.T) {
	prepareTest()

	// Prepare dummy folder
	dir := "testing"
	os.MkdirAll(dir, 0755)

	// Case 1: Connect to existed database (Graceful)
	path1 := filepath.Join(dir, "case1.db")
	f1, _ := os.Create(path1)
	defer f1.Close()

	// Case 2: Connect to database that is not existing
	path2 := filepath.Join(dir, "case2.db")

	// Temp LocalTags struct
	lt := &LocalTags{}

	// Start Test
	err1 := lt.connectTo(path1)
	if err1 != nil {
		t.Errorf("%s LocalTags.connectTo() error = %v, wantErr %v", "Case 1", err1, nil)
	}

	err2 := lt.connectTo(path2)
	if err2 == nil || !os.IsNotExist(err2) {
		t.Errorf("%s LocalTags.connectTo() error = %v, expect error is not exist.", "Case 2", err2)
	}
}
