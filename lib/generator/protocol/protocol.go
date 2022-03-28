package protocol

import (
	"fmt"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type ProtocolGen interface {
	CreateInternalProtocolDir() error
	CreateProtocolTypeDir() error
	CreateProtocolDir() error
}

type ProtocolGenImpl struct {
	ProtocolType ProtocolType
	ProtocolPath string
	GomodName    string
}

// default protocol is http
func NewProtocolGen(
	protocol string,
	InternalPath string,
	GomodName string,
) *ProtocolGenImpl {
	protocolType := NewProtocolType(protocol)
	protocolPath := fs.ConcatDirPath(InternalPath, "protocol")
	return &ProtocolGenImpl{
		ProtocolType: protocolType,
		ProtocolPath: protocolPath,
		GomodName:    GomodName,
	}
}

func (s *ProtocolGenImpl) CreateProtocolTypeDir() error {
	dirPath := fs.ConcatDirPath(s.ProtocolPath, s.ProtocolType.ToString())
	err := fs.CreateFolderIsNotExist(dirPath)
	if err != nil {
		return err
	}

	switch s.ProtocolType {
	case Http:
		err = s.createProtocolInternalHttp(dirPath)
	}

	if err != nil {
		return err
	}

	return nil
}

// create protocol under src dir
func (s *ProtocolGenImpl) CreateProtocolDir() error {
	ProtocolTypePath := fs.ConcatDirPath(s.ProtocolPath, s.ProtocolType.ToString())
	err := fs.CreateFolderIsNotExist(ProtocolTypePath)
	if err != nil {
		return err
	}
	return nil
}

// create protocol under internal dir
func (s *ProtocolGenImpl) createProtocolInternalHttp(path string) error {
	var err error

	err = s.createInternalHttpErrorDir(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpMiddlewareDir(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpResponseDir(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpRouteDir(path)
	if err != nil {
		return err
	}

	err = s.createProtocolInternalHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) CreateInternalProtocolDir() error {
	err := fs.CreateFolderIsNotExist(s.ProtocolPath)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProtocolGenImpl) createInternalHttpErrorDir(path string) error {
	path = fs.ConcatDirPath(path, "errors")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = s.createInternalErrorHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) createInternalErrorHttpFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "errors",
		Template:    templates.ProtocolHttpChiErrorTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "fmt",
			},
			{
				FilePath: "net/http",
			},
			{
				FilePath: "regexp",
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(path, "error.go", templateString)

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpMiddlewareDir(path string) error {
	path = fs.ConcatDirPath(path, "middleware")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = s.createInternalMiddlewareHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) createInternalMiddlewareHttpFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "middleware",
		Template:    templates.ProtocolHttpChiMiddlewareTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "fmt",
			},
			{
				FilePath: fs.ConcatDirPath(s.GomodName, "config"),
			},
			{
				FilePath: "net/http",
			},
			{
				FilePath: "strings",
			},
			{
				FilePath: "time",
			},
			{
				Alias:    "httpresponse",
				FilePath: fs.ConcatDirPath(s.GomodName, "internal/protocol/http/response"),
			},
			{
				FilePath: fs.ConcatDirPath(s.GomodName, "internal/utils/encryption"),
			},
			{
				FilePath: "github.com/golang-jwt/jwt",
			},
			{
				FilePath: "github.com/golang-jwt/jwt/request",
			},
			{
				FilePath: "github.com/rs/zerolog/log",
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(path, "jwt.go", templateString)

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpResponseDir(path string) error {
	path = fs.ConcatDirPath(path, "response")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpResponseFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpResponseFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "response",
		Template:    templates.ProtocolHttpChiResponseTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "encoding/json",
			},
			{
				FilePath: fs.ConcatDirPath(s.GomodName, "internal/protocol/http/errors"),
			},
			{
				FilePath: "net/http",
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(path, "response.go", templateString)

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpRouteDir(path string) error {
	path = fs.ConcatDirPath(path, "router")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpRouteFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpRouteFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "router",
		Import: []templates.ImportedPackage{
			{
				FilePath: "github.com/go-chi/chi/v5",
			},
		},
		IsDependency: true,
		Dependency: templates.Dependency{
			HaveInterface:  false,
			DependencyName: "HttpRouter",
			DependencyMethod: []templates.DependencyMethod{
				{
					MethodImpl: "func (h *HttpRouterImpl) Router(r *chi.Mux) {}",
				},
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(path, "route.go", templateString)
	if err != nil {
		return nil
	}

	return nil
}

func (s *ProtocolGenImpl) createProtocolInternalHttpFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "http",
		Import: []templates.ImportedPackage{
			{
				FilePath: "fmt",
			},
			{
				FilePath: fs.ConcatDirPath(s.GomodName, "config"),
			},
			{
				FilePath: fs.ConcatDirPath(s.GomodName, "internal/protocol/http/router"),
			},
			{
				FilePath: "net/http",
			},
			{
				FilePath: "github.com/go-chi/chi/v5",
			},
		},
		IsDependency: true,
		Dependency: templates.Dependency{
			HaveInterface:  true,
			DependencyName: "Http",
			FuncParams: []templates.DependencyFuncParam{
				{
					ParamName:     "HttpRouter",
					ParamDataType: "*router.HttpRouterImpl",
				},
			},
			DependencyMethod: []templates.DependencyMethod{
				{
					MethodImpl: `func (p *HttpImpl) setupRouter(app *chi.Mux) {
						p.HttpRouter.Router(app)
					}`,
				},
				{
					Method: "Listen()",
					MethodImpl: `func (p *HttpImpl) Listen() {
						app := chi.NewRouter()

						p.setupRouter(app)

						http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Application.Port), app)
					}`,
				},
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(path, fmt.Sprintf("%s.go", "http"), templateString)
	if err != nil {
		return nil
	}

	return nil
}
