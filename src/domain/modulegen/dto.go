package modulegen

type ModuleType int

func NewModuleTypeFromString(s string) ModuleType {
	var mT ModuleType
	switch s {
	case ServiceModuleString:
		mT = ServiceModule
	case RepositoryModuleString:
		mT = RepositoryModule
	case ProtocolModuleString:
		mT = ProtocolModule
	case DatabaseModuleString:
		mT = DatabaseModule
	default:
		mT = AnyModule
	}
	return mT
}

const (
	ServiceModule ModuleType = iota + 1
	RepositoryModule
	ProtocolModule
	DatabaseModule
	AnyModule
)

const (
	ServiceModuleString    string = "service"
	RepositoryModuleString string = "repository"
	ProtocolModuleString   string = "protocol"
	DatabaseModuleString   string = "database"
	AnyModuleString        string = "any"
)

func (m ModuleType) ToString() string {
	moduleType := ""
	switch m {
	case ServiceModule:
		moduleType = ServiceModuleString
	case RepositoryModule:
		moduleType = RepositoryModuleString
	case ProtocolModule:
		moduleType = ProtocolModuleString
	case DatabaseModule:
		moduleType = DatabaseModuleString
	case AnyModule:
		moduleType = AnyModuleString
	}

	return moduleType
}

type ModuleDto struct {
	PackageName string
	ModuleName  string
	Path        string
	GomodName   string
}
