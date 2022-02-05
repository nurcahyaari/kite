package model

import (
	"github.com/nurcahyaari/kite/utils/fs"
)

type DTOOption struct {
	Options
	ModulesPath      string
	ModuleEntityPath string
}

func NewDTO(opt DTOOption) (AppGenerator, error) {
	entityPath := fs.ConcatDirPath(opt.ModulesPath, "dto")
	return &DTOOption{
		Options:          opt.Options,
		ModulesPath:      opt.ModulesPath,
		ModuleEntityPath: entityPath,
	}, nil
}

func (o DTOOption) Run() error {
	err := o.createDTODir()
	if err != nil {
		return err
	}

	return nil
}

func (o DTOOption) createDTODir() error {
	err := fs.CreateFolderIsNotExist(o.ModuleEntityPath)
	if err != nil {
		return err
	}

	return nil
}

func (o DTOOption) createBaseFile() error {
	return nil
}
