package ast

type Comment struct {
	// FunctionName indicates that the comment is before the function
	FunctionName string
	// Value indicates the comment value
	Value string
}
