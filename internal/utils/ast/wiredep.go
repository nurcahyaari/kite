package ast

// GlobDecl here actually uses to create a dependency injector for wire
type WireDependencyInjection struct {
	VarName                   string
	TargetInjectName          string
	TargetInjectConstructName string
	InterfaceLib              string
	InterfaceName             string
	StructLib                 string
	StructName                string
}
