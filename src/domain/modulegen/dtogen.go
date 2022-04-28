package modulegen

// type DtoGen interface {
// 	CreateDtoDir() error
// 	CreateDtoFile() error
// }

// type DtoGenImpl struct {
// 	DtoPath string
// 	fs      *database.FileImpl
// }

// func NewDtoGen(modulePath string) *DtoGenImpl {
// 	dtoPath := utils.ConcatDirPath(modulePath, "dto")
// 	return &DtoGenImpl{
// 		DtoPath: dtoPath,
// 		fs:      database.NewFileSystem(dtoPath),
// 	}
// }

// func (s *DtoGenImpl) CreateDtoDir() error {
// 	err := s.fs.CreateFolderIfNotExists()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *DtoGenImpl) CreateDtoFile() error {
// 	return nil
// }
