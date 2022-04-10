package protocol

const (
	ProtocolHttp = "http"
	// TODO: implement these protocol
	ProtocolGrpc = "grpc"
	ProtocolAmqp = "amqp"
)

type ProtocolType int

const (
	Http ProtocolType = iota
	Grpc
	Amqp
)

func NewProtocolType(protocol string) ProtocolType {
	var pt ProtocolType
	switch protocol {
	case ProtocolHttp:
		pt = Http
	case ProtocolGrpc:
		pt = Grpc
	case ProtocolAmqp:
		pt = Amqp
	}

	return pt
}

func (r ProtocolType) ToString() string {
	protocolType := ""
	switch r {
	case Http:
		protocolType = ProtocolHttp
	case Grpc:
		protocolType = ProtocolGrpc
	case Amqp:
		protocolType = ProtocolAmqp
	}
	return protocolType
}
