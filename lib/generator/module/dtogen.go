package module

import "github.com/nurcahyaari/kite/utils/fs"

type DtoGen interface {
	CreateDtoDir() error
	CreateDtoFile() error
}

type DtoGenImpl struct {
	DtoPath string
}

func NewDtoGen(modulePath string) *DtoGenImpl {
	dtoPath := fs.ConcatDirPath(modulePath, "dto")
	return &DtoGenImpl{
		DtoPath: dtoPath,
	}
}

func (s *DtoGenImpl) CreateDtoDir() error {
	err := fs.CreateFolderIsNotExist(s.DtoPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *DtoGenImpl) CreateDtoFile() error {
	return nil
}
