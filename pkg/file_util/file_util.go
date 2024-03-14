package file_util

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func IsDirectoryWritable(directory string) bool {
	tempFile := filepath.Join(directory, "file_util.tmp")
	_, err := os.Create(tempFile)
	if err != nil {
		return false
	}
	_ = os.Remove(tempFile)
	return true
}

func SolveAbsPath(path string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("invalid path")
	}
	if strings.Contains(path, "~") {
		currentUser, err := user.Current()
		if err != nil {
			return "", err
		}
		homePath := currentUser.HomeDir
		path = strings.Replace(path, "~", homePath, -1)
	}
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absolutePath, nil
}
