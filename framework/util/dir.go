package util

import (
	"os"
	"path/filepath"
)

func GetExecDirectory() string {
	if dir, err := os.Getwd(); err == nil {
		return filepath.Join(dir, "/")
	}
	return ""
}
