package ast

type FunctionArg struct {
	IsPointer bool
	Name      string
	LibName   string
	DataType  string
}

type FunctionArgList []*FunctionArg

type FunctionReturnArgsSpec struct {
	IsPointer bool
	// FuncName means the name of the function
	// example
	// func NewHttp(), so the FuncName = NewHttp
	FuncName string
	// ReturnName means the name of return
	// example
	// return &Http{}, so the return name is Http
	ReturnName string
	// DataTypeKey name of the data type that we want to inject
	DataTypeKey string
	// DataTypeValue name of the data type that we want to inject
	DataTypeValue string
}

// FunctionReturnSpec currently only be able to return a custom struct
// and an error. it cannot return a number or anythink
type FunctionReturnSpec struct {
	IsPointer bool
	// IsStruct define that the return is a struct, if the IsStruct is false
	// so it define it an error
	IsStruct bool
	LibName  string
	DataType string
	// Return define the variable name that will be returned
	// example
	//	func a() interface{} {
	//		var a := interface{}
	// 		return a <- a here means Return name
	//	}
	Return string
}

type FunctionReturnSpecList []*FunctionReturnSpec

type FunctionStructSpec struct {
	// Name means as the struct alias
	Name string
	// DataTypes means the struct name
	DataTypes string
	IsPointer bool
	// IsConstruct define that the function is a construct of an object
	IsConstruct bool
}

type FunctionBodySpec struct {
}

type FunctionSpec struct {
	Name string
	// StructName define that the function is a method of an object
	StructSpec *FunctionStructSpec
	Args       FunctionArgList
	Returns    *FunctionReturnSpecList
	Body       FunctionBodySpec
}

type FunctionSpecList []*FunctionSpec
