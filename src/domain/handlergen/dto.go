package handlergen

import "github.com/nurcahyaari/kite/src/domain/protocolgen"

type HandlerDto struct {
	Name         string
	GomodName    string
	Path         string
	ProtocolType protocolgen.ProtocolType
}
