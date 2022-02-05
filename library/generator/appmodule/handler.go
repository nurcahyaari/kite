package appmodule

import (
	"github.com/nurcahyaari/kite/library/generator/misc"
)

type HandlerOption struct {
	Options
	DirPath    string
	ModuleName string
}

func NewHandler(opt HandlerOption) (AppGenerator, error) {
	return HandlerOption{
		Options:    opt.Options,
		DirPath:    opt.DirPath,
		ModuleName: opt.ModuleName,
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
		Options:    o.Options,
		DirPath:    o.DirPath,
		IsModule:   true,
		RouteType:  misc.Http.ToString(),
		ModuleName: o.ModuleName,
	}
	protocol := misc.NewProtocols(protocolOption)
	err := protocol.Run()
	if err != nil {
		return err
	}

	return nil
}
