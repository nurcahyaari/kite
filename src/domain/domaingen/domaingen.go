package domaingen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/src/domain/domaingen/repositorygen"
	"github.com/nurcahyaari/kite/src/domain/domaingen/servicegen"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
)

type DomainGen interface {
	repositorygen.RepositoryGen
	servicegen.ServiceGen
	CreateDomain(dto DomainDto) error
	createDomainFull(dto DomainDto) error
	createDomainFolderOnly(path string) error
}

type DomainGenImpl struct {
	fs database.FileSystem
	*repositorygen.RepositoryGenImpl
	*servicegen.ServiceGenImpl
}

func NewDomainGen(
	fs database.FileSystem,
	moduleGen modulegen.ModuleGen,
) *DomainGenImpl {
	return &DomainGenImpl{
		fs:                fs,
		RepositoryGenImpl: repositorygen.NewRepositoryGen(fs, moduleGen),
		ServiceGenImpl:    servicegen.NewServiceGen(fs, moduleGen),
	}
}

func (s DomainGenImpl) CreateDomain(dto DomainDto) error {
	var err error

	switch dto.DomainCreationalType {
	case DomainFullCreation:
		err = s.createDomainFull(dto)
	case DomainFolderOnlyCreation:
		err = s.createDomainFolderOnly(dto.Path)
	}

	return err
}

func (s DomainGenImpl) createDomainFull(dto DomainDto) error {
	if s.fs.IsFolderExists(dto.Path) {
		return database.PrintFsErr(database.FolderWasCreated, dto.Path)
	}
	repoPath := utils.ConcatDirPath(dto.Path, "repository")
	servicePath := utils.ConcatDirPath(dto.Path, "service")

	repoDto := repositorygen.RepositoryDto{
		Path:      repoPath,
		GomodName: dto.GomodName,
	}
	s.RepositoryGenImpl.CreateRepository(repoDto)

	serviceDto := servicegen.ServiceDto{
		Path:       servicePath,
		GomodName:  dto.GomodName,
		DomainName: dto.Name,
	}
	s.ServiceGenImpl.CreateService(serviceDto)

	return nil
}

func (s DomainGenImpl) createDomainFolderOnly(path string) error {
	err := s.fs.CreateFolderIfNotExists(path)
	if err != nil {
		return err
	}

	return nil
}
