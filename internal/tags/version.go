package tags

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Error represents program is outdated
var ErrProgramOutdated = fmt.Errorf("outdated program")

// Error represents database schema is outdated
var ErrSchemaOutdated = fmt.Errorf("outdated schema")

// Error represents database user_version is invalid,
// usually appear in empty database
var ErrInvalidVersion = fmt.Errorf("invalid user_version")

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
		return ErrInvalidVersion
	}

	var userVersion int
	err := row.Scan(&userVersion)
	if err != nil {
		return err
	}

	// Check schema version
	logrus.Infof("DB User version: %d; Latest Support:%d", userVersion, LatestSchema)

	if userVersion == 0 {
		logrus.Error("invalid user_version")
		return ErrInvalidVersion
	}

	if userVersion > LatestSchema {
		logrus.Error("Program version is not support this version of database.")
		return ErrProgramOutdated
	}

	if userVersion < LatestSchema {
		logrus.Warn("Database version is outdated.")

		// TODO: Migration scripts
		return ErrSchemaOutdated
	}

	logrus.Info("Database version is supported.")
	return nil

}
