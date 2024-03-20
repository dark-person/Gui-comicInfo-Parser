package tags

import (
	"database/sql"
	"fmt"
	"gui-comicinfo/internal/constant"
	"gui-comicinfo/internal/files"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

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
	db, err := sql.Open(constant.DatabaseType, path)
	if err != nil {
		return err
	}

	// Assign to local database
	l.db = db
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

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// Close immediately
	file.Close()

	// TODO: Add table
	return err
}
