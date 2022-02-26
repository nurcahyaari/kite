package generator

import (
	"github.com/nurcahyaari/kite/lib/generator/misc"
	"github.com/nurcahyaari/kite/lib/impl"

	"github.com/nurcahyaari/kite/utils/fs"
)

type InternalOption struct {
	impl.GeneratorOptions
	DirName string
	DirPath string
}

func NewInternal(opt InternalOption) impl.AppGenerator {
	opt.DirName = "internal"
	opt.DirPath = fs.ConcatDirPath(opt.ProjectPath, opt.DirName)

	return &opt
}

func (o InternalOption) Run() error {

	o.createInternalDir()

	protocolOption := misc.ProtocolOption{
		GeneratorOptions: o.GeneratorOptions,
		DirPath:          o.DirPath,
		IsModule:         false,
		RouteType:        misc.Http.ToString(),
	}
	protocol := misc.NewProtocols(protocolOption)
	err := protocol.Run()
	if err != nil {
		return err
	}

	utilOption := misc.UtilsOption{
		GeneratorOptions: o.GeneratorOptions,
		InternalPath:     o.DirPath,
	}
	util := misc.NewUtils(utilOption)
	util.Run()

	loggerOption := misc.LoggerOption{
		GeneratorOptions:   o.GeneratorOptions,
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
