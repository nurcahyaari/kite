package generator

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/src/domain/domaingen"
)

type DomainGen interface {
	CreateNewDomain(dto DomainNewDto) error
}

type DomainGenImpl struct {
	fs        database.FileSystem
	domainGen domaingen.DomainGen
}

func NewDomainGen(
	fs database.FileSystem,
	domainGen domaingen.DomainGen,
) *DomainGenImpl {
	return &DomainGenImpl{
		fs:        fs,
		domainGen: domainGen,
	}
}

func (s DomainGenImpl) CreateNewDomain(dto DomainNewDto) error {
	domainCreationalType := domaingen.TypeDomainFullCreation
	if !dto.IsDomainFullCreational {
		domainCreationalType = domaingen.TypeDomainFolderOnlyCreation
	}

	err := s.domainGen.CreateDomain(domaingen.DomainDto{
		Name:                 dto.Name,
		Path:                 dto.ProjectPath,
		GomodName:            dto.GoModName,
		DomainCreationalType: domaingen.NewDomainCreationalType(domainCreationalType),
	})

	return err
}
