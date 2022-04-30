package handlergen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
)

type HandlerGen interface {
	CreateHandler(dto HandlerDto) error
}

type HandlerGenImpl struct {
	fs          database.FileSystem
	protocolGen protocolgen.ProtocolGen
}

func NewHandlerGen(
	fs database.FileSystem,
	protocolGen protocolgen.ProtocolGen,
) *HandlerGenImpl {
	return &HandlerGenImpl{
		fs:          fs,
		protocolGen: protocolGen,
	}
}

func (s HandlerGenImpl) CreateHandler(dto HandlerDto) error {
	return s.protocolGen.CreateProtocolSrcHandler(protocolgen.ProtocolDto{
		Path:         dto.Path,
		Name:         dto.Name,
		GomodName:    dto.GomodName,
		ProtocolType: dto.ProtocolType,
	})
}
