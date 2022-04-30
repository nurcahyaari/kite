package generator

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/src/domain/handlergen"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
)

type HandlerGen interface {
	CreateNewHandler(dto HandlerNewDto) error
}

type HandlerGenImpl struct {
	fs         database.FileSystem
	handlerGen handlergen.HandlerGen
}

func NewHandlerGen(
	fs database.FileSystem,
	handlerGen handlergen.HandlerGen,
) *HandlerGenImpl {
	return &HandlerGenImpl{
		fs:         fs,
		handlerGen: handlerGen,
	}
}

func (s HandlerGenImpl) CreateNewHandler(dto HandlerNewDto) error {
	handlerPath := utils.ConcatDirPath(dto.ProjectPath, "src", "handlers")

	return s.handlerGen.CreateHandler(handlergen.HandlerDto{
		Name:         dto.Name,
		Path:         handlerPath,
		GomodName:    dto.GoModName,
		ProtocolType: protocolgen.NewProtocolType(dto.ProtocolType),
	})
}
