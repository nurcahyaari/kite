package dbgen

type DbType int

const (
	DbMysql DbType = iota + 1
)

const (
	MysqlCode string = "mysql"
)

func (s DbType) ToString() string {
	var dbType string
	switch s {
	case DbMysql:
		dbType = MysqlCode
	}
	return dbType
}

type DatabaseDto struct {
	DatabasePath string
	GomodName    string
	DatabaseType DbType
}
