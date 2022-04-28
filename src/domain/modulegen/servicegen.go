package modulegen

// type ServiceGen interface {
// 	CreateServiceDir() error
// 	CreateServiceFile() error
// }

// type ServiceGenImpl struct {
// 	ServicePath string
// 	ModuleName  string
// 	GomodName   string
// 	fs          *database.FileImpl
// }

// func NewServiceGen(moduleName, modulePath, gomodName string) *ServiceGenImpl {
// 	servicePath := utils.ConcatDirPath(modulePath, "service")
// 	return &ServiceGenImpl{
// 		ServicePath: servicePath,
// 		ModuleName:  moduleName,
// 		GomodName:   gomodName,
// 		fs:          database.NewFileSystem(servicePath),
// 	}
// }

// func (s *ServiceGenImpl) CreateServiceDir() error {
// 	err := s.fs.CreateFolderIfNotExists()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *ServiceGenImpl) CreateServiceFile() error {
// 	templateNew := templates.NewTemplateNewImpl("repository", "")
// 	templateCode, err := templateNew.Render("", nil)
// 	if err != nil {
// 		return err
// 	}

// 	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
// 	abstractCode.AddFunction(ast.FunctionSpecList{
// 		&ast.FunctionSpec{
// 			Name: "NewService",
// 			Args: ast.FunctionArgList{
// 				&ast.FunctionArg{
// 					Name:     fmt.Sprintf("%sRepo", s.ModuleName),
// 					LibName:  fmt.Sprintf("%srepo", s.ModuleName),
// 					DataType: "Repository",
// 				},
// 			},
// 			Returns: &ast.FunctionReturnSpecList{
// 				&ast.FunctionReturnSpec{
// 					IsPointer: true,
// 					IsStruct:  true,
// 					DataType:  "ServiceImpl",
// 					Return:    "ServiceImpl",
// 				},
// 			},
// 		},
// 	})
// 	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
// 		FuncName:      "NewService",
// 		ReturnName:    "ServiceImpl",
// 		DataTypeKey:   fmt.Sprintf("%sRepo", s.ModuleName),
// 		DataTypeValue: fmt.Sprintf("%sRepo", s.ModuleName),
// 	})
// 	abstractCode.AddStructs(ast.StructSpecList{
// 		&ast.StructSpec{
// 			Name: "ServiceImpl",
// 		},
// 	})
// 	abstractCode.AddStructVarDecl(ast.StructArgList{
// 		&ast.StructArg{
// 			StructName: "ServiceImpl",
// 			Name:       fmt.Sprintf("%sRepo", s.ModuleName),
// 			DataType: ast.StructDtypes{
// 				LibName:  fmt.Sprintf("%srepo", s.ModuleName),
// 				TypeName: "Repository",
// 			},
// 		},
// 	})
// 	abstractCode.AddInterfaces(ast.InterfaceSpecList{
// 		&ast.InterfaceSpec{
// 			Name:       "Service",
// 			StructName: "ServiceImpl",
// 		},
// 	})
// 	abstractCode.AddImport(ast.ImportSpec{
// 		Name: fmt.Sprintf("%srepo", s.ModuleName),
// 		Path: fmt.Sprintf("\"%s/src/module/%s/repository\"", s.GomodName, s.ModuleName),
// 	})
// 	err = abstractCode.RebuildCode()
// 	if err != nil {
// 		return err
// 	}
// 	templateBaseFileString := abstractCode.GetCode()

// 	err = s.fs.CreateFileIfNotExists("service.go", templateBaseFileString)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
