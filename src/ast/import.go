package ast

type ImportSpec struct {
	Name string
	Path string
}

type ImportSpecList []*ImportSpec
