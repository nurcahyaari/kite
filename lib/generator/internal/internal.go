package internal

import (
	"github.com/nurcahyaari/kite/lib/generator/internal/logger"
	"github.com/nurcahyaari/kite/lib/generator/internal/utils"
	"github.com/nurcahyaari/kite/lib/generator/protocol"
	"github.com/nurcahyaari/kite/utils/fs"
)

type InternalGen interface {
	logger.LoggerGen
	utils.UtilGen
	protocol.ProtocolGen
	CreateInternalDir() error
	CreateInternalProtocolDir() error
}

type InternalGenImpl struct {
	InternalPath string
	*logger.LoggerGenImpl
	*utils.UtilGenImpl
	*protocol.ProtocolGenImpl
}

func NewInternal(
	projectPath string,
	gomodName string,
) InternalGen {
	internalPath := fs.ConcatDirPath(projectPath, "internal")
	return &InternalGenImpl{
		InternalPath:    internalPath,
		LoggerGenImpl:   logger.NewLoggerGen(internalPath),
		UtilGenImpl:     utils.NewUtil(internalPath),
		ProtocolGenImpl: protocol.NewProtocolGen(protocol.Http.ToString(), internalPath, gomodName),
	}
}

func (s *InternalGenImpl) CreateInternalDir() error {
	return fs.CreateFolderIsNotExist(s.InternalPath)
}

func (s *InternalGenImpl) CreateInternalProtocolDir() error {
	err := s.ProtocolGenImpl.CreateInternalProtocolDir()
	if err != nil {
		return err
	}
	return nil
}
