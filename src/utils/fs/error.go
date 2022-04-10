package fs

import (
	"errors"
	"fmt"

	"github.com/nurcahyaari/kite/src/utils/logger"
)

type filesystemErr string

const (
	FileNotFoundErr      filesystemErr = "File is not found"
	FileIsNotFileButADir filesystemErr = "File is not a file but it's a directory"
	FileIsNotDirButAFile filesystemErr = "File is not a directory but it's a file"
	FileWasCreated       filesystemErr = "File already created"
	FolderWasCreated     filesystemErr = "Folder already created"
	GoFmtErr             filesystemErr = "Failed to run go fmt"
	GoModInitErr         filesystemErr = "Failed to go mod init"
	MkdirErr             filesystemErr = "Failed to create a folder"
	MkfileErr            filesystemErr = "Failed to create a file"
)

// panicErr return an error message
func printErr(errorName filesystemErr) error {
	return errors.New(logger.Errorf(fmt.Sprintf("%s", errorName)))
}
