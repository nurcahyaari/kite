package infrastructure

import (
	"github.com/nurcahyaari/kite/utils/fs"
)

type QueueOption struct {
	Options
	InfrastructurePath string
	DirName            string
	DirPath            string
	QueueType          string
}

func NewQueue(options QueueOption) AppGenerator {
	options.DirName = "queue"
	options.DirPath = fs.ConcatDirPath(options.InfrastructurePath, options.DirName)

	return options
}

func (o QueueOption) Run() error {

	// o.createQueueDir()

	return nil
}

// func (o QueueOption) createQueueDir() error {
// 	err := fs.CreateFolderIsNotExist(o.DirPath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
