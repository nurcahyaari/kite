package generator

import (
	"github.com/nurcahyaari/kite/library/generator/misc"

	"github.com/nurcahyaari/kite/utils/fs"
)

type InternalOption struct {
	GeneratorOptions
	DirName string
	DirPath string
}

func NewInternal(opt InternalOption) AppGenerator {
	opt.DirName = "internal"
	opt.DirPath = fs.ConcatDirPath(opt.ProjectPath, opt.DirName)

	return &opt
}

func (o InternalOption) Run() error {

	o.createInternalDir()

	protocolOption := misc.ProtocolOption{
		Options:   o.GeneratorOptions,
		DirPath:   o.DirPath,
		IsModule:  false,
		RouteType: misc.Http.ToString(),
	}
	protocol := misc.NewProtocols(protocolOption)
	err := protocol.Run()
	if err != nil {
		return err
	}

	utilOption := misc.UtilsOption{
		Options:      o.GeneratorOptions,
		InternalPath: o.DirPath,
	}
	util := misc.NewUtils(utilOption)
	util.Run()

	loggerOption := misc.LoggerOption{
		Options:            o.GeneratorOptions,
		InfrastructurePath: o.DirPath,
	}
	db := misc.NewLogger(loggerOption)
	db.Run()

	return nil
}

func (o InternalOption) createInternalDir() error {
	err := fs.CreateFolderIsNotExist(o.DirPath)
	if err != nil {
		return err
	}

	return nil
}
