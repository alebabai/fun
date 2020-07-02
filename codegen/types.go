package main

import (
	"fmt"
	"strings"
)

const (
	ptrPrefix      = "*"
	slicePrefix    = "[]"
	ptrSlicePrefix = "[]*"
)

var (
	// BuiltinTypes represents golang builtin types
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

// Type holds type specific information
type Type struct {
	Name  string
	Title string
}

func generateComplexTypes(builtinType string) []string {
	if builtinType == "" {
		return make([]string, 0)
	}
	return []string{
		fmt.Sprintf("*%s", builtinType),
		fmt.Sprintf("[]%s", builtinType),
		fmt.Sprintf("[]*%s", builtinType),
	}
}

func getComplexTypeTitle(complexType string) string {
	title := strings.Title(complexType)
	if strings.HasPrefix(complexType, ptrPrefix) {
		title = fmt.Sprintf("%sPtr", strings.TrimLeft(title, ptrPrefix))
	} else if strings.HasPrefix(complexType, ptrSlicePrefix) {
		title = fmt.Sprintf("%sPtrSlice", strings.TrimLeft(title, ptrSlicePrefix))
	} else if strings.HasPrefix(complexType, slicePrefix) {
		title = fmt.Sprintf("%sSlice", strings.TrimLeft(title, slicePrefix))
	}
	return title
}

func getTypes(builtinTypes []string) []*Type {
	types := make([]*Type, 0)
	for _, bt := range builtinTypes {
		if bt == "" {
			continue
		}
		if bt != "interface{}" {
			t := &Type{
				Name:  bt,
				Title: strings.Title(bt),
			}
			types = append(types, t)

			for _, ct := range generateComplexTypes(bt) {
				t := &Type{
					Name:  ct,
					Title: getComplexTypeTitle(ct),
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
