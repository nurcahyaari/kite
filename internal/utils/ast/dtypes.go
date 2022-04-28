// Currently this code is unused
// I still don't know to use my own decls data type here
// but I created this file
package ast

type DataTypes int

const (
	INT DataTypes = iota + 1
	INT8
	INT16
	INT32
	INT64
	UINT
	UINT8
	UINT16
	UINT32
	UINT64
	UINTPTR
	FLOAT
	FLOAT32
	FLOAT64
	BYTE
	STRING
	BOOL
	RUNE
	COMPLEX64
	COMPLEX128
)

func (d DataTypes) String() string {
	var s string

	switch d {
	case INT:
		s = "int"
	case INT8:
		s = "int8"
	case INT16:
		s = "int16"
	case INT32:
		s = "int32"
	case INT64:
		s = "int64"
	case UINT:
		s = "uint"
	case UINT8:
		s = "uint8"
	case UINT16:
		s = "uint16"
	case UINT32:
		s = "uint32"
	case UINT64:
		s = "uint64"
	case UINTPTR:
		s = "uintptr"
	case FLOAT:
		s = "float"
	case FLOAT32:
		s = "float32"
	case FLOAT64:
		s = "float64"
	case BYTE:
		s = "byte"
	case STRING:
		s = "string"
	case RUNE:
		s = "rune"
	case COMPLEX64:
		s = "complex64"
	case COMPLEX128:
		s = "complex128"
	}

	return s
}

type CustomDtypesProp struct {
	Name      string
	Path      string
	PathAlias string
}

type CustomDtypes CustomDtypesProp

type InterfaceDtypes string

func (i InterfaceDtypes) String() string {
	return "interface"
}
