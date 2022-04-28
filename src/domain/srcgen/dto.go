package srcgen

import "github.com/nurcahyaari/kite/src/domain/protocolgen"

type SrcDto struct {
	Path         string
	GomodName    string
	ProtocolType protocolgen.ProtocolType
}
