package protocoltype

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/src/domain/emptygen"
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
) *ProtocolTypeImpl {
	return &ProtocolTypeImpl{
		ProtocolHttpGenImpl: NewProtocolHttp(fs, emptyGen),
	}
}
