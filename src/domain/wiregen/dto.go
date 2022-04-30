package wiregen

import "github.com/nurcahyaari/kite/internal/utils/ast"

type WireDto struct {
	ProjectPath string
	// function name to inject modules
	FunctionName string
	GomodName    string
}

type WireEntryPointDto struct {
	WireDto
	Import ast.ImportSpec
	Return *ast.FunctionReturnSpecList
}

type WireAddModuleDto struct {
	WireDto
	Import     ast.ImportSpec
	Dependency ast.WireDependencyInjection
}
