package protocol

const (
	ProtocolHttp = "http"
)

type ProtocolType int

const (
	Http ProtocolType = iota
)

func NewProtocolType(protocol string) ProtocolType {
	var pt ProtocolType
	switch protocol {
	case ProtocolHttp:
		pt = Http
	}

	return pt
}

func (r ProtocolType) ToString() string {
	protocolType := ""
	switch r {
	case Http:
		protocolType = ProtocolHttp
	}
	return protocolType
}
