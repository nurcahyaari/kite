package model

import (
	"github.com/nurcahyaari/kite/utils/fs"
)

type EntityOption struct {
	Options
	ModulesPath      string
	ModuleEntityPath string
}

func NewEntity(opt EntityOption) (AppGenerator, error) {
	entityPath := fs.ConcatDirPath(opt.ModulesPath, "entity")
	return &EntityOption{
		Options:          opt.Options,
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
