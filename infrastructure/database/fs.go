package database

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/nurcahyaari/kite/internal/logger"
	"github.com/nurcahyaari/kite/internal/utils"
)

type filesystemErr string

const (
	PathNotSet           filesystemErr = "Path is not set"
	FileNotFoundErr      filesystemErr = "File is not found"
	FileIsNotFileButADir filesystemErr = "File is not a file but it's a directory"
	FileIsNotDirButAFile filesystemErr = "File is not a directory but it's a file"
	FileWasCreated       filesystemErr = "File already created"
	FolderWasCreated     filesystemErr = "Folder already created %s"
	GoFmtErr             filesystemErr = "Failed to run go fmt"
	GitinitErr           filesystemErr = "Failed to run gitinit, maybe git isn't installed on your device"
	MkdirErr             filesystemErr = "Failed to create a folder"
	MkfileErr            filesystemErr = "Failed to create a file"
)

// panicErr return an error message
func PrintFsErr(errorName filesystemErr, opt ...interface{}) error {
	return fmt.Errorf(logger.Errorf(errorName), opt)
}

type FileSystem interface {
	IsFileExists(filepath string) bool
	IsFolderExists(path string) bool
	IsFolderEmpty(path string) bool
	CreateFolder(path string) error
	CreateFolderIfNotExists(path string) error
	CreateFile(path string, fileName string, fileTemplate string) error
	CreateFileIfNotExists(path string, fileName string, fileTemplate string) error
	DeleteFolder(path string) error
	DeleteFile(path string) error
	ReadFile(path string) (string, error)
	ReadFolderList(path string) ([]string, error)
	ReplaceFile(path string, fileName string, fileTemplate string) error
	CommandExec(path string, name string, args ...string) error
	writeStringToFile(fileTemplate string, file *os.File) error
}

type FileSystemImpl struct {
}

func NewFileSystem() *FileSystemImpl {
	return &FileSystemImpl{}
}

func (f FileSystemImpl) IsFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func (f FileSystemImpl) IsFolderExists(path string) bool {
	return f.IsFileExists(path)
}

func (f FileSystemImpl) IsFolderEmpty(path string) bool {
	fs, err := os.Open(path)
	if err != nil {
		return false
	}
	defer fs.Close()

	_, err = fs.Readdirnames(1)
	return err == io.EOF
}

func (f FileSystemImpl) CreateFolder(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return PrintFsErr(MkdirErr)
	}

	return nil
}

func (f FileSystemImpl) CreateFolderIfNotExists(path string) error {

	if f.IsFolderExists(path) {
		return PrintFsErr(FolderWasCreated, path)
	}

	return f.CreateFolder(path)
}

func (f FileSystemImpl) CreateFile(path string, fileName string, fileTemplate string) error {
	resultFile, err := os.Create(utils.ConcatDirPath(path, fileName))
	if err != nil {
		return PrintFsErr(MkfileErr)
	}

	err = f.writeStringToFile(fileTemplate, resultFile)
	if err != nil {
		return PrintFsErr(filesystemErr(err.Error()))
	}

	return nil
}

func (f FileSystemImpl) DeleteFolder(path string) error {

	return os.RemoveAll(fmt.Sprintf("%s/", path))
}

func (f FileSystemImpl) DeleteFile(path string) error {

	return os.Remove(path)
}

func (f FileSystemImpl) CreateFileIfNotExists(path string, fileName string, fileTemplate string) error {
	if f.IsFileExists(utils.ConcatDirPath(path, fileName)) {
		return PrintFsErr(FolderWasCreated, path)
	}

	err := f.CreateFile(path, fileName, fileTemplate)
	if err != nil {
		return err
	}

	return nil
}

func (f FileSystemImpl) ReadFile(path string) (string, error) {
	var fileValue string
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileValue += fmt.Sprintf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fileValue, nil
}

func (f FileSystemImpl) ReadFolderList(path string) ([]string, error) {
	var dirList []string
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return dirList, err
	}

	for _, file := range fileInfo {
		if file.IsDir() {
			dirList = append(dirList, file.Name())
		}
	}
	return dirList, nil
}

func (f FileSystemImpl) ReplaceFile(path string, fileName string, fileTemplate string) error {
	err := f.DeleteFile(utils.ConcatDirPath(path, fileName))
	if err != nil {
		return err
	}

	err = f.CreateFile(path, fileName, fileTemplate)
	if err != nil {
		return err
	}

	return nil
}

func (f FileSystemImpl) CommandExec(path string, name string, args ...string) error {

	os.Chdir(path)
	cmd := exec.Command(name, args...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (f FileSystemImpl) writeStringToFile(fileTemplate string, file *os.File) error {
	_, err := file.WriteString(fileTemplate)

	if err != nil {
		return err
	}

	return nil
}
