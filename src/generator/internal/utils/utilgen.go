package utils

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/src/ast"
	"github.com/nurcahyaari/kite/src/templates/internaltemplate/utilstemplate/encryptiontemplate"
	"github.com/nurcahyaari/kite/src/utils/fs"
)

type UtilGen interface {
	CreateUtilDir() error
	CreateRsaReader() error
}

type UtilGenImpl struct {
	UtilPath string
}

func NewUtil(
	internalPath string,
) *UtilGenImpl {
	utilPath := fs.ConcatDirPath(internalPath, "utils")
	return &UtilGenImpl{
		UtilPath: utilPath,
	}
}

func (s *UtilGenImpl) CreateUtilDir() error {
	err := fs.CreateFolderIsNotExist(s.UtilPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *UtilGenImpl) CreateRsaReader() error {
	// first creating the dir path
	path := fs.ValidatePath(s.UtilPath)
	path = fmt.Sprintf("%s%s", path, "encryption")
	// no need to check if error, because we still need to continue the proses
	fs.CreateFolderIsNotExist(path)

	templateNew := encryptiontemplate.NewRsaTemplate()
	rsaReaderTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	rsaReaderAbstractCode := ast.NewAbstractCode(rsaReaderTemplate, parser.ParseComments)
	rsaReaderAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"crypto/rsa\"",
	})
	rsaReaderAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"crypto/x509\"",
	})
	rsaReaderAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"encoding/pem\"",
	})
	rsaReaderAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"errors\"",
	})

	// after manipulate the code, rebuild the code
	err = rsaReaderAbstractCode.RebuildCode()
	if err != nil {
		return err
	}
	// get the manipulate code
	rsaRaderCode := rsaReaderAbstractCode.GetCode()

	return fs.CreateFileIfNotExist(path, "rsa.go", rsaRaderCode)
}
