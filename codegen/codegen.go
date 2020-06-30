package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

const tmplFileExtension = ".go.tmpl"

type Data struct {
	Type   *Type
	Source string
}

func generateCode(tmpl *template.Template, data *Data) ([]byte, error) {
	if tmpl == nil {
		return nil, fmt.Errorf("template is required")
	}

	codeBuff := new(bytes.Buffer)
	err := tmpl.Execute(codeBuff, data)
	if err != nil {
		return nil, fmt.Errorf("error executing template: %w", err)
	}

	fmtBytes, err := format.Source(codeBuff.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error formatting generated code: %w", err)
	}

	return fmtBytes, nil
}

func generateFilename(path string, t *Type) string {
	if t == nil {
		return path
	}

	var typePrefix string
	if t.Name != "interface{}" {
		lt := strings.ToLower(t.Title)
		typePrefix = fmt.Sprintf(
			"%s_",
			lt,
		)
	}

	return fmt.Sprintf(
		"%s/%s%s.go",
		filepath.Dir(path),
		typePrefix,
		strings.TrimSuffix(filepath.Base(path), tmplFileExtension),
	)
}

func Generate(dir string) {
	pattern := fmt.Sprintf("%s/*%s", dir, tmplFileExtension)
	paths, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	types := getTypes()
	for _, p := range paths {
		source := filepath.Base(p)
		tmpl, err := createTemplate(source, p)
		if err != nil {
			log.Fatal(err)
		}
		for _, t := range types {
			data := &Data{
				Type:   t,
				Source: source,
			}
			code, err := generateCode(tmpl, data)
			filename := generateFilename(p, t)
			err = ioutil.WriteFile(filename, code, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
