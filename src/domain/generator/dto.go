package generator

import (
	"github.com/nurcahyaari/kite/internal/utils/errcustom"
)

type ProjectInfo struct {
	GoModName   string
	ProjectPath string
	// Name can be use to indentify app name, domain name, module name, or etc
	Name         string
	ProtocolType string
}

func (s ProjectInfo) Validate() error {
	var errList []string

	if s.GoModName == "" {
		errList = append(errList, "go.mod name cannot be empty")
	}

	if s.ProjectPath == "" {
		errList = append(errList, "project path name cannot be empty")
	}

	if s.Name == "" {
		errList = append(errList, "name cannot be empty")
	}

	err := errcustom.NewErrorResp()
	err.AddListToErrList(errList)

	if err.IsEmpty() {
		return err.ToErrorAsString()
	}

	return nil
}

type AppNewDto struct {
	ProjectInfo
}

func (s AppNewDto) Validate() error {
	return s.ProjectInfo.Validate()
}

type DomainNewDto struct {
	ProjectInfo
	IsCreateDomainFolderOnly bool
}

func (s DomainNewDto) Validate() error {
	err := s.ProjectInfo.Validate()

	errRespList := errcustom.NewErrRespFromError(err)
	if errRespList != nil {
		return errRespList.ToErrorAsString()
	}

	return nil
}

type HandlerNewDto struct {
	ProjectInfo
	ProtocolType string
}

func (s HandlerNewDto) Validate() error {
	err := s.ProjectInfo.Validate()

	errRespList := errcustom.NewErrRespFromError(err)
	if errRespList != nil {
		return errRespList.ToErrorAsString()
	}

	return nil
}

type ModuleNewDto struct {
	ProjectInfo
	PackageName string
}

func (s ModuleNewDto) Validate() error {
	err := s.ProjectInfo.Validate()

	errRespList := errcustom.NewErrRespFromError(err)
	if errRespList != nil {
		return errRespList.ToErrorAsString()
	}

	return nil
}
