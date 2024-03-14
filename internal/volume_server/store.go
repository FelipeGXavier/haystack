package volume_server

import "errors"

type Store struct {
	volumes   map[VolumeId]*Volume
	directory string
}

func NewStore(directory string) (*Store, error) {
	store := &Store{directory: directory}
	return store, nil
}

func (store *Store) loadVolumes() {

}

func (store *Store) addVolume(volume *Volume) error {
	if store.volumes[volume.Id] != nil {
		//
		return errors.New("this volume id already exists on this volume server")
	}
	store.volumes[volume.Id] = volume
	return nil
}
