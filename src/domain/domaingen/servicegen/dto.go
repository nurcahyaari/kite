package servicegen

type ServiceDto struct {
	Path              string
	ProjectPath       string
	GomodName         string
	DomainName        string
	IsInjectRepo      bool
	IsInjectToHandler bool
}
