package entitygen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/src/domain/emptygen"
)

type EntityGen interface {
	CreateEntityDir(dto EntityGenDto) error
	CreateEntityFile(dto EntityGenDto) error
	CreateEntity(dto EntityGenDto) error
}

type EntityGenImpl struct {
	fs       database.FileSystem
	emptyGen emptygen.EmptyGen
}

func NewEntityGen(
	fs database.FileSystem,
	emptyGen emptygen.EmptyGen,
) *EntityGenImpl {
	return &EntityGenImpl{
		fs:       fs,
		emptyGen: emptyGen,
	}
}

func (s EntityGenImpl) CreateEntityDir(dto EntityGenDto) error {
	return s.fs.CreateFolderIfNotExists(dto.Path)
}

func (s EntityGenImpl) CreateEntityFile(dto EntityGenDto) error {
	return s.emptyGen.CreateEmptyGolangFile(emptygen.EmptyDto{
		Path:        dto.Path,
		FileName:    "entity",
		PackageName: "entity",
	})
}

func (s EntityGenImpl) CreateEntity(dto EntityGenDto) error {
	err := s.CreateEntityDir(dto)
	if err != nil {
		return err
	}
	return s.CreateEntityFile(dto)
}
