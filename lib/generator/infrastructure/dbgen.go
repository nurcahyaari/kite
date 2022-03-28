package infrastructure

import (
	"fmt"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type DatabaseGen interface {
	CreateMysqlConnection() error
}

type DatabaseGenImpl struct {
	AppName            string
	InfrastructurePath string
	GomodName          string
}

func NewDatabaseGen(dbGenImpl DatabaseGenImpl) *DatabaseGenImpl {
	return &dbGenImpl
}

func (s *DatabaseGenImpl) CreateMysqlConnection() error {
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
				FilePath: fmt.Sprintf("%s/config", s.GomodName),
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
			"AppName":            s.AppName,
			"ConnectionTemplate": connectionTemplate,
			"DBDialeg":           "mysql",
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.InfrastructurePath, "mysql.go", templateString)
}
