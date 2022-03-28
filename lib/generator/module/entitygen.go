package module

import "github.com/nurcahyaari/kite/utils/fs"

type EntityGen interface {
	CreateEntityDir() error
	CreateEntityFile() error
}

type EntityGenImpl struct {
	EntityPath string
}

func NewEntityGen(modulePath string) *EntityGenImpl {
	entityPath := fs.ConcatDirPath(modulePath, "entity")
	return &EntityGenImpl{
		EntityPath: entityPath,
	}
}

func (s *EntityGenImpl) CreateEntityDir() error {
	err := fs.CreateFolderIsNotExist(s.EntityPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *EntityGenImpl) CreateEntityFile() error {
	return nil
}
