package utilsgen

import (
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates/internaltemplate/utilstemplate/encryptiontemplate"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

type UtilEncryption interface {
	CreateRsaReader(dto EncryptionDto) error
}

type UtilEncryptionImpl struct {
	fs database.FileSystem
}

func NewUtilEncryption(
	fs database.FileSystem,
) *UtilEncryptionImpl {
	return &UtilEncryptionImpl{
		fs: fs,
	}
}

func (s *UtilEncryptionImpl) CreateRsaReader(dto EncryptionDto) error {
	s.fs.CreateFolderIfNotExists(dto.Path)

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

	return s.fs.CreateFileIfNotExists(dto.Path, "rsa.go", rsaRaderCode)
}
