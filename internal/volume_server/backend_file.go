package volume_server

import "os"

// .dat
type BackendFile struct {
	File         *os.File
	fullFilePath string
	fileSize     uint64
}
