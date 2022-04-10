package http

type ProtocolHttpGen interface{}

type ProtocolHttpGenImpl struct {
}

func NewProtocolHttp() *ProtocolHttpGenImpl {
	return &ProtocolHttpGenImpl{}
}
