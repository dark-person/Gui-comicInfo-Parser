package tags

import (
	"database/sql"
	"embed"
	"fmt"
	"gui-comicinfo/internal/constant"
	"gui-comicinfo/internal/files"
	"os"
	"path/filepath"
	"sort"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

//go:embed sql/*.sql
var sqlScript embed.FS

type LocalTags struct {
	db *sql.DB
}

func NewLocalTags() (*LocalTags, error) {
	return new(constant.DatabaseFile)
}

func new(path string) (*LocalTags, error) {
	// If database file not exist, then create a new db file
	if !files.IsFileExist(path) {
		err := createDb(path)
		if err != nil {
			return nil, err
		}
	}

	// Return
	return &LocalTags{}, nil
}

func (l *LocalTags) connectTo(path string) error {
	logrus.Infof("Connecting to database: %s", path)

	if !files.IsFileExist(path) {
		return os.ErrNotExist
	}

	// Open Database connection
	var err error
	l.db, err = sql.Open(constant.DatabaseType, path)
	if err != nil {
		return err
	}

	// Test DB connection by user version
	err = checkVer(l.db)
	if err != nil {
		return err
	}

	return nil
}

func (l *LocalTags) Connect() error {
	// Get Home Directory
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Create database file if not exist
	path := filepath.Join(home, constant.RootDir, constant.DatabaseFile)
	if !files.IsFileExist(path) {
		createDb(path)
	}

	// Use private connect method
	return l.connectTo(path)
}

func (l *LocalTags) Close() error {
	return l.db.Close()
}

func createDb(path string) error {
	// Prevent Not database file
	if filepath.Ext(path) != ".db" {
		return fmt.Errorf("invalid database path")
	}

	// Prevent already existing database
	if files.IsFileExist(path) {
		return fmt.Errorf("database already exists")
	}

	// Create Connection, which will create file if not exist
	db, err := sql.Open(constant.DatabaseType, path)
	if err != nil {
		return err
	}

	// Get all sql files in embedded file system
	files, err := sqlScript.ReadDir("sql")
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// Run all script available
	for _, file := range files {
		// Not using filepath.Join as embedded file system not support window slashes
		path := "sql/" + file.Name()

		// Get SQL file content as string
		str, err := sqlScript.ReadFile(path)
		if err != nil {
			db.Close()
			return err
		}

		_, err = db.Exec(string(str))
		if err != nil {
			db.Close()
			return err
		}
	}

	// Close Database connection
	db.Close()
	return err
}
