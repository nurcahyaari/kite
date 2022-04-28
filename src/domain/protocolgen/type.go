package protocolgen

const (
	// HTTP is default protocol
	ProtocolHttp = "http"
	// TODO: implement these protocol
	ProtocolGrpc = "grpc"
	ProtocolAmqp = "amqp"
	ProtocolCli  = "cli"
	ProcolEmpty  = "without"
)

type ProtocolType int

const (
	Http ProtocolType = iota + 1
	Grpc
	Amqp
	Cli
	Empty
)

func NewProtocolType(protocol string) ProtocolType {
	var pt ProtocolType
	switch protocol {
	case ProtocolGrpc:
		pt = Grpc
	case ProtocolAmqp:
		pt = Amqp
	case ProtocolCli:
		pt = Cli
	case ProcolEmpty:
		pt = Empty
	default:
		pt = Http
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

func (r ProtocolType) NotEmpty() bool {
	if r > 0 {
		return true
	}
	return false
}
