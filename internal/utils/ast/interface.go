package ast

type InterfaceMethodArg struct {
	Name      string
	IsPointer bool
}

type InterfaceMethodArgList []*InterfaceMethodArg

type InterfaceMethod struct {
	Name string
	Args InterfaceMethodArgList
}

type InterfaceMethodList []*InterfaceMethod

type InterfaceSpec struct {
	Name       string
	StructName string
}

type InterfaceSpecList []*InterfaceSpec
