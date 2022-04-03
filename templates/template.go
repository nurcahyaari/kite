package templates

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
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
	template := ""
	if s.PackageName != "" {
		body, err := s.renderPackageTemplate()
		if err != nil {
			return "", err
		}
		template += "\n" + body
	}
	if templateLoc != "" {
		body, err := s.renderBodyTemplate(templateLoc, data)
		if err != nil {
			return "", err
		}
		template += "\n" + body
	}

	return template, nil
}

func (s *TemplateNewImpl) renderPackageTemplate() (string, error) {
	var packageTmplBuffer bytes.Buffer
	var scanner *bufio.Scanner

	packageTemplateByte := []byte(packageTemplate)
	scanner = bufio.NewScanner(bytes.NewReader(packageTemplateByte))

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
	return packageTemplateBuf.String(), nil
}

func (s *TemplateNewImpl) renderBodyTemplate(templateLoc string, data interface{}) (string, error) {
	var bodyTmplBuffer bytes.Buffer
	var scanner *bufio.Scanner

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

	return bodyTemplateBuf.String(), nil
}
