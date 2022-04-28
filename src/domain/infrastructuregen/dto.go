package infrastructuregen

import "github.com/nurcahyaari/kite/src/domain/dbgen"

type InfrastructureDto struct {
	GomodName          string
	InfrastructurePath string
	DatabaseType       dbgen.DbType
}
