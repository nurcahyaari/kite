package infrastructure

import (
	"fmt"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type DBOption struct {
	Options
	InfrastructurePath string
	DirName            string
	DirPath            string
}

func NewDB(options DBOption) AppGenerator {
	options.DirPath = options.InfrastructurePath

	return options
}

func (o DBOption) Run() error {
	err := o.createMysqlConnection()
	if err != nil {
		return err
	}

	return nil
}

func (o DBOption) createMysqlConnection() error {

	connectionTemplate := `
	dbHost := config.Get().DB.Mysql.Host
	dbPort := config.Get().DB.Mysql.Port
	dbName := config.Get().DB.Mysql.Name
	dbUser := config.Get().DB.Mysql.User
	dbPass := config.Get().DB.Mysql.Pass

	sHost := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	`
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "infrastructure",
		Template:    templates.DBTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "fmt",
			},
			{
				FilePath: fmt.Sprintf("%s/config", o.GoModName),
			},
			{
				FilePath: "github.com/go-sql-driver/mysql",
				Alias:    "_",
			},
			{
				FilePath: "github.com/jmoiron/sqlx",
			},
			{
				FilePath: "github.com/rs/zerolog/log",
			},
		},
		Data: map[string]interface{}{
			"AppName":            o.AppName,
			"ConnectionTemplate": connectionTemplate,
			"DBDialeg":           o.DefaultDBDialeg,
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(o.DirPath, fmt.Sprintf("%s.go", o.DefaultDBDialeg), templateString)

	return nil
}
