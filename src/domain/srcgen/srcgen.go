package srcgen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
)

// srcgen uses to create src folder
// all of the domain and routing works on src
// src/handler for the handler
// src/domains for the domain

type SrcGen interface {
	CreateSrcDirectory(dto SrcDto) error
}

type SrcGenImpl struct {
	fs          database.FileSystem
	protocolGen protocolgen.ProtocolGen
	// domainGen   domaingen.
}

func NewSrcGen(
	fs database.FileSystem,
	protocolgen protocolgen.ProtocolGen,
) *SrcGenImpl {
	return &SrcGenImpl{
		fs:          fs,
		protocolGen: protocolgen,
	}
}

func (s SrcGenImpl) CreateSrcDirectory(dto SrcDto) error {
	// create /src
	s.fs.CreateFolderIfNotExists(dto.Path)

	handlerPath := utils.ConcatDirPath(dto.Path, "handlers")
	domainPath := utils.ConcatDirPath(dto.Path, "domains")

	// create /src/handlers
	s.fs.CreateFolderIfNotExists(handlerPath)

	// create /src/domains
	s.fs.CreateFolderIfNotExists(domainPath)

	// create /src/handlers/[protocols]
	protocolDto := protocolgen.ProtocolDto{
		ProtocolType: dto.ProtocolType,
		GomodName:    dto.GomodName,
		Path:         handlerPath,
	}
	err := s.protocolGen.CreateProtocolSrc(protocolDto)
	if err != nil {
		return err
	}

	return nil
}
