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

type HeaderData struct {
	Source string
}

type CodeData struct {
	Type   string
	GoType string
}

func main() {
	types := []string{
		"Int", "String", "Bool",
	}

	ht, err := template.ParseFiles("./_codegen/header.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	paths, err := filepath.Glob("./[^_codegen]**/*.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range paths {
		hb := new(bytes.Buffer)
		hd := &HeaderData{
			Source: filepath.Base(p),
		}
		err := ht.Execute(hb, hd)
		if err != nil {
			log.Fatal(err)
		}

		ct, err := template.ParseFiles(p)
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range types {
			cb := new(bytes.Buffer)
			cb.WriteString(hb.String())

			cd := &CodeData{
				Type:   t,
				GoType: strings.ToLower(t),
			}
			err = ct.Execute(cb, cd)
			if err != nil {
				log.Fatal(err)
			}

			fn := fmt.Sprintf(
				"%s/%s_%s.go",
				filepath.Dir(p),
				strings.ToLower(t),
				strings.TrimSuffix(filepath.Base(p), ".go.tmpl"),
			)
			fcb, err := format.Source(cb.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			err = ioutil.WriteFile(fn, fcb, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
