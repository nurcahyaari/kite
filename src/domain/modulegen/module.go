package modulegen

// create module, what is module
// all of these package is based on module.
// dependency in here called as a module
type ModuleGen interface {
	CreateNewModule(dto ModuleDto) error
}

type ModuleGenImpl struct {
	path string
}

func NewModuleGen(path string) *ModuleGenImpl {
	return &ModuleGenImpl{
		path: path,
	}
}

func (s ModuleGenImpl) CreateNewModule(dto ModuleDto) error {
	return nil
}

// type ModuleGen interface {
// 	DtoGen
// 	EntityGen
// 	RepositoryGen
// 	ServiceGen
// 	HandlerGen
// 	// CreateSrcDir() error
// 	// CreateBaseModuleDir() error
// 	CreateNewModule() error
// 	// AppendModuleToWire() error
// }

// type ModuleGenImpl struct {
// 	// directory path is the path of the directory that store module
// 	DirectoryPath string
// 	// the base of module path
// 	BaseModulePath string
// 	// module name is the name of the module
// 	ModuleName string
// 	// module path is the place of the module
// 	ModulePath string
// 	// BaseHandlerPath  base of handler path
// 	BaseHandlerPath string
// 	// path of the project
// 	ProjectPath string
// 	// Derived module
// 	GomodName string
// 	*DtoGenImpl
// 	*EntityGenImpl
// 	*RepositoryGenImpl
// 	*ServiceGenImpl
// 	*HandlerGenImpl
// }

// func NewModuleGen(projectPath string, moduleName string, gomodName string) *ModuleGenImpl {
// 	directoryPath := fs.ConcatDirPath(
// 		projectPath, "src",
// 	)
// 	baseModulePath := fs.ConcatDirPath(directoryPath, "module")

// 	modulePath := fs.ConcatDirPath(
// 		baseModulePath,
// 		moduleName,
// 	)

// 	return &ModuleGenImpl{
// 		DirectoryPath:     directoryPath,
// 		BaseModulePath:    baseModulePath,
// 		ModulePath:        modulePath,
// 		ModuleName:        moduleName,
// 		ProjectPath:       projectPath,
// 		GomodName:         gomodName,
// 		DtoGenImpl:        NewDtoGen(modulePath),
// 		EntityGenImpl:     NewEntityGen(modulePath),
// 		RepositoryGenImpl: NewRepositoryGen(moduleName, modulePath, gomodName),
// 		ServiceGenImpl:    NewServiceGen(moduleName, modulePath, gomodName),
// 		HandlerGenImpl:    NewHandlerGen(directoryPath, moduleName, gomodName),
// 	}
// }

// // func (s *ModuleGenImpl) CreateSrcDir() error {
// // 	err := fs.CreateFolderIsNotExist(s.DirectoryPath)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }

// // func (s *ModuleGenImpl) CreateBaseModuleDir() error {
// // 	// validate is project exist
// // 	if !fs.IsFolderExist(s.ProjectPath) {
// // 		return fmt.Errorf("%s project path is not exist", s.ProjectPath)
// // 	}

// // 	moduleDir := fs.ConcatDirPath(s.DirectoryPath, "module")
// // 	fs.CreateFolderIsNotExist(moduleDir)

// // 	return nil
// // }

// func (s *ModuleGenImpl) CreateNewModule() error {
// 	err := fs.CreateFolderIsNotExist(s.ModulePath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // func (s *ModuleGenImpl) AppendModuleToWire() error {
// // 	// append repo
// // 	wireGen := misc.NewWire(s.ProjectPath, "")
// // 	wireGen.AddDependencyAfterCreatingModule(
// // 		ast.ImportSpec{
// // 			Name: fmt.Sprintf("%srepo", s.ModuleName),
// // 			Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "src/module", s.ModuleName, "repository")),
// // 		},
// // 		ast.WireDependencyInjection{
// // 			VarName:                   fmt.Sprintf("%sRepo", s.ModuleName),
// // 			TargetInjectName:          fmt.Sprintf("%srepo", s.ModuleName),
// // 			TargetInjectConstructName: "NewRepository",
// // 			InterfaceLib:              fmt.Sprintf("%srepo", s.ModuleName),
// // 			InterfaceName:             "Repository",
// // 			StructLib:                 fmt.Sprintf("%srepo", s.ModuleName),
// // 			StructName:                "RepositoryImpl",
// // 		},
// // 	)

// // 	// append service
// // 	wireGen.AddDependencyAfterCreatingModule(
// // 		ast.ImportSpec{
// // 			Name: fmt.Sprintf("%ssvc", s.ModuleName),
// // 			Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "src/module", s.ModuleName, "service")),
// // 		},
// // 		ast.WireDependencyInjection{
// // 			VarName:                   fmt.Sprintf("%sSvc", s.ModuleName),
// // 			TargetInjectName:          fmt.Sprintf("%ssvc", s.ModuleName),
// // 			TargetInjectConstructName: "NewService",
// // 			InterfaceLib:              fmt.Sprintf("%ssvc", s.ModuleName),
// // 			InterfaceName:             "Service",
// // 			StructLib:                 fmt.Sprintf("%ssvc", s.ModuleName),
// // 			StructName:                "ServiceImpl",
// // 		},
// // 	)

// // 	return nil
// // }
