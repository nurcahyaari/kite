package protocoltype

type ProtocolTypeGen interface {
	CreateProtocolInternal(dto ProtocolDto) error
	CreateProtocolSrc(dto ProtocolDto) error
	CreateProtocolSrcBaseFile(dto ProtocolDto) error
}
