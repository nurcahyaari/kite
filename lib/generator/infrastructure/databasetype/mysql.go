package databasetype

type MysqlGen interface {
	CreateMysqlConfig() error
}

type MysqlGenImpl struct{}

func (s *MysqlGenImpl) CreateMysqlConfig() error {
	return nil
}
