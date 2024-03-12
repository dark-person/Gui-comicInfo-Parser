package locals

import (
	"gui-comicinfo/internal/locals"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Load config from specified folder. Core of `LoadConfig()` method.
// Config will be found under `{folder}/gui-comicinfo/config.yaml`.
//
// In real usage, folder should be `$HOME` directory of current user.
// This method allows customization of folder, which is not allowed,
// therefore, this method is an private method.
//
// When config file is not found, an os.ErrNotExist will be returned.
func loadConfig(folder string) error {

	configPath := filepath.Join(folder, locals.RootDir, locals.ConfigFile)

	// Check file is exist or not
	_, err := os.Stat(configPath)
	if err != nil {
		return err
	}

	viper.AddConfigPath(filepath.Join(folder, locals.RootDir))
	viper.SetConfigName(locals.ConfigName)
	viper.SetConfigType(locals.ConfigType)

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

// Load config from `$HOME` directory of current user.
func LoadConfig() error {
	home, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	return loadConfig(home)
}
