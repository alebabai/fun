package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
)

const headerTemplate = `// CODE GENERATED AUTOMATICALLY
// SOURCE: {{.Source}}
// DO NOT EDIT
`

func createTemplate(name, path string) (*template.Template, error) {
	tmplBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading template %s file: %w", name, err)
	}

	tsb := new(strings.Builder)
	tsb.WriteString(headerTemplate)
	tsb.WriteString("\n")
	tsb.Write(tmplBytes)

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
		"ToUpper": strings.ToUpper,
	}
	tmpl, err := template.
		New(name).
		Funcs(funcMap).
		Parse(tsb.String())
	if err != nil {
		return nil, fmt.Errorf("error during template parsing: %w", err)
	}

	return tmpl, nil
}
