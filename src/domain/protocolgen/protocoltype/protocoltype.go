package protocoltype

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/src/domain/emptygen"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
)

type ProtocolType interface {
	ProtocolHttpGen
}

type ProtocolTypeImpl struct {
	*ProtocolHttpGenImpl
}

func NewProtocolType(
	fs database.FileSystem,
	emptyGen emptygen.EmptyGen,
	wireGen wiregen.WireGen,
) *ProtocolTypeImpl {
	return &ProtocolTypeImpl{
		ProtocolHttpGenImpl: NewProtocolHttp(fs, emptyGen, wireGen),
	}
}
