package file_util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/user"
	"testing"
)

func Test_writableDirectoryShouldReturnTrue(t *testing.T) {
	_ = os.Mkdir("temp", 0744)
	isWritable := IsDirectoryWritable("temp")
	_ = os.Remove("temp")
	assert.Equal(t, true, isWritable)
}

func Test_not_writableDirectoryShouldReturnFalse(t *testing.T) {
	isWritable := IsDirectoryWritable("99999999999999")
	assert.Equal(t, false, isWritable)
}

func Test_absolutePathFromHome(t *testing.T) {
	homePath, _ := user.Current()
	dir := ".test"
	path := "~/" + dir
	fullPath, err := SolveAbsPath(path)
	assert.Nil(t, err)
	assert.Equal(t, homePath.HomeDir+"/"+dir, fullPath)
}

func Test_absolutePathFrom(t *testing.T) {
	currentDir, _ := os.Getwd()
	dir := "dir1"
	fullPath, err := SolveAbsPath("./" + dir)
	assert.Nil(t, err)
	assert.Equal(t, currentDir+"/"+dir, fullPath)
}
