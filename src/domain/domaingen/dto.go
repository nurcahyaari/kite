package domaingen

type DomainCreationalType int

const (
	DomainFolderOnlyCreation DomainCreationalType = iota + 1
	DomainFullCreation
)

const (
	TypeDomainFolderOnlyCreation = "DomainFolderOnlyCreation"
	TypeDomainFullCreation       = "DomainFullCreation"
)

func NewDomainCreationalType(t string) DomainCreationalType {
	var dt DomainCreationalType
	switch t {
	case TypeDomainFolderOnlyCreation:
		dt = DomainFolderOnlyCreation
	default:
		dt = DomainFullCreation
	}
	return dt
}

type DomainType int

const (
	SingleDomain DomainType = iota
	MultiDomain
)

const (
	TypeSingleDomain = "Singledomain"
	TypeMultiDomain  = "Multidomain"
)

func NewDomainType(t string) DomainType {
	var dt DomainType
	switch t {
	case TypeMultiDomain:
		dt = MultiDomain
	default:
		dt = SingleDomain
	}
	return dt
}

type DomainDto struct {
	Name                 string
	Path                 string
	ProjectPath          string
	GomodName            string
	DomainCreationalType DomainCreationalType
	DomainType           DomainType
}
