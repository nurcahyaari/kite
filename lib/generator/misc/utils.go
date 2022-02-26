package misc

import (
	"fmt"

	"github.com/nurcahyaari/kite/lib/impl"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type UtilsOption struct {
	impl.GeneratorOptions
	InternalPath string
	DirName      string
	DirPath      string
}

func NewUtils(
	options UtilsOption,
) impl.AppGenerator {
	options.DirName = "utils"
	options.DirPath = fs.ConcatDirPath(options.InternalPath, options.DirName)

	return options
}

func (o UtilsOption) Run() error {

	o.createUtilsDir()

	return nil
}

func (o UtilsOption) createUtilsDir() error {
	err := fs.CreateFolderIsNotExist(o.DirPath)
	if err != nil {
		return err
	}

	o.createEncryptionFolder()

	return nil
}

func (o UtilsOption) createEncryptionFolder() error {
	path := fs.ValidatePath(o.DirPath)
	path = fmt.Sprintf("%s%s", path, "encryption")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	o.createRSAReaderFile(path)

	return nil
}

func (o UtilsOption) createRSAReaderFile(path string) error {
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

	fs.CreateFileIfNotExist(path, "rsa.go", templateString)

	return nil
}

func (o UtilsOption) createBaseFile() error {
	return nil
}
