package tags

import (
	"database/sql"
	"fmt"
	"gui-comicinfo/internal/constant"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// Testing Purpose ONLY. Create Database file with specified `user_version` value.
//
// The returned *sql.DB will remain opened state.
// Developer MUST ensure Close() is called after usage.
func createDbWithVersion(folder string, dbName string, version int) *sql.DB {
	// Get Path
	path := filepath.Join(folder, dbName)

	// Create database file
	file, _ := os.Create(path)
	file.Close()

	// Open Database
	db, _ := sql.Open(constant.DatabaseType, path)

	// Run PRAGMA user_version
	_, e := db.Exec(fmt.Sprintf("PRAGMA user_version = %d", version))
	if e != nil {
		logrus.Error(e)
	}

	return db
}

func Test_checkVer(t *testing.T) {
	// Prepare temp directory
	dir := "testing"

	// Test Case Struct
	type testCase struct {
		db          *sql.DB // Opened database connection
		expectedMsg string  // Expected error message. Empty string if nil.
	}

	// Prepare test case
	tests := []testCase{
		// 0. Graceful test case
		{createDbWithVersion(dir, "case1.db", LatestSchema), ""},
		// 1. Database version > Program Version
		{createDbWithVersion(dir, "case2.db", LatestSchema+1), ErrProgramOutdated.Error()},
		// 2. Database version < Program Version
		{createDbWithVersion(dir, "case3.db", -1), ErrSchemaOutdated.Error()},
		// 3. Database version = 0
		{createDbWithVersion(dir, "case4.db", 0), ErrInvalidVersion.Error()},
		// 4. Nil *sql.DB
		{nil, ErrNilDatabase.Error()},
	}

	for _, tt := range tests {
		err := checkVer(tt.db)

		if tt.db != nil {
			defer tt.db.Close()
		}

		// Check if error is nil (only when expectedMsg is empty)
		if tt.expectedMsg == "" {
			assert.Nil(t, err, "Expected nil error, got %v", err)
			continue
		}

		// Check error content
		assert.EqualErrorf(t, err, tt.expectedMsg, "Unexpected Error: %v", err)
	}
}
