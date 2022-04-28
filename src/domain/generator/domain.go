package generator

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
)

type DomainGen interface {
	CreateNewDomain() error
}

type DomainGenImpl struct {
	fs database.FileSystem
}

func NewDomainGen(
	fs database.FileSystem,
) *DomainGenImpl {
	// domainPath := utils.ConcatDirPath(info.ProjectPath, "src", name)
	return &DomainGenImpl{
		// fs:         database.NewFileSystem(domainPath),
		fs: fs,
	}
}

func (s DomainGenImpl) CreateNewDomain(info ProjectInfo) error {
	return nil
}
