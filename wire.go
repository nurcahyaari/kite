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
	"github.com/nurcahyaari/kite/src/domain/envgen"
	"github.com/nurcahyaari/kite/src/domain/generator"
	"github.com/nurcahyaari/kite/src/domain/infrastructuregen"
	"github.com/nurcahyaari/kite/src/domain/internalgen"
	"github.com/nurcahyaari/kite/src/domain/internalgen/loggergen"
	"github.com/nurcahyaari/kite/src/domain/internalgen/utilsgen"
	"github.com/nurcahyaari/kite/src/domain/misc"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
	"github.com/nurcahyaari/kite/src/domain/protocolgen/protocolhttpgen"
	"github.com/nurcahyaari/kite/src/domain/srcgen"
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

var wireGen = wire.NewSet(
	misc.NewWire,
	wire.Bind(
		new(misc.WireGen),
		new(*misc.WireGenImpl),
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

var protocolHttpGen = wire.NewSet(
	protocolhttpgen.NewProtocolHttp,
	wire.Bind(
		new(protocolhttpgen.ProtocolHttpGen),
		new(*protocolhttpgen.ProtocolHttpGenImpl),
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

var appGen = wire.NewSet(
	generator.NewApp,
	wire.Bind(
		new(generator.AppGenNew),
		new(*generator.AppGenNewImpl),
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
		configGen,
		wireGen,
		envGen,
		loggerGen,
		encryptionGen,
		utilGen,
		cacheGen,
		dbGen,
		infrastructureGen,
		internalGen,
		protocolGen,
		protocolHttpGen,
		srcGen,
		appGen,
		cliApp,
		cliRouter,
		cli.NewCliApp,
	)
	return &cli.CliImpl{}
}
