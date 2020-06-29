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

	types := GetTypes()
	for _, p := range paths {
		tmplBytes, err := ioutil.ReadFile(p)
		if err != nil {
			log.Fatal(err)
		}
		tsb := new(strings.Builder)
		tsb.WriteString(HeaderTemplate)
		tsb.WriteString("\n")
		tsb.Write(tmplBytes)

		source := filepath.Base(p)
		funcMap := template.FuncMap{
			"ToLower": strings.ToLower,
			"ToUpper": strings.ToUpper,
		}
		tmpl, err := template.
			New(source).
			Funcs(funcMap).
			Parse(tsb.String())
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range types {
			resBuff := new(bytes.Buffer)
			data := &Data{
				Type:   t,
				Source: source,
			}
			err := tmpl.Execute(resBuff, data)
			if err != nil {
				log.Fatal(err)
			}

			var typePrefix string
			if t.Name != "interface{}" {
				lt := strings.ToLower(t.Title)
				typePrefix = fmt.Sprintf(
					"%s_",
					lt,
				)
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
