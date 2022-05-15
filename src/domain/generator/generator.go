package generator

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/src/domain/configgen"
	"github.com/nurcahyaari/kite/src/domain/domaingen"
	"github.com/nurcahyaari/kite/src/domain/envgen"
	"github.com/nurcahyaari/kite/src/domain/handlergen"
	"github.com/nurcahyaari/kite/src/domain/infrastructuregen"
	"github.com/nurcahyaari/kite/src/domain/internalgen"
	"github.com/nurcahyaari/kite/src/domain/miscgen"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
	"github.com/nurcahyaari/kite/src/domain/srcgen"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
)

type AppGenerator interface {
	AppGenNew
	DomainGen
	HandlerGen
	ModuleGen
}

type AppGeneratorImpl struct {
	*AppGenNewImpl
	*DomainGenImpl
	*HandlerGenImpl
	*ModuleGenImpl
}

func NewAppGenerator(
	fs database.FileSystem,
	configGen configgen.ConfigGen,
	envGen envgen.EnvGen,
	wireGen wiregen.WireGen,
	internalGen internalgen.InternalGen,
	infrastructureGen infrastructuregen.InfrastructureGen,
	srcGen srcgen.SrcGen,
	domainGen domaingen.DomainGen,
	handlerGen handlergen.HandlerGen,
	gitignoreGen miscgen.GitIgnoreGen,
	makefileGen miscgen.MakefileGen,
	moduleGen modulegen.ModuleGen,
) *AppGeneratorImpl {
	return &AppGeneratorImpl{
		AppGenNewImpl: NewApp(
			fs,
			configGen,
			envGen,
			wireGen,
			internalGen,
			infrastructureGen,
			srcGen,
			domainGen,
			gitignoreGen,
			makefileGen,
		),
		DomainGenImpl:  NewDomainGen(fs, domainGen),
		HandlerGenImpl: NewHandlerGen(fs, handlerGen),
		ModuleGenImpl:  NewModuleGen(fs, moduleGen),
	}
}
