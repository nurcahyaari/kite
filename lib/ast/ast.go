package ast

type AbstractCodeImpl struct {
	Path string
}

type AbstractCode interface {
	ReadImport() error
	AddImport() error
	ReadInterface() error
	AddInterface() error
	AddInterfaceFuncDecl() error
	ReadStruct() error
	AddStruct() error
	AddStructVarDecl() error
	ReadDISignature() error
	AddDISignature() error
}
