package module

import (
	"github.com/nurcahyaari/kite/lib/generator/module/handler"
	"github.com/nurcahyaari/kite/utils/fs"
)

type HandlerGen interface {
	CreateBaseHandlerDir() error
	CreateHandlerFile() error
	handler.HttpHandlerGen
}

type HandlerGenImpl struct {
	HandlerPath string
	ModuleName  string
	*handler.HttpHandlerGenImpl
}

func NewHandlerGen(modulePath, moduleName, gomodName string) *HandlerGenImpl {
	HandlerPath := fs.ConcatDirPath(modulePath, "handler")
	return &HandlerGenImpl{
		HandlerPath:        HandlerPath,
		ModuleName:         moduleName,
		HttpHandlerGenImpl: handler.NewHttpHandlerGen(moduleName, HandlerPath, gomodName),
	}
}

func (s *HandlerGenImpl) CreateBaseHandlerDir() error {
	err := fs.CreateFolderIsNotExist(s.HandlerPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *HandlerGenImpl) CreateHandlerFile() error {
	s.CreateHttpHandlerBaseDir()
	return s.CreateHttpHandlerBaseFile()
}
