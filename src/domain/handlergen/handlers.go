package handlergen

import "github.com/nurcahyaari/kite/infrastructure/database"

type HandlerGen interface{}

type HandlerGenImpl struct {
	fs database.FileSystem
}

func NewHandlerGen(
	fs database.FileSystem,
) *HandlerGenImpl {
	return &HandlerGenImpl{
		fs: fs,
	}
}
