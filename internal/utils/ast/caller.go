package ast

import "go/token"

type CallerArgSelectorStmt struct {
	LibName  string
	DataType string
}

type CallerArgBasicLit struct {
	Kind  token.Token
	Value string
}

type CallerArgIdent struct {
	Name string
}

type CallerArg struct {
	SelectorStmt *CallerArgSelectorStmt
	BasicLit     *CallerArgBasicLit
	Ident        *CallerArgIdent
}

type CallerArgList []*CallerArg

type CallerSelecterExpr struct {
	Name     string
	Selector string
}

type CallerFunc struct {
	// Func here has 2 object
	// first is name and the second is selector
	// the Name means the name of the library and the selector is the function
	// example we have fmt.Println. the fmt means the Name and the Println means the selector
	Name     CallerSelecterExpr
	Selector string
}

type CallerSpec struct {
	Func CallerFunc
	Args CallerArgList
}
