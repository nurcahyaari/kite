package modulegen

// type EntityGen interface {
// 	CreateEntityDir() error
// 	CreateEntityFile() error
// }

// type EntityGenImpl struct {
// 	EntityPath string
// 	fs         *database.FileImpl
// }

// func NewEntityGen(modulePath string) *EntityGenImpl {
// 	entityPath := utils.ConcatDirPath(modulePath, "entity")
// 	return &EntityGenImpl{
// 		EntityPath: entityPath,
// 		fs:         database.NewFileSystem(entityPath),
// 	}
// }

// func (s *EntityGenImpl) CreateEntityDir() error {
// 	err := s.fs.CreateFolderIfNotExists()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *EntityGenImpl) CreateEntityFile() error {
// 	return nil
// }
