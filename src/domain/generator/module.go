package generator

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
)

type ModuleGen interface {
	CreateNewModule(dto ModuleNewDto) error
}

type ModuleGenImpl struct {
	fs        database.FileSystem
	moduleGen modulegen.ModuleGen
}

func NewModuleGen(
	fs database.FileSystem,
	moduleGen modulegen.ModuleGen,
) *ModuleGenImpl {
	return &ModuleGenImpl{
		fs:        fs,
		moduleGen: moduleGen,
	}
}

func (s ModuleGenImpl) CreateNewModule(dto ModuleNewDto) error {
	return s.moduleGen.CreateNewModule(modulegen.ModuleDto{
		PackageName: dto.PackageName,
		FileName:    dto.Name,
		ModuleName:  dto.Name,
		Path:        dto.ProjectPath,
		ProjectPath: utils.GetProjectPath(dto.ProjectPath, "src"),
		GomodName:   utils.GetGoModName(utils.GetProjectPath(dto.ProjectPath, "src")),
	})
}
