package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	DefaultTypes = []string{
		"interface{}", "int", "string", "bool",
	}
)

const HeaderTemplate = `// CODE GENERATED AUTOMATICALLY
// SOURCE: {{.Source}}
// DO NOT EDIT
`

type Type struct {
	Name  string
	Title string
}

type Data struct {
	Type   *Type
	Source string
}

func GetTypes(names []string) []*Type {
	types := make([]*Type, 0)
	for _, name := range names {
		var title string
		if name != "interface{}" {
			title = strings.Title(name)
		}
		t := &Type{
			Name:  name,
			Title: title,
		}
		types = append(types, t)
	}
	return types
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := flag.String(
		"dir",
		wd,
		"specify directory to process",
	)
	flag.Parse()

	pattern := fmt.Sprintf("%s/*.go.tmpl", *dir)
	paths, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	types := GetTypes(DefaultTypes)
	for _, p := range paths {
		fileBytes, err := ioutil.ReadFile(p)
		if err != nil {
			log.Fatal(err)
		}
		tmplBuff := bytes.NewBuffer(nil)
		tmplBuff.WriteString(HeaderTemplate)
		tmplBuff.WriteString("\n")
		tmplBuff.Write(fileBytes)

		source := filepath.Base(p)
		funcMap := template.FuncMap{
			"ToLower": strings.ToLower,
		}
		tmpl, err := template.
			New(source).
			Funcs(funcMap).
			Parse(tmplBuff.String())
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range types {
			resBuff := bytes.NewBuffer(nil)
			data := &Data{
				Type:   t,
				Source: source,
			}
			err := tmpl.Execute(resBuff, data)
			if err != nil {
				log.Fatal(err)
			}

			typePrefix := fmt.Sprintf(
				"%s_",
				strings.ToLower(t.Title),
			)
			if t.Name == "interface{}" {
				typePrefix = ""
			}

			filename := fmt.Sprintf(
				"%s/%s%s.go",
				*dir,
				typePrefix,
				strings.TrimSuffix(filepath.Base(p), ".go.tmpl"),
			)

			formattedBytes, err := format.Source(resBuff.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			err = ioutil.WriteFile(filename, formattedBytes, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
