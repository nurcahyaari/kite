package generator

type ProjectInfo struct {
	GoModName   string
	ProjectPath string
	// Name can be use to indentify app name, domain name, module name, or etc
	Name         string
	ProtocolType string
}
