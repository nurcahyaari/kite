//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/protocol/cli"
	"github.com/nurcahyaari/kite/src/domain/cachegen"
	"github.com/nurcahyaari/kite/src/domain/configgen"
	"github.com/nurcahyaari/kite/src/domain/dbgen"
	"github.com/nurcahyaari/kite/src/domain/dbgen/databasetype"
	"github.com/nurcahyaari/kite/src/domain/domaingen"
	"github.com/nurcahyaari/kite/src/domain/emptygen"
	"github.com/nurcahyaari/kite/src/domain/envgen"
	"github.com/nurcahyaari/kite/src/domain/generator"
	"github.com/nurcahyaari/kite/src/domain/handlergen"
	"github.com/nurcahyaari/kite/src/domain/infrastructuregen"
	"github.com/nurcahyaari/kite/src/domain/internalgen"
	"github.com/nurcahyaari/kite/src/domain/internalgen/loggergen"
	"github.com/nurcahyaari/kite/src/domain/internalgen/utilsgen"
	"github.com/nurcahyaari/kite/src/domain/miscgen"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
	"github.com/nurcahyaari/kite/src/domain/protocolgen/protocoltype"
	"github.com/nurcahyaari/kite/src/domain/srcgen"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
	clirouter "github.com/nurcahyaari/kite/src/protocol/cli"
	cliapp "github.com/urfave/cli/v2"
)

// filesystem database
var filesystemdb = wire.NewSet(
	database.NewFileSystem,
	wire.Bind(
		new(database.FileSystem),
		new(*database.FileSystemImpl),
	),
)

// domains
var configGen = wire.NewSet(
	configgen.NewConfig,
	wire.Bind(
		new(configgen.ConfigGen),
		new(*configgen.ConfigGenImpl),
	),
)

var envGen = wire.NewSet(
	envgen.NewEnvGen,
	wire.Bind(
		new(envgen.EnvGen),
		new(*envgen.EnvGenImpl),
	),
)

var emptyGen = wire.NewSet(
	emptygen.NewEmptyGen,
	wire.Bind(
		new(emptygen.EmptyGen),
		new(*emptygen.EmptyGenImpl),
	),
)

var wireGen = wire.NewSet(
	wiregen.NewWire,
	wire.Bind(
		new(wiregen.WireGen),
		new(*wiregen.WireGenImpl),
	),
)

var loggerGen = wire.NewSet(
	loggergen.NewLoggerGen,
	wire.Bind(
		new(loggergen.LoggerGen),
		new(*loggergen.LoggerGenImpl),
	),
)

var encryptionGen = wire.NewSet(
	utilsgen.NewUtilEncryption,
	wire.Bind(
		new(utilsgen.UtilEncryption),
		new(*utilsgen.UtilEncryptionImpl),
	),
)

var utilGen = wire.NewSet(
	utilsgen.NewUtil,
	wire.Bind(
		new(utilsgen.UtilGen),
		new(*utilsgen.UtilGenImpl),
	),
)

var internalGen = wire.NewSet(
	internalgen.NewInternal,
	wire.Bind(
		new(internalgen.InternalGen),
		new(*internalgen.InternalGenImpl),
	),
)

var dbGen = wire.NewSet(
	dbgen.NewDatabaseGen,
	wire.Bind(
		new(dbgen.DatabaseGen),
		new(*dbgen.DatabaseGenImpl),
	),
)

var mysqlGen = wire.NewSet(
	databasetype.NewMysqlGen,
	wire.Bind(
		new(databasetype.MysqlGen),
		new(*databasetype.MysqlGenImpl),
	),
)

var cacheGen = wire.NewSet(
	cachegen.NewCacheGen,
	wire.Bind(
		new(cachegen.CacheGen),
		new(*cachegen.CacheGenImpl),
	),
)

var infrastructureGen = wire.NewSet(
	infrastructuregen.NewInfrastructureGen,
	wire.Bind(
		new(infrastructuregen.InfrastructureGen),
		new(*infrastructuregen.InfrastructureGenImpl),
	),
)

var protocolType = wire.NewSet(
	protocoltype.NewProtocolType,
	wire.Bind(
		new(protocoltype.ProtocolType),
		new(*protocoltype.ProtocolTypeImpl),
	),
)

var protocolGen = wire.NewSet(
	protocolgen.NewProtocolGen,
	wire.Bind(
		new(protocolgen.ProtocolGen),
		new(*protocolgen.ProtocolGenImpl),
	),
)

var srcGen = wire.NewSet(
	srcgen.NewSrcGen,
	wire.Bind(
		new(srcgen.SrcGen),
		new(*srcgen.SrcGenImpl),
	),
)

var domainGen = wire.NewSet(
	domaingen.NewDomainGen,
	wire.Bind(
		new(domaingen.DomainGen),
		new(*domaingen.DomainGenImpl),
	),
)

var moduleGen = wire.NewSet(
	modulegen.NewModuleGen,
	wire.Bind(
		new(modulegen.ModuleGen),
		new(*modulegen.ModuleGenImpl),
	),
)

var gitignoreGen = wire.NewSet(
	miscgen.NewGitignoreGen,
	wire.Bind(
		new(miscgen.GitIgnoreGen),
		new(*miscgen.GitIgnoreGenImpl),
	),
)

var makefileGen = wire.NewSet(
	miscgen.NewMakefileGen,
	wire.Bind(
		new(miscgen.MakefileGen),
		new(*miscgen.MakefileGenImpl),
	),
)

var handlerGen = wire.NewSet(
	handlergen.NewHandlerGen,
	wire.Bind(
		new(handlergen.HandlerGen),
		new(*handlergen.HandlerGenImpl),
	),
)

var appGenerator = wire.NewSet(
	generator.NewAppGenerator,
	wire.Bind(
		new(generator.AppGenerator),
		new(*generator.AppGeneratorImpl),
	),
)

// type of protocol
var cliApp = wire.NewSet(cliapp.NewApp)

var cliRouter = wire.NewSet(
	clirouter.NewCliRouter,
	wire.Bind(
		new(clirouter.CliRouter),
		new(*clirouter.CliRouterImpl),
	),
)

func InitCliApp() *cli.CliImpl {
	wire.Build(
		filesystemdb,
		moduleGen,
		configGen,
		wireGen,
		envGen,
		emptyGen,
		loggerGen,
		encryptionGen,
		utilGen,
		cacheGen,
		mysqlGen,
		dbGen,
		infrastructureGen,
		internalGen,
		protocolGen,
		protocolType,
		srcGen,
		handlerGen,
		domainGen,
		appGenerator,
		gitignoreGen,
		makefileGen,
		cliApp,
		cliRouter,
		cli.NewCliApp,
	)
	return &cli.CliImpl{}
}
