package templates

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"text/template"
)

//go:embed packagetemplate.tmpl
var packageTemplate string

type TemplateNew interface {
	AddTemplateFunction(name string, function interface{})
	Render() (string, error)
}

type TemplateNewImpl struct {
	PackageName             string
	TemplateLocation        string
	defaultTemplateFunction map[string]interface{}
}

func NewTemplateNewImpl(packageName, templateLoc string) *TemplateNewImpl {
	return &TemplateNewImpl{
		PackageName:             packageName,
		TemplateLocation:        templateLoc,
		defaultTemplateFunction: make(map[string]interface{}),
	}
}

func (s *TemplateNewImpl) AddTemplateFunction(name string, function interface{}) {
	s.defaultTemplateFunction[name] = function
}

func (s *TemplateNewImpl) Render(templateLoc string, data interface{}) (string, error) {
	var packageTmplBuffer bytes.Buffer
	var bodyTmplBuffer bytes.Buffer
	packageTemplateByte := []byte(packageTemplate)
	scanner := bufio.NewScanner(bytes.NewReader(packageTemplateByte))

	for scanner.Scan() {
		row := scanner.Text()
		packageTmplBuffer.WriteString(row)
		packageTmplBuffer.WriteString("\n")
	}
	if scanner.Err() != nil {
		return "", errors.New("error scan file")
	}

	packageTmpl := template.Must(template.New("template").
		Parse(packageTmplBuffer.String()))

	packageTmplData := map[string]interface{}{
		"PackageName": s.PackageName,
	}

	var packageTemplateBuf bytes.Buffer
	if err := packageTmpl.Execute(&packageTemplateBuf, packageTmplData); err != nil {
		return "", err
	}

	bodyByte := []byte(templateLoc)
	scanner = bufio.NewScanner(bytes.NewReader(bodyByte))

	for scanner.Scan() {
		row := scanner.Text()
		bodyTmplBuffer.WriteString(row)
		bodyTmplBuffer.WriteString("\n")
	}
	if scanner.Err() != nil {
		return "", errors.New("error scan file")
	}

	bodyTmpl := template.Must(template.New("template").
		Funcs(s.defaultTemplateFunction).
		Parse(bodyTmplBuffer.String()))

	var bodyTemplateBuf bytes.Buffer
	if err := bodyTmpl.Execute(&bodyTemplateBuf, data); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n%s",
		packageTemplateBuf.String(),
		bodyTemplateBuf.String(),
	), nil
}
