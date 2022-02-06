package templates

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"
)

type DependencyMethod struct {
	// Method
	Method string
	// Method with the implementation
	MethodImpl string
}

type DependencyFuncParam struct {
	ParamName     string
	ParamDataType string
}

type Dependency struct {
	HaveInterface    bool
	DependencyName   string
	FuncParams       []DependencyFuncParam
	DependencyMethod []DependencyMethod
}

type ImportedPackage struct {
	Alias    string
	FilePath string
}

type Template struct {
	PackageName  string
	Import       []ImportedPackage
	Template     TemplateBody
	Data         map[string]interface{}
	FilePath     string
	FuncMap      map[string]interface{}
	Dependency   Dependency
	IsDependency bool
	Header       string

	// private
	importPackage  bool
	importPackages bool
	headerPrefix   bool
}

func NewTemplate(
	template Template,
) Template {

	if len(template.Import) > 1 {
		template.importPackage = false
		template.importPackages = true
	} else if len(template.Import) == 1 {
		template.importPackage = true
		template.importPackages = false
	} else {
		template.importPackage = false
		template.importPackages = false
	}

	if template.Header != "" {
		template.headerPrefix = true
	}

	for i, imp := range template.Import {
		if imp.Alias != "" {
			template.Import[i].Alias = fmt.Sprintf("%s ", imp.Alias)
		}

		template.Import[i].FilePath = fmt.Sprintf(`"%s"`, imp.FilePath)
	}

	tmpl := Template{
		PackageName:    template.PackageName,
		Import:         template.Import,
		Template:       template.Template,
		importPackages: template.importPackages,
		importPackage:  template.importPackage,
		headerPrefix:   template.headerPrefix,
		Header:         fmt.Sprintf("%s\n", template.Header),
		Data:           template.Data,
		FilePath:       template.FilePath,
		IsDependency:   template.IsDependency,
		Dependency:     template.Dependency,
	}

	// defined the default function
	tmpl.FuncMap = map[string]interface{}{
		"Title":   strings.Title,
		"ToUpper": strings.ToUpper,
	}

	return tmpl
}

func (t Template) renderBodyTemplate() (string, error) {
	var templateBuffer bytes.Buffer

	templateByte := []byte(t.Template)
	var data interface{}
	data = t.Data
	if t.IsDependency {
		data = t.Dependency
		templateByte = []byte(DependencyTemplate)
	}

	scanner := bufio.NewScanner(bytes.NewReader(templateByte))

	for scanner.Scan() {
		row := scanner.Text()
		templateBuffer.WriteString(row)
		templateBuffer.WriteString("\n")
	}

	if scanner.Err() != nil {
		return "", errors.New("error scan file")
	}

	tmpl := template.Must(template.New("template").Funcs(t.FuncMap).Parse(templateBuffer.String()))

	var tmplString bytes.Buffer
	if err := tmpl.Execute(&tmplString, data); err != nil {
		return "", err
	}

	result := tmplString.String()

	return result, nil
}

func (t *Template) AddPackage(importPackage ImportedPackage) {
	t.Import = append(t.Import, importPackage)
}

func (t Template) Render() (string, error) {
	if t.PackageName != "" {
		var templateBuffer bytes.Buffer
		scanner := bufio.NewScanner(bytes.NewReader([]byte(PackageTemplate)))

		for scanner.Scan() {
			row := scanner.Text()
			templateBuffer.WriteString(row)
			templateBuffer.WriteString("\n")
		}

		if scanner.Err() != nil {
			return "", errors.New("error scan file")
		}

		tmpl := template.Must(template.New("template").Funcs(t.FuncMap).Parse(templateBuffer.String()))

		body, err := t.renderBodyTemplate()
		if err != nil {
			return "", err
		}

		data := map[string]interface{}{
			"PackageName":     t.PackageName,
			"importPackage":   t.importPackage,
			"importPackages":  t.importPackages,
			"ImportedPackage": t.Import,
			"headerPrefix":    t.headerPrefix,
			"Header":          t.Header,
			"Body":            body,
		}

		var tmplString bytes.Buffer
		if err := tmpl.Execute(&tmplString, data); err != nil {
			return "", err
		}

		result := tmplString.String()

		return result, nil
	} else {
		body, err := t.renderBodyTemplate()
		if err != nil {
			return "", err
		}
		return body, nil
	}
}
