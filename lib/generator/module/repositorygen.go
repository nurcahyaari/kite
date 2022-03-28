package module

import (
	"fmt"
	"strings"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type RepositoryGen interface {
	CreateRepositoryDir() error
	CreateRepositoryFile() error
}

type RepositoryGenImpl struct {
	RepositoryPath string
	ModuleName     string
	GomodName      string
}

func NewRepositoryGen(moduleName, modulePath, gomodName string) *RepositoryGenImpl {
	RepositoryPath := fs.ConcatDirPath(modulePath, "repository")
	return &RepositoryGenImpl{
		RepositoryPath: RepositoryPath,
		ModuleName:     moduleName,
		GomodName:      gomodName,
	}
}

func (s *RepositoryGenImpl) CreateRepositoryDir() error {
	err := fs.CreateFolderIsNotExist(s.RepositoryPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *RepositoryGenImpl) CreateRepositoryFile() error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "repository",
		Import: []templates.ImportedPackage{
			{
				FilePath: fs.ConcatDirPath(s.GomodName, "infrastructure"),
			},
		},
		IsDependency: true,
		Dependency: templates.Dependency{
			HaveInterface:  true,
			DependencyName: fmt.Sprintf("%sRepository", strings.Title(s.ModuleName)),
			FuncParams: []templates.DependencyFuncParam{
				{
					ParamName:     "db",
					ParamDataType: "*infrastructure.MysqlImpl",
				},
			},
			DependencyMethod: []templates.DependencyMethod{},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(s.RepositoryPath, "repository.go", templateString)
	if err != nil {
		return err
	}

	return nil
}
