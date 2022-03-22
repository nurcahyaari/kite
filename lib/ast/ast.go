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
	AddInterfaces(interfaceSpecs InterfaceSpecList)
	AddInterfaceFuncDecl() error
	AddStructs(structSpecs StructSpecList)
	AddStructVarDecl(structArgs StructArgList)
	AddWireDependencyInjection(wireDependency WireDependencyInjection)
	AddFunction(functionSpecs FunctionSpecList)
	AddFunctionArgs(functionSpec FunctionSpec)
	AddFunctionCaller(funcName string, callerSpec CallerSpec)
	AddArgsToCallExpr(callerSpec CallerSpec)
	AddFunctionArgsToReturn(functionReturnArgs FunctionReturnArgsSpec)
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

func (a *AbstractCodeImpl) AddInterfaces(interfaceSpecs InterfaceSpecList) {
	for _, interfaceSpec := range interfaceSpecs {
		isAddBeforeStruct := interfaceSpec.StructName != ""
		decls := []ast.Decl{}
		structIndex := 0
		for i, decl := range a.file.Decls {
			gendecl, ok := decl.(*ast.GenDecl)
			if ok {
				if isAddBeforeStruct {
					// if the StructName is not empty
					// then we must add the interface before the struct
					// example what I expected
					// =======================
					// type Http interface {}
					// type HttpImpl struct {}
					// =======================
					// so the struct must defined after the interface

					// get the index of Decls
					for _, spec := range gendecl.Specs {
						typeSpec, ok := spec.(*ast.TypeSpec)
						if ok {
							_, ok := typeSpec.Type.(*ast.StructType)
							if ok {
								if typeSpec.Name.Name == interfaceSpec.StructName {
									structIndex = i
									break
								}
							}
						}
					}
				} else if gendecl.Tok == token.IMPORT {
					structIndex = i + 1
					break
				}
			}
		}

		decls = append(a.file.Decls[:structIndex+1], a.file.Decls[structIndex:]...)
		decls[structIndex] = &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: &ast.Ident{
						Name: interfaceSpec.Name,
					},
					Type: &ast.InterfaceType{
						Interface:  token.NoPos,
						Methods:    &ast.FieldList{},
						Incomplete: false,
					},
				},
			},
		}
		a.file.Decls = decls
	}
}

func (a *AbstractCodeImpl) AddInterfaceFuncDecl() error {
	return nil
}

func (a *AbstractCodeImpl) AddStructs(structSpecs StructSpecList) {
	for _, structSpec := range structSpecs {
		isAddAfterInterface := structSpec.InterfaceName != ""
		decls := []ast.Decl{}
		structIndex := 0
		for i, decl := range a.file.Decls {
			gendecl, ok := decl.(*ast.GenDecl)
			if ok {
				if isAddAfterInterface {
					// if the InterfaceName is not empty
					// then we must add the struct after the interface
					// example what I expected
					// =======================
					// type Http interface {}
					// type HttpImpl struct {}
					// =======================
					// so the struct must defined after the interface

					// get the index of Decls
					for _, spec := range gendecl.Specs {
						typeSpec, ok := spec.(*ast.TypeSpec)
						if ok {
							_, ok := typeSpec.Type.(*ast.InterfaceType)
							if ok {
								if typeSpec.Name.Name == structSpec.InterfaceName {
									structIndex = i + 1
									break
								}
							}
						}
					}
				} else if gendecl.Tok == token.IMPORT {
					structIndex = i + 1
					break
				}
			}
		}

		decls = append(a.file.Decls[:structIndex+1], a.file.Decls[structIndex:]...)
		decls[structIndex] = &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: &ast.Ident{
						Name: structSpec.Name,
					},
					Type: &ast.StructType{
						Fields:     &ast.FieldList{},
						Incomplete: false,
					},
				},
			},
		}
		a.file.Decls = decls
	}
}

func (a *AbstractCodeImpl) AddStructVarDecl(structArgs StructArgList) {
	for _, structArg := range structArgs {
		field := &ast.Field{}
		for _, decl := range a.file.Decls {
			gendecl, ok := decl.(*ast.GenDecl)
			if ok {
				if gendecl.Tok != token.TYPE {
					continue
				}
				for _, spec := range gendecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if ok {
						structType, ok := typeSpec.Type.(*ast.StructType)
						if ok {
							if typeSpec.Name.Name == structArg.StructName {
								field = &ast.Field{}

								if structArg.Name != "" {
									field.Names = []*ast.Ident{
										{
											Name: structArg.Name,
										},
									}
								}

								if structArg.IsPointer {
									if structArg.DataType.TypeName != "" {
										field.Type = &ast.StarExpr{
											X: &ast.SelectorExpr{
												X:   ast.NewIdent(structArg.DataType.LibName),
												Sel: ast.NewIdent(structArg.DataType.TypeName),
											},
										}
									} else {
										field.Type = &ast.StarExpr{
											X: ast.NewIdent(structArg.DataType.LibName),
										}
									}
								} else {
									if structArg.DataType.TypeName != "" {
										field.Type = &ast.SelectorExpr{
											X:   ast.NewIdent(structArg.DataType.LibName),
											Sel: ast.NewIdent(structArg.DataType.TypeName),
										}
									} else {
										field.Type = &ast.Ident{
											Name: structArg.DataType.LibName,
										}
									}
								}

								structType.Fields.List = append(structType.Fields.List, field)
								break
							}
						}
					}
				}
			}
		}
	}
}

func (a *AbstractCodeImpl) AddFunction(functionSpecs FunctionSpecList) {
	for _, functionSpec := range functionSpecs {
		decls := []ast.Decl{}
		structIndex := 0
		for i, decl := range a.file.Decls {
			gendecl, ok := decl.(*ast.GenDecl)
			if ok {
				if functionSpec.StructSpec != nil {
					if functionSpec.StructSpec.IsConstruct {
						// it define the function should placed after the struct
						// example, we have the object like this
						// type Http interface{}
						// type HttpImpl struct{}
						// then the function should be written after the struct
						// type Http interface{}
						// type HttpImpl struct{}
						// func NewHttp() {}
						for _, spec := range gendecl.Specs {
							typeSpec, ok := spec.(*ast.TypeSpec)
							if ok {
								_, ok := typeSpec.Type.(*ast.StructType)
								if ok {
									if functionSpec.StructSpec != nil {
										if typeSpec.Name.Name == functionSpec.StructSpec.DataTypes {
											structIndex = i + 1
											break
										}
									}
								}
							}
						}
					}
					structIndex = len(a.file.Decls) - 1
					break
				} else if gendecl.Tok == token.IMPORT {
					structIndex = i + 1
					break
				}
			}
		}

		newFunc := &ast.FuncDecl{
			Name: ast.NewIdent(functionSpec.Name),
			Type: &ast.FuncType{
				Params:  &ast.FieldList{},
				Results: &ast.FieldList{},
			},
			Body: &ast.BlockStmt{},
		}

		if functionSpec.StructSpec != nil {
			if functionSpec.StructSpec.Name != "" {
				if functionSpec.StructSpec.IsPointer {
					newFunc.Recv = &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent(functionSpec.StructSpec.Name),
								},
								Type: &ast.StarExpr{X: ast.NewIdent(functionSpec.StructSpec.DataTypes)},
							},
						},
					}
				} else {
					newFunc.Recv = &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent(functionSpec.StructSpec.Name),
								},
								Type: ast.NewIdent(functionSpec.StructSpec.DataTypes),
							},
						},
					}
				}
			}
		}

		for _, paramStmt := range functionSpec.Args {
			if paramStmt.IsPointer {
				if paramStmt.LibName != "" {
					newFunc.Type.Params.List = append(newFunc.Type.Params.List, &ast.Field{
						Names: []*ast.Ident{
							ast.NewIdent(paramStmt.Name),
						},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   ast.NewIdent(paramStmt.LibName),
								Sel: ast.NewIdent(paramStmt.DataType),
							},
						},
					})
				} else {
					newFunc.Type.Params.List = append(newFunc.Type.Params.List, &ast.Field{
						Names: []*ast.Ident{
							ast.NewIdent(paramStmt.Name),
						},
						Type: &ast.StarExpr{
							X: ast.NewIdent(paramStmt.DataType),
						},
					})
				}
			} else {
				if paramStmt.LibName != "" {
					newFunc.Type.Params.List = append(newFunc.Type.Params.List, &ast.Field{
						Names: []*ast.Ident{
							ast.NewIdent(paramStmt.Name),
						},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent(paramStmt.LibName),
							Sel: ast.NewIdent(paramStmt.DataType),
						},
					})
				} else {
					newFunc.Type.Params.List = append(newFunc.Type.Params.List, &ast.Field{
						Names: []*ast.Ident{
							ast.NewIdent(paramStmt.Name),
						},
						Type: ast.NewIdent(paramStmt.DataType),
					})
				}
			}
		}

		if functionSpec.Returns != nil {
			astReturnStmt := ast.ReturnStmt{}
			for _, returnStmt := range *functionSpec.Returns {
				// if the returnStmt is a pointer
				if returnStmt.IsPointer {
					if returnStmt.LibName != "" {
						// return http.HttpImpl but using pointer
						newFunc.Type.Results.List = append(newFunc.Type.Results.List, &ast.Field{
							Type: &ast.StarExpr{
								X: &ast.SelectorExpr{
									X:   ast.NewIdent(returnStmt.LibName),
									Sel: ast.NewIdent(returnStmt.DataType),
								},
							},
						})

						astReturnStmt.Results = append(astReturnStmt.Results, &ast.UnaryExpr{
							Op: token.AND,
							X: &ast.CompositeLit{
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent(returnStmt.LibName),
									Sel: ast.NewIdent(returnStmt.Return),
								},
							},
						})
					} else {
						// return only single statement without data type or anything
						// return HttpImpl but using pointer
						newFunc.Type.Results.List = append(newFunc.Type.Results.List, &ast.Field{
							Type: &ast.StarExpr{
								X: ast.NewIdent(returnStmt.DataType),
							},
						})

						astReturnStmt.Results = append(astReturnStmt.Results, &ast.UnaryExpr{
							Op: token.AND,
							X: &ast.CompositeLit{
								Type: ast.NewIdent(returnStmt.Return),
							},
						})
					}
				} else {
					if returnStmt.LibName != "" {
						// return http.HttpImpl
						newFunc.Type.Results.List = append(newFunc.Type.Results.List, &ast.Field{
							Type: &ast.SelectorExpr{
								X:   ast.NewIdent(returnStmt.LibName),
								Sel: ast.NewIdent(returnStmt.DataType),
							},
						})

						astReturnStmt.Results = append(astReturnStmt.Results, &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X:   ast.NewIdent(returnStmt.LibName),
								Sel: ast.NewIdent(returnStmt.Return),
							},
						})
					} else {
						// return only single statement without data type or anything
						// return HttpImpl
						newFunc.Type.Results.List = append(newFunc.Type.Results.List, &ast.Field{
							Type: ast.NewIdent(returnStmt.DataType),
						})

						if returnStmt.IsStruct {
							astReturnStmt.Results = append(astReturnStmt.Results, &ast.CompositeLit{
								Type: ast.NewIdent(returnStmt.Return),
							})
						} else {
							astReturnStmt.Results = append(astReturnStmt.Results, ast.NewIdent(returnStmt.Return))
						}
					}
				}
			}
			newFunc.Body.List = append(newFunc.Body.List, &astReturnStmt)
		}

		if len(a.file.Decls) == 0 {
			decls = append(a.file.Decls, newFunc)
		} else {
			decls = append(a.file.Decls[:structIndex+1], a.file.Decls[structIndex:]...)
			decls[structIndex] = newFunc
		}
		a.file.Decls = decls
	}
}

func (a *AbstractCodeImpl) AddFunctionArgs(functionSpec FunctionSpec) {
	for _, decl := range a.file.Decls {
		funcdecl, ok := decl.(*ast.FuncDecl)
		if ok {
			if funcdecl.Name.Name == functionSpec.Name {
				for _, paramStmt := range functionSpec.Args {
					if paramStmt.IsPointer {
						if paramStmt.LibName != "" {
							funcdecl.Type.Params.List = append(funcdecl.Type.Params.List, &ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent(paramStmt.Name),
								},
								Type: &ast.StarExpr{
									X: &ast.SelectorExpr{
										X:   ast.NewIdent(paramStmt.LibName),
										Sel: ast.NewIdent(paramStmt.DataType),
									},
								},
							})
						} else {
							funcdecl.Type.Params.List = append(funcdecl.Type.Params.List, &ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent(paramStmt.Name),
								},
								Type: &ast.StarExpr{
									X: ast.NewIdent(paramStmt.DataType),
								},
							})
						}
					} else {
						if paramStmt.LibName != "" {
							funcdecl.Type.Params.List = append(funcdecl.Type.Params.List, &ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent(paramStmt.Name),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent(paramStmt.LibName),
									Sel: ast.NewIdent(paramStmt.DataType),
								},
							})
						} else {
							funcdecl.Type.Params.List = append(funcdecl.Type.Params.List, &ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent(paramStmt.Name),
								},
								Type: ast.NewIdent(paramStmt.DataType),
							})
						}
					}
				}
				// after appending the function args then we must break it
				break
			}
		}
	}
}

// AddWireDependencyInjection
// add this thing like this
// var productRepo = wire.NewSet(
// 	productrepo.NewRepository,
// 	wire.Bind(
// 		new(productrepo.Repositories),
// 		new(*productrepo.RepositoriesImpl),
// 	),
// )
// func InitHttpProtocol() *http.HttpImpl {
// 	wire.Build(
// 		productRepo, <-- and append to here
// 	)
// 	return &http.HttpImpl{}
// }
// TODO: should make the function more globaly
func (a *AbstractCodeImpl) AddWireDependencyInjection(wireDependency WireDependencyInjection) {
	a.file.Decls = append(a.file.Decls[:len(a.file.Decls)], a.file.Decls...)

	if wireDependency.InterfaceLib == "" && wireDependency.InterfaceName == "" &&
		wireDependency.StructLib == "" && wireDependency.StructName == "" {
		a.file.Decls[0] = &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						{
							Name: wireDependency.VarName,
							Obj: &ast.Object{
								Kind: ast.ObjKind(token.VAR),
								Name: wireDependency.VarName,
							},
						},
					},
					Values: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("wire"),
								Sel: ast.NewIdent("NewSet"),
							},
							Args: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent(wireDependency.TargetInjectName),
									Sel: ast.NewIdent(wireDependency.TargetInjectConstructName),
								},
							},
						},
					},
				},
			},
		}
	} else {
		a.file.Decls[0] = &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						{
							Name: wireDependency.VarName,
							Obj: &ast.Object{
								Kind: ast.ObjKind(token.VAR),
								Name: wireDependency.VarName,
							},
						},
					},
					Values: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("wire"),
								Sel: ast.NewIdent("NewSet"),
							},
							Args: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent(wireDependency.TargetInjectName),
									Sel: ast.NewIdent(wireDependency.TargetInjectConstructName),
								},
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("wire"),
										Sel: ast.NewIdent("Bind"),
									},
									Args: []ast.Expr{
										&ast.CallExpr{
											Fun: ast.NewIdent("new"),
											Args: []ast.Expr{
												&ast.SelectorExpr{
													X:   ast.NewIdent(wireDependency.InterfaceLib),
													Sel: ast.NewIdent(wireDependency.InterfaceName),
												},
											},
										},
										&ast.CallExpr{
											Fun: ast.NewIdent("new"),
											Args: []ast.Expr{
												&ast.StarExpr{
													X: &ast.SelectorExpr{
														X:   ast.NewIdent(wireDependency.StructLib),
														Sel: ast.NewIdent(wireDependency.StructName),
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
	}
}

func (a *AbstractCodeImpl) AddFunctionCaller(funcName string, callerSpec CallerSpec) {
	ast.Inspect(a.file, func(n ast.Node) bool {
		funcDecl, ok := n.(*ast.FuncDecl)
		if ok {
			if funcDecl.Name.Name == funcName {
				funcDecl.Body.List = append(funcDecl.Body.List, funcDecl.Body.List...)

				funcDecl.Body.List[0] = &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent(callerSpec.Func.Name),
							Sel: ast.NewIdent(callerSpec.Func.Selector),
						},
					},
				}
			}
		}
		return true
	})
}

func (a *AbstractCodeImpl) AddFuncWireBuild(funcName string) {
	for _, decl := range a.file.Decls {
		funcdecls, ok := decl.(*ast.FuncDecl)
		if ok {
			if funcdecls.Name.Name == funcName {
				funcdecls.Body.List = append(funcdecls.Body.List[:len(funcdecls.Body.List)], funcdecls.Body.List...)
				funcdecls.Body.List[0] = &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("wire"),
							Sel: ast.NewIdent("Build"),
						},
					},
				}
				break
			}
		}
	}
}

func (a *AbstractCodeImpl) AddArgsToCallExpr(callerSpec CallerSpec) {
	ast.Inspect(a.file, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if ok {
			selectorStmt, ok := callExpr.Fun.(*ast.SelectorExpr)
			if ok {
				if selectorStmt.X.(*ast.Ident).Name == callerSpec.Func.Name && selectorStmt.Sel.Name == callerSpec.Func.Selector {
					for _, arg := range callerSpec.Args {
						if arg.SelectorStmt != nil {
							callExpr.Args = append(callExpr.Args,
								&ast.SelectorExpr{
									X:   ast.NewIdent(arg.SelectorStmt.LibName),
									Sel: ast.NewIdent(arg.SelectorStmt.DataType),
								},
							)
						} else if arg.BasicLit != nil {
							callExpr.Args = append(callExpr.Args,
								&ast.BasicLit{
									Kind:  arg.BasicLit.Kind,
									Value: arg.BasicLit.Value,
								},
							)
						} else if arg.Ident != nil {
							callExpr.Args = append(callExpr.Args,
								ast.NewIdent(arg.Ident.Name),
							)
						}
					}
				}
			}

		}
		return true
	})
}

func (a *AbstractCodeImpl) AddFunctionArgsToReturn(functionReturnArgs FunctionReturnArgsSpec) {
	for _, decl := range a.file.Decls {
		funcdecls, ok := decl.(*ast.FuncDecl)
		if ok {
			if funcdecls.Name.Name == functionReturnArgs.FuncName {
				for _, body := range funcdecls.Body.List {
					returnstmt, ok := body.(*ast.ReturnStmt)
					if ok {
						for _, result := range returnstmt.Results {
							unary, okUnary := result.(*ast.UnaryExpr)
							composite, okComposite := result.(*ast.CompositeLit)
							if okUnary {
								composite, okComposite = unary.X.(*ast.CompositeLit)
							}
							if okComposite {
								ident, ok := composite.Type.(*ast.Ident)
								if ok {
									if ident.Name == functionReturnArgs.ReturnName {
										composite.Elts = append(composite.Elts, &ast.KeyValueExpr{
											Key:   ast.NewIdent(functionReturnArgs.DataTypeKey),
											Value: ast.NewIdent(functionReturnArgs.DataTypeValue),
										})
										break
									}
								}
							}
						}
					}
				}
				break
			}
		}
	}
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
