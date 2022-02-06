package impl

type GeneratorOptions struct {
	AppName         string
	GoModName       string
	ProjectPath     string
	Path            string
	DefaultDBDialeg string
}

type AppGenerator interface {
	Run() error
}
