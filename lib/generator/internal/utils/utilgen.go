package utils

import (
	"fmt"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
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

	// then creating the file
	rsaTemplateData := map[string]interface{}{
		"ReadPublicKey":  templates.ReadPublicKeyTemplate,
		"ReadPrivateKey": templates.ReadPrivateKeyTemplate,
	}

	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "encryption",
		Template:    templates.RSABaseTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "crypto/rsa",
			},
			{
				FilePath: "crypto/x509",
			},
			{
				FilePath: "encoding/pem",
			},
			{
				FilePath: "errors",
			},
		},
		Data: rsaTemplateData,
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(path, "rsa.go", templateString)
}
