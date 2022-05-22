package domaingen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/src/domain/domaingen/dtogen"
	"github.com/nurcahyaari/kite/src/domain/domaingen/entitygen"
	"github.com/nurcahyaari/kite/src/domain/domaingen/repositorygen"
	"github.com/nurcahyaari/kite/src/domain/domaingen/servicegen"
	"github.com/nurcahyaari/kite/src/domain/emptygen"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
)

type DomainGen interface {
	repositorygen.RepositoryGen
	servicegen.ServiceGen
	dtogen.DtoGen
	entitygen.EntityGen
	CreateDomain(dto DomainDto) error
	createDomainFull(dto DomainDto) error
	createDomainFolderOnly(path string) error
}

type DomainGenImpl struct {
	fs database.FileSystem
	*repositorygen.RepositoryGenImpl
	*servicegen.ServiceGenImpl
	*dtogen.DtoGenImpl
	*entitygen.EntityGenImpl
}

func NewDomainGen(
	fs database.FileSystem,
	moduleGen modulegen.ModuleGen,
	protocolGen protocolgen.ProtocolGen,
	wireGen wiregen.WireGen,
	emptyGen emptygen.EmptyGen,
) *DomainGenImpl {
	return &DomainGenImpl{
		fs:                fs,
		RepositoryGenImpl: repositorygen.NewRepositoryGen(fs, moduleGen, wireGen),
		ServiceGenImpl:    servicegen.NewServiceGen(fs, moduleGen, protocolGen, wireGen),
		DtoGenImpl:        dtogen.NewEntityGen(fs, emptyGen),
		EntityGenImpl:     entitygen.NewEntityGen(fs, emptyGen),
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
	dtoPath := utils.ConcatDirPath(dto.Path, "dto")
	entityPath := utils.ConcatDirPath(dto.Path, "entity")

	dtoDto := dtogen.DtoGenDto{
		Path:       dtoPath,
		DomainName: dto.Name,
	}
	err := s.DtoGenImpl.CreateDto(dtoDto)
	if err != nil {
		return err
	}

	entityDto := entitygen.EntityGenDto{
		Path:       entityPath,
		DomainName: dto.Name,
	}
	err = s.EntityGenImpl.CreateEntity(entityDto)
	if err != nil {
		return err
	}

	repoDto := repositorygen.RepositoryDto{
		Path:        repoPath,
		ProjectPath: dto.ProjectPath,
		GomodName:   dto.GomodName,
		DomainName:  dto.Name,
	}
	err = s.RepositoryGenImpl.CreateRepository(repoDto)
	if err != nil {
		return err
	}

	serviceDto := servicegen.ServiceDto{
		Path:              servicePath,
		ProjectPath:       dto.ProjectPath,
		GomodName:         dto.GomodName,
		DomainName:        dto.Name,
		IsInjectRepo:      true,
		IsInjectToHandler: true,
	}
	err = s.ServiceGenImpl.CreateService(serviceDto)
	if err != nil {
		return err
	}

	return nil
}

func (s DomainGenImpl) createDomainFolderOnly(path string) error {
	err := s.fs.CreateFolderIfNotExists(path)
	if err != nil {
		return err
	}

	return nil
}
