package locals

import (
	"gui-comicinfo/internal/locals"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
)

// Prepare testing. This function should be called in 1st line of every test.
func prepareTest() {
	// Set Log Level & Output
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
}

func TestLoadConfig_Exist(t *testing.T) {
	prepareTest()

	// Prepare Dummy folder
	dir := t.TempDir()

	// Prepare two cases, one has config, one has not
	case1 := filepath.Join(dir, "case1")
	os.MkdirAll(case1, 0755)

	case2 := filepath.Join(dir, "case2")
	os.MkdirAll(filepath.Join(case2, locals.RootDir), 0755)

	// Create config.yaml
	tmp, _ := os.Create(filepath.Join(case2, locals.RootDir, locals.ConfigFile))
	defer tmp.Close()

	// Load config
	err1 := loadConfig(case1)
	if err1 == nil {
		t.Error("Case 1 failed. Expected ErrNotExist, return nil")
	}

	err2 := loadConfig(case2)
	if err2 != nil {
		t.Errorf("Case 1 failed. Expected nil, return %v", err2)
	}
}
