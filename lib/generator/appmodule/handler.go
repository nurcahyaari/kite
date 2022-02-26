package appmodule

import (
	"github.com/nurcahyaari/kite/lib/generator/misc"
	"github.com/nurcahyaari/kite/lib/impl"
)

type HandlerOption struct {
	impl.GeneratorOptions
	DirPath    string
	ModuleName string
}

func NewHandler(opt HandlerOption) (impl.AppGenerator, error) {
	return HandlerOption{
		GeneratorOptions: opt.GeneratorOptions,
		DirPath:          opt.DirPath,
		ModuleName:       opt.ModuleName,
	}, nil
}

func (o HandlerOption) Run() error {
	err := o.createHandlerPath()
	if err != nil {
		return err
	}

	return nil
}

func (o HandlerOption) createHandlerPath() error {
	protocolOption := misc.ProtocolOption{
		GeneratorOptions: o.GeneratorOptions,
		DirPath:          o.DirPath,
		IsModule:         true,
		RouteType:        misc.Http.ToString(),
		ModuleName:       o.ModuleName,
	}
	protocol := misc.NewProtocols(protocolOption)
	err := protocol.Run()
	if err != nil {
		return err
	}

	return nil
}
