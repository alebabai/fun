package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

const HeaderTemplate = `// CODE GENERATED AUTOMATICALLY
// SOURCE: {{.Source}}
// DO NOT EDIT
`

type HeaderData struct {
	Source string
}

type CodeData struct {
	Type   string
	GoType string
}

func main() {
	types := [...]string{
		"int", "string", "bool",
	}

	headerTmpl, err := template.New("header.go.tmpl").Parse(HeaderTemplate)
	if err != nil {
		log.Fatal(err)
	}

	paths, err := filepath.Glob("./[^_]**/[^_]*.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range paths {
		headerBuff := new(bytes.Buffer)
		headerData := &HeaderData{
			Source: filepath.Base(p),
		}
		err := headerTmpl.Execute(headerBuff, headerData)
		if err != nil {
			log.Fatal(err)
		}

		codeTmpl, err := template.ParseFiles(p)
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range types {
			codeBuff := new(bytes.Buffer)
			codeBuff.WriteString(headerBuff.String())

			codeData := &CodeData{
				Type:   strings.ToTitle(t),
				GoType: t,
			}
			err = codeTmpl.Execute(codeBuff, codeData)
			if err != nil {
				log.Fatal(err)
			}

			filename := fmt.Sprintf(
				"%s/%s_%s.go",
				filepath.Dir(p),
				strings.ToLower(t),
				strings.TrimSuffix(filepath.Base(p), ".go.tmpl"),
			)
			formattedBytes, err := format.Source(codeBuff.Bytes())
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
