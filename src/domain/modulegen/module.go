package modulegen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

// create module, what is module
// all of these package is based on module.
// dependency in here called as a module
type ModuleGen interface {
	BuildModuleTemplate(dto ModuleDto) (string, error)
}

type ModuleGenImpl struct {
}

func NewModuleGen() *ModuleGenImpl {
	return &ModuleGenImpl{}
}

// BuildModuleTemplate will build base file of the module, such as interface struct and the construct
func (s ModuleGenImpl) BuildModuleTemplate(dto ModuleDto) (string, error) {
	templateNew := templates.NewTemplateNewImpl(dto.PackageName, "")
	templateCode, err := templateNew.Render("", nil)
	if err != nil {
		return "", err
	}

	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: fmt.Sprintf("New%s", dto.ModuleName),
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					DataType:  fmt.Sprintf("%sImpl", dto.ModuleName),
					Return:    fmt.Sprintf("%sImpl", dto.ModuleName),
				},
			},
		},
	})
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: fmt.Sprintf("%sImpl", dto.ModuleName),
		},
	})
	abstractCode.AddInterfaces(ast.InterfaceSpecList{
		&ast.InterfaceSpec{
			Name:       dto.ModuleName,
			StructName: fmt.Sprintf("%sImpl", dto.ModuleName),
		},
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return "", err
	}
	templateBaseFileString := abstractCode.GetCode()

	return templateBaseFileString, nil
}
