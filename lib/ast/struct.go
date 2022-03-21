package ast

type StructDtypes struct {
	// LibName example in fmt library
	LibName string
	// TypeName example in fmt library have type Scanner
	TypeName string
}

type StructArg struct {
	StructName string
	Name       string
	DataType   StructDtypes
	IsPointer  bool
}

type StructArgList []*StructArg

type StructSpec struct {
	Name          string
	InterfaceName string
}

type StructSpecList []*StructSpec
