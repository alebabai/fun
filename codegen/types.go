package main

import (
	"fmt"
	"strings"
)

var (
	BuiltinTypes = [...]string{
		"interface{}",
		"string",
		"rune",
		"bool",
		"byte",
		"uintptr",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"int", "int8", "int16", "int32", "int64",
		"float32", "float64",
	}
)

func GenerateComplexTypes(builtinType string) []string {
	return []string{
		fmt.Sprintf("*%s", builtinType),
		fmt.Sprintf("[]%s", builtinType),
		fmt.Sprintf("[]*%s", builtinType),
	}
}

func GetTypes() []*Type {
	types := make([]*Type, 0)
	for _, bt := range BuiltinTypes {
		var btTitle string
		if bt != "interface{}" {
			btTitle = strings.Title(bt)
			complexTypes := GenerateComplexTypes(bt)
			for _, ct := range complexTypes {
				ctTitle := btTitle
				if strings.HasPrefix(ct, "*") {
					ctTitle = fmt.Sprintf("%sPtr", btTitle)
				} else if strings.HasPrefix(ct, "[]*") {
					ctTitle = fmt.Sprintf("%sPtrSlice", btTitle)
				} else if strings.HasPrefix(ct, "[]") {
					ctTitle = fmt.Sprintf("%sSlice", btTitle)
				}
				t := &Type{
					Name:  ct,
					Title: ctTitle,
				}
				types = append(types, t)
			}
		}
		t := &Type{
			Name:  bt,
			Title: btTitle,
		}
		types = append(types, t)
	}
	return types
}
