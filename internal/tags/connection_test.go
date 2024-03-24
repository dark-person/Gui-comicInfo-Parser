package tags

import (
	"database/sql"
	"fmt"
	"gui-comicinfo/internal/constant"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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
	dir := t.TempDir()

	// Case 1: Connect to existed database (Graceful)
	path1 := filepath.Join(dir, "case1.db")
	f1, _ := os.Create(path1)
	f1.Close()

	db, _ := sql.Open(constant.DatabaseType, path1)
	_, e := db.Exec(fmt.Sprintf("PRAGMA user_version = %d", LatestSchema))
	if e != nil {
		logrus.Error(e)
	}
	db.Close()

	// Case 2: Connect to database that is not existing
	path2 := filepath.Join(dir, "case2.db")

	// Case 3: Connect to existed invalid database
	path3 := filepath.Join(dir, "case3.db")
	f3, _ := os.Create(path3)
	defer f3.Close()

	// Test Case structure
	type testCase struct {
		path    string
		wantErr bool
	}

	// Prepare tests
	tests := []testCase{
		{path1, false},
		{path2, true},
		{path3, true},
	}

	// Start Test
	for idx, tt := range tests {
		lt := &LocalTags{}
		err := lt.connectTo(tt.path)

		if lt.db != nil {
			lt.db.Close()
		}

		assert.EqualValuesf(t, tt.wantErr, err != nil, "Case %d: Error expected %v, got %v", idx, tt.wantErr, err)
	}
}

func Test_createDb(t *testing.T) {

	err := createDb("temp/abc.db")
	if err != nil {
		logrus.Error(err)
	}

	// type args struct {
	// 	path string
	// }
	// tests := []struct {
	// 	name    string
	// 	args    args
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if err := createDb(tt.args.path); (err != nil) != tt.wantErr {
	// 			t.Errorf("createDb() error = %v, wantErr %v", err, tt.wantErr)
	// 		}
	// 	})
	// }
}
