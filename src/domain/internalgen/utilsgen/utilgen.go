package utilsgen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
)

type UtilGen interface {
	CreateUtilDir(dto UtilDto) error
	CreateUtilModules(dto UtilDto) error
}

type UtilGenImpl struct {
	fs  database.FileSystem
	enc UtilEncryption
}

func NewUtil(
	fs database.FileSystem,
	enc UtilEncryption,
) *UtilGenImpl {
	return &UtilGenImpl{
		enc: enc,
		fs:  fs,
	}
}

func (s *UtilGenImpl) CreateUtilDir(dto UtilDto) error {
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return err
	}

	return nil
}

func (s *UtilGenImpl) CreateUtilDefaultFile(dto UtilDto) error {
	// create empty utils file
	// temp
	return nil
}

func (s *UtilGenImpl) CreateUtilModules(dto UtilDto) error {
	encryptionPath := utils.ConcatDirPath(dto.Path, "encryption")

	s.enc.CreateRsaReader(EncryptionDto{
		Path: encryptionPath,
	})

	return nil
}
