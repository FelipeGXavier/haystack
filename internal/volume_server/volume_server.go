package volume_server

import (
	"haystack/pkg/file_util"
)

type VolumeServer struct {
	Ip    string
	Port  int
	store *Store
}

func NewVolumeServer(directory string) (*VolumeServer, error) {
	defaultIp := "localhost"
	defaultPort := 8081
	absoluteVolumeDirectory, err := file_util.SolveAbsPath(directory)
	if err != nil {

	}
	if isDirectoryWritable := file_util.IsDirectoryWritable(absoluteVolumeDirectory); !isDirectoryWritable {

	}
	store, err := NewStore(absoluteVolumeDirectory)
	if err != nil {

	}
	volumeServer := &VolumeServer{Ip: defaultIp, Port: defaultPort, store: store}
	return volumeServer, nil
}
