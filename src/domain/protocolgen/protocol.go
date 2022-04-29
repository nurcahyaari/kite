package protocolgen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/src/domain/protocolgen/protocolhttpgen"
)

type ProtocolGen interface {
	createProtocolDir(dto ProtocolDto) (string, error)
	CreateProtocolInternalType(dto ProtocolDto) error
	CreateProtocolSrc(dto ProtocolDto) error
}

type ProtocolGenImpl struct {
	fs           database.FileSystem
	httpProtocol protocolhttpgen.ProtocolHttpGen
}

// default protocol is http
func NewProtocolGen(
	fs database.FileSystem,
	httpProtocol protocolhttpgen.ProtocolHttpGen,
) *ProtocolGenImpl {
	return &ProtocolGenImpl{
		fs:           fs,
		httpProtocol: httpProtocol,
	}
}

func (s *ProtocolGenImpl) createProtocolDir(dto ProtocolDto) (string, error) {
	dirPath := utils.ConcatDirPath(dto.Path, dto.ProtocolType.ToString())
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return "", err
	}
	return dirPath, nil
}

func (s *ProtocolGenImpl) CreateProtocolInternalType(dto ProtocolDto) error {
	dirPath, err := s.createProtocolDir(dto)
	if err != nil {
		return err
	}

	if !s.fs.IsFolderExists(dirPath) {
		s.fs.CreateFolderIfNotExists(dirPath)
	}

	switch dto.ProtocolType {
	case Http:
		err = s.httpProtocol.CreateProtocolInternalHttp(dirPath)
	}

	if err != nil {
		return err
	}

	return nil
}

// create protocol directory. under src or under internal
func (s *ProtocolGenImpl) CreateProtocolSrc(dto ProtocolDto) error {
	var err error

	if dto.ProtocolType.NotEmpty() {
		path := utils.ConcatDirPath(dto.Path, dto.ProtocolType.ToString())

		if !s.fs.IsFolderExists(path) {
			s.fs.CreateFolderIfNotExists(path)
		}

		switch dto.ProtocolType {
		case Http:
			err = s.httpProtocol.CreateProtocolSrcHttp(path)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
