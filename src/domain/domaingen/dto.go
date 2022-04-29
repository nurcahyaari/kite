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

type DomainDto struct {
	Name                 string
	Path                 string
	ProjectPath          string
	GomodName            string
	DomainCreationalType DomainCreationalType
}
