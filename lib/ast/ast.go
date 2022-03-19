package ast

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	"github.com/nurcahyaari/kite/utils/logger"
)

type AbstractCodeImpl struct {
	code string
	fset *token.FileSet
	file *ast.File
}

type AbstractCode interface {
	GetCode() string
	AddImport(importSpec ImportSpec)
	AddInterface() error
	AddInterfaceFuncDecl() error
	AddStruct() error
	AddStructVarDecl() error
	AddGlobalDecl() error
	AddComments() error
	RebuildCode() error
}

func NewAbstractCode(code string, parserMode parser.Mode) AbstractCode {
	fset := token.NewFileSet()
	// parse source code from the string
	file, err := parser.ParseFile(fset, "", code, parserMode)
	if err != nil {
		logger.Errorln("Error when parse code")
		return nil
	}
	return &AbstractCodeImpl{
		code: code,
		file: file,
		fset: fset,
	}
}

func (a *AbstractCodeImpl) GetCode() string {
	return a.code
}

func (a *AbstractCodeImpl) AddImport(importSpec ImportSpec) {
	for _, decl := range a.file.Decls {
		gendecl, ok := decl.(*ast.GenDecl)
		if ok {
			for _, genspecs := range gendecl.Specs {
				_, ok := genspecs.(*ast.ImportSpec)
				if ok {
					gendecl.Specs = append(gendecl.Specs, &ast.ImportSpec{
						Name: &ast.Ident{
							Name: importSpec.Name,
						},
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: importSpec.Path,
						},
					})
					// breaking after importing in the import spec
					break
				}
				// if not ok means that the genspec is not ast.ImportSpec
				// so break this
				break
			}
		}
	}
}

func (a *AbstractCodeImpl) AddInterface() error {
	return nil
}

func (a *AbstractCodeImpl) AddInterfaceFuncDecl() error {
	return nil
}

func (a *AbstractCodeImpl) AddStruct() error {
	return nil
}

func (a *AbstractCodeImpl) AddStructVarDecl() error {
	return nil
}

func (a *AbstractCodeImpl) AddGlobalDecl() error {
	return nil
}

func (a *AbstractCodeImpl) AddComments() error {
	return nil
}

func (a *AbstractCodeImpl) RebuildCode() error {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, a.fset, a.file)
	if err != nil {
		logger.Error(err)
		return err
	}

	a.code = buf.String()
	return nil
}
