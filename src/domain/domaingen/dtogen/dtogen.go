package dtogen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/src/domain/emptygen"
)

type DtoGen interface {
	CreateDtoDir(dto DtoGenDto) error
	CreateDtoFile(dto DtoGenDto) error
	CreateDto(dto DtoGenDto) error
}

type DtoGenImpl struct {
	fs       database.FileSystem
	emptyGen emptygen.EmptyGen
}

func NewEntityGen(
	fs database.FileSystem,
	emptyGen emptygen.EmptyGen,
) *DtoGenImpl {
	return &DtoGenImpl{
		fs:       fs,
		emptyGen: emptyGen,
	}
}

func (s DtoGenImpl) CreateDtoDir(dto DtoGenDto) error {
	return s.fs.CreateFolderIfNotExists(dto.Path)
}

func (s DtoGenImpl) CreateDtoFile(dto DtoGenDto) error {
	return s.emptyGen.CreateEmptyGolangFile(emptygen.EmptyDto{
		Path:        dto.Path,
		FileName:    "dto",
		PackageName: "dto",
	})
}

func (s DtoGenImpl) CreateDto(dto DtoGenDto) error {
	err := s.CreateDtoDir(dto)
	if err != nil {
		return err
	}
	return s.CreateDtoFile(dto)
}
