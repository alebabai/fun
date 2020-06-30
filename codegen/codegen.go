package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
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

func Generate(dir string) error {
	pattern := fmt.Sprintf("%s/*%s", dir, tmplFileExtension)
	paths, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("error during glob pattern matching: %w", err)
	}

	types := getTypes()
	for _, p := range paths {
		source := filepath.Base(p)
		tmpl, err := createTemplate(source, p)
		if err != nil {
			return fmt.Errorf("error during template creation: %w", err)
		}
		for _, t := range types {
			data := &Data{
				Type:   t,
				Source: source,
			}
			code, err := generateCode(tmpl, data)
			if err != nil {
				return fmt.Errorf("error during code generation: %w", err)
			}
			filename := generateFilename(p, t)
			err = ioutil.WriteFile(filename, code, 0644)
			if err != nil {
				return fmt.Errorf("error during file %s writing: %w", filename, err)
			}
		}
	}
	return nil
}
