package volume_server

type VolumeId uint8

type Volume struct {
	Id           VolumeId
	needleMapper NeedleMapper
	backendFile  *BackendFile
}
