package misc

import (
	"github.com/nurcahyaari/kite/lib/impl"
	"github.com/nurcahyaari/kite/utils/fs"
)

type EntityOption struct {
	impl.GeneratorOptions
	ModulesPath      string
	ModuleEntityPath string
}

func NewEntity(opt EntityOption) (impl.AppGenerator, error) {
	entityPath := fs.ConcatDirPath(opt.ModulesPath, "entity")
	return &EntityOption{
		GeneratorOptions: opt.GeneratorOptions,
		ModulesPath:      opt.ModulesPath,
		ModuleEntityPath: entityPath,
	}, nil
}

func (o EntityOption) Run() error {
	err := o.createEntityDir()
	if err != nil {
		return err
	}

	return nil
}

func (o EntityOption) createEntityDir() error {
	err := fs.CreateFolderIsNotExist(o.ModuleEntityPath)
	if err != nil {
		return err
	}

	return nil
}

func (o EntityOption) createBaseFile() error {
	return nil
}
