package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_generateComplexTypes(t *testing.T) {
	type args struct {
		builtinType string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok",
			args: args{
				builtinType: "string",
			},
			want: []string{
				"*string",
				"[]string",
				"[]*string",
			},
		},
		{
			name: "builtin_type__empty",
			args: args{
				builtinType: "",
			},
			want: make([]string, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			got := generateComplexTypes(tt.args.builtinType)
			r.Equal(tt.want, got)
		})
	}
}

func Test_getComplexTypeTitle(t *testing.T) {
	type args struct {
		complexType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty_string",
			args: args{
				complexType: "",
			},
			want: "",
		},
		{
			name: "to_title",
			args: args{
				complexType: "string",
			},
			want: "String",
		},
		{
			name: "pointer",
			args: args{
				complexType: "*string",
			},
			want: "StringPtr",
		},
		{
			name: "slice",
			args: args{
				complexType: "[]string",
			},
			want: "StringSlice",
		},
		{
			name: "slice_of_pointers",
			args: args{
				complexType: "[]*string",
			},
			want: "StringPtrSlice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			got := getComplexTypeTitle(tt.args.complexType)
			r.Equal(tt.want, got)
		})
	}
}

func Test_getTypes(t *testing.T) {
	type args struct {
		builtinTypes []string
	}
	tests := []struct {
		name string
		args args
		want []*Type
	}{
		{
			name: "for_empty_interface_type",
			args: args{
				builtinTypes: []string{
					"interface{}",
				},
			},
			want: []*Type{
				{
					Name:  "interface{}",
					Title: "",
				},
			},
		},
		{
			name: "for_any_other_type",
			args: args{
				builtinTypes: []string{
					"string",
				},
			},
			want: []*Type{
				{
					Name:  "string",
					Title: "String",
				},
				{
					Name:  "*string",
					Title: "StringPtr",
				},
				{
					Name:  "[]string",
					Title: "StringSlice",
				},
				{
					Name:  "[]*string",
					Title: "StringPtrSlice",
				},
			},
		},
		{
			name: "empty slice",
			want: make([]*Type, 0),
		},
		{
			name: "empty type",
			args: args{
				builtinTypes: []string{
					"",
				},
			},
			want: make([]*Type, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			got := getTypes(tt.args.builtinTypes)
			r.Equal(tt.want, got)
		})
	}
}
