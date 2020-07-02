package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
)

const tmplFileExtension = ".tmpl"

// Data represent all needed data for code generation
type Data struct {
	Type   *Type
	Source string
}

func generateCode(tmpl *template.Template, data *Data) ([]byte, error) {
	if tmpl == nil {
		return nil, errors.New("template is required")
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

func generateFilename(tmplFilename string, t *Type) (string, error) {
	if len(tmplFilename) <= len(tmplFileExtension) || !strings.HasSuffix(tmplFilename, tmplFileExtension) {
		return "", errors.New("template filename is invalid")
	}
	if t == nil || t.Name == "" {
		return "", errors.New("type is invalid")
	}

	var typePrefix string
	if t.Name != "interface{}" && t.Title != "" {
		lt := strings.ToLower(t.Title)
		typePrefix = fmt.Sprintf(
			"%s_",
			lt,
		)
	}

	fn := fmt.Sprintf(
		"%s%s",
		typePrefix,
		strings.TrimRight(tmplFilename, tmplFileExtension),
	)
	return fn, nil
}

// Generate encapsulates code generation flow
func Generate(dir string) error {
	pattern := fmt.Sprintf("%s/*%s", dir, tmplFileExtension)
	paths, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("error during glob pattern matching: %w", err)
	}

	types := getTypes(BuiltinTypes)
	for _, p := range paths {
		tmplFilename := filepath.Base(p)
		tmpl, err := createTemplate(tmplFilename, p)
		if err != nil {
			return fmt.Errorf("error during template creation: %w", err)
		}
		for _, t := range types {
			data := &Data{
				Type:   t,
				Source: tmplFilename,
			}
			code, err := generateCode(tmpl, data)
			if err != nil {
				return fmt.Errorf("error during code generation: %w", err)
			}

			fn, err := generateFilename(tmplFilename, t)
			if err != nil {
				return fmt.Errorf("error generating filename: %w", err)
			}

			fp := filepath.Join(dir, fn)
			err = ioutil.WriteFile(fp, code, 0644)
			if err != nil {
				return fmt.Errorf("error during file %s writing: %w", fn, err)
			}
		}
	}
	return nil
}
