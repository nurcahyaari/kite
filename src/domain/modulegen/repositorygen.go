package modulegen

// type RepositoryGen interface {
// 	CreateRepositoryDir() error
// 	CreateRepositoryFile() error
// }

// type RepositoryGenImpl struct {
// 	RepositoryPath string
// 	ModuleName     string
// 	GomodName      string
// 	fs             database.FileSystem
// }

// func NewRepositoryGen(moduleName, modulePath, gomodName string) *RepositoryGenImpl {
// 	repositoryPath := utils.ConcatDirPath(modulePath, "repository")
// 	return &RepositoryGenImpl{
// 		RepositoryPath: repositoryPath,
// 		ModuleName:     moduleName,
// 		GomodName:      gomodName,
// 		fs:             database.NewFileSystem(repositoryPath),
// 	}
// }

// func (s *RepositoryGenImpl) CreateRepositoryDir() error {
// 	err := s.fs.CreateFolderIfNotExists()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *RepositoryGenImpl) CreateRepositoryFile() error {
// 	templateNew := templates.NewTemplateNewImpl("repository", "")
// 	templateCode, err := templateNew.Render("", nil)
// 	if err != nil {
// 		return err
// 	}

// 	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
// 	abstractCode.AddFunction(ast.FunctionSpecList{
// 		&ast.FunctionSpec{
// 			Name: "NewRepository",
// 			Args: ast.FunctionArgList{
// 				&ast.FunctionArg{
// 					IsPointer: true,
// 					Name:      "db",
// 					LibName:   "database",
// 					DataType:  "MysqlImpl",
// 				},
// 			},
// 			Returns: &ast.FunctionReturnSpecList{
// 				&ast.FunctionReturnSpec{
// 					IsPointer: true,
// 					IsStruct:  true,
// 					DataType:  "RepositoryImpl",
// 					Return:    "RepositoryImpl",
// 				},
// 			},
// 		},
// 	})
// 	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
// 		FuncName:      "NewRepository",
// 		ReturnName:    "RepositoryImpl",
// 		DataTypeKey:   "db",
// 		DataTypeValue: "db",
// 	})
// 	abstractCode.AddStructs(ast.StructSpecList{
// 		&ast.StructSpec{
// 			Name: "RepositoryImpl",
// 		},
// 	})
// 	abstractCode.AddStructVarDecl(ast.StructArgList{
// 		&ast.StructArg{
// 			StructName: "RepositoryImpl",
// 			IsPointer:  true,
// 			Name:       "db",
// 			DataType: ast.StructDtypes{
// 				LibName:  "database",
// 				TypeName: "MysqlImpl",
// 			},
// 		},
// 	})
// 	abstractCode.AddInterfaces(ast.InterfaceSpecList{
// 		&ast.InterfaceSpec{
// 			Name:       "Repository",
// 			StructName: "RepositoryImpl",
// 		},
// 	})
// 	abstractCode.AddImport(ast.ImportSpec{
// 		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(s.GomodName, "infrastructure", "database")),
// 	})
// 	err = abstractCode.RebuildCode()
// 	if err != nil {
// 		return err
// 	}
// 	templateBaseFileString := abstractCode.GetCode()

// 	err = s.fs.CreateFileIfNotExists("repository.go", templateBaseFileString)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
