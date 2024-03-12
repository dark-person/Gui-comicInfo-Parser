package locals

// This file will stored the default file location for
// 1. config.yaml
// 2. database file

// Root Directory to storing all config files.
const RootDir = "gui-comicinfo"

// Config filename, without extension. Used by `viper`.
const ConfigName = "config"

// Config file type. Used by `viper`.
const ConfigType = "yaml"

// Config filename, with extension. Used to check file existence.
const ConfigFile = ConfigName + "." + ConfigType

// Database filename.
const DatabaseFile = "storage.db"

// Database type.
const DatabaseType = "sqlite3"
