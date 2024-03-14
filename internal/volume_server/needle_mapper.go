package volume_server

type Offset struct {
	startOffset uint64
	endOffset   uint64
}

type NeedleId uint64

type Cookie uint64

type Needle struct {
	Cookie   Cookie
	Id       NeedleId
	Data     []byte
	Flags    byte
	Name     []byte
	Checksum uint64
}

type NeedleMapper interface {
	Get(needleId NeedleId) (Offset, error)
	Set(needle Needle) error
}
