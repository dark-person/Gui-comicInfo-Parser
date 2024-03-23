package tags

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Latest Version of schema supported.
//
// Every database schema changes,
// should also change this value at same time.
const LatestSchema = 1

// Check database user version value is supported by program.
//
// Any mismatched values will consider as outdated,
// and should run migration scripts before program actually function.
func checkVer(db *sql.DB) error {
	row := db.QueryRow("PRAGMA user_version")
	if row == nil {
		return fmt.Errorf("empty user_version")
	}

	var userVersion int
	err := row.Scan(&userVersion)
	if err != nil {
		return err
	}

	// Check schema version
	logrus.Infof("DB User version: %d; Latest Support:%d", userVersion, LatestSchema)
	if userVersion > LatestSchema {
		logrus.Error("Program version is not support this version of database.")
		return fmt.Errorf("outdated program")
	}

	if userVersion < LatestSchema {
		logrus.Warn("Database version is outdated.")

		// TODO: Migration scripts
		return fmt.Errorf("outdated schema")
	}

	logrus.Info("Database version is supported.")
	return nil

}
