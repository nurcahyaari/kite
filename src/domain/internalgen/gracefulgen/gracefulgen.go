package gracefulgen

import (
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates/internaltemplate/gracefultemplate"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

type GracefulGen interface {
	CreateGraceful(dto GracefulGenDto) error
	CreateGracefulDir(dto GracefulGenDto) error
	CreateGracefulFile(dto GracefulGenDto) error
}

type GracefulGenImpl struct {
	fs database.FileSystem
}

func NewGracefulGen(fs database.FileSystem) *GracefulGenImpl {
	return &GracefulGenImpl{
		fs: fs,
	}
}

func (s GracefulGenImpl) CreateGraceful(dto GracefulGenDto) error {
	err := s.CreateGracefulDir(dto)
	if err != nil {
		return err
	}

	return s.CreateGracefulFile(dto)
}

func (s GracefulGenImpl) CreateGracefulDir(dto GracefulGenDto) error {
	return s.fs.CreateFolderIfNotExists(dto.Path)
}

func (s GracefulGenImpl) CreateGracefulFile(dto GracefulGenDto) error {
	gracefulNew := gracefultemplate.NewGracefulTemplate()
	gracefulTemplate, err := gracefulNew.Render()
	if err != nil {
		return err
	}

	gracefulAbstractCode := ast.NewAbstractCode(gracefulTemplate, parser.ParseComments)
	gracefulAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"context\"",
	})
	gracefulAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"os\"",
	})
	gracefulAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"os/signal\"",
	})
	gracefulAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"sync\"",
	})
	gracefulAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"syscall\"",
	})
	gracefulAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"time\"",
	})
	gracefulAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/rs/zerolog/log\"",
	})
	// after manipulate the code, rebuild the code
	err = gracefulAbstractCode.RebuildCode()
	if err != nil {
		return err
	}
	// get the manipulate code
	gracefulCode := gracefulAbstractCode.GetCode()

	return s.fs.CreateFileIfNotExists(dto.Path, "graceful.go", gracefulCode)
}
