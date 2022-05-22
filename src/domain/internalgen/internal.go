package internalgen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"

	"github.com/nurcahyaari/kite/src/domain/internalgen/gracefulgen"
	"github.com/nurcahyaari/kite/src/domain/internalgen/loggergen"
	"github.com/nurcahyaari/kite/src/domain/internalgen/utilsgen"
	protocol "github.com/nurcahyaari/kite/src/domain/protocolgen"
)

type InternalGen interface {
	CreateInternalDir(dto InternalDto) error
	CreateInternalModules(dto InternalDto) error
}

type InternalGenImpl struct {
	loggerGen   loggergen.LoggerGen
	utilGen     utilsgen.UtilGen
	protocolGen protocol.ProtocolGen
	gracefulGen gracefulgen.GracefulGen
	fs          database.FileSystem
}

func NewInternal(
	fs database.FileSystem,
	loggerGen loggergen.LoggerGen,
	utilGen utilsgen.UtilGen,
	protocolGen protocol.ProtocolGen,
	gracefulGen gracefulgen.GracefulGen,
) *InternalGenImpl {
	return &InternalGenImpl{
		fs:          fs,
		loggerGen:   loggerGen,
		utilGen:     utilGen,
		protocolGen: protocolGen,
		gracefulGen: gracefulGen,
	}
}

func (s *InternalGenImpl) CreateInternalDir(dto InternalDto) error {
	return s.fs.CreateFolderIfNotExists(dto.Path)
}

func (s InternalGenImpl) CreateInternalModules(dto InternalDto) error {
	loggerPath := utils.ConcatDirPath(dto.Path, "logger")
	utilsPath := utils.ConcatDirPath(dto.Path, "utils")
	protocolPath := utils.ConcatDirPath(dto.Path, "protocols")
	gracefulPath := utils.ConcatDirPath(dto.Path, "graceful")

	loggerGenDto := loggergen.LoggerDto{
		Path: loggerPath,
	}
	err := s.loggerGen.CreateLoggerDir(loggerGenDto)
	if err != nil {
		return err
	}
	s.loggerGen.CreateDefaultLoggerFile(loggerGenDto)

	utilDto := utilsgen.UtilDto{
		Path: utilsPath,
	}
	err = s.utilGen.CreateUtilDir(utilDto)
	if err != nil {
		return err
	}

	err = s.utilGen.CreateUtilModules(utilDto)
	if err != nil {
		return err
	}

	gracefulDto := gracefulgen.GracefulGenDto{
		ProjectPath: dto.ProjectPath,
		Path:        gracefulPath,
	}
	err = s.gracefulGen.CreateGraceful(gracefulDto)
	if err != nil {
		return err
	}

	protocolDto := protocol.ProtocolDto{
		ProtocolType: protocol.Http,
		Path:         protocolPath,
		GomodName:    dto.GomodName,
		ProjectPath:  dto.ProjectPath,
	}
	err = s.protocolGen.CreateProtocolInternalType(protocolDto)
	if err != nil {
		return err
	}

	return nil
}
