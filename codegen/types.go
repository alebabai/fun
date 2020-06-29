package main

import (
	"fmt"
	"strings"
)

var (
	BuiltinTypes = []string{
		"interface{}",
		"string",
		"rune",
		"bool",
		"byte",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"int", "int8", "int16", "int32", "int64",
		"float32", "float64",
	}
)

type Type struct {
	Name  string
	Title string
}

func generateComplexTypes(builtinType string) []string {
	return []string{
		fmt.Sprintf("*%s", builtinType),
		fmt.Sprintf("[]%s", builtinType),
		fmt.Sprintf("[]*%s", builtinType),
	}
}

func getComplexTypeTitle(builtinType, complexType string) string {
	title := strings.Title(builtinType)
	if strings.HasPrefix(complexType, "*") {
		title = fmt.Sprintf("%sPtr", title)
	} else if strings.HasPrefix(complexType, "[]*") {
		title = fmt.Sprintf("%sPtrSlice", title)
	} else if strings.HasPrefix(complexType, "[]") {
		title = fmt.Sprintf("%sSlice", title)
	}
	return title
}

func GetTypes() []*Type {
	types := make([]*Type, 0)
	for _, bt := range BuiltinTypes {
		if bt != "interface{}" {
			t := &Type{
				Name:  bt,
				Title: strings.Title(bt),
			}
			types = append(types, t)

			for _, ct := range generateComplexTypes(bt) {
				t := &Type{
					Name:  ct,
					Title: getComplexTypeTitle(bt, ct),
				}
				types = append(types, t)
			}
		} else {
			t := &Type{
				Name: bt,
			}
			types = append(types, t)
		}
	}
	return types
}
