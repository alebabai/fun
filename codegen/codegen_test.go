package main

import (
	"errors"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	BuiltinTypes = []string{
		"string",
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				dir: "./_test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			err := Generate(tt.args.dir)
			if tt.wantErr {
				r.Error(err)
			} else {
				r.NoError(err)
				r.FileExists("./_test/string_test.go")
				r.FileExists("./_test/stringptr_test.go")
				r.FileExists("./_test/stringslice_test.go")
				r.FileExists("./_test/stringptrslice_test.go")
			}
			t.Cleanup(func() {
				//clean up generated files
				err := os.Remove("./_test/string_test.go")
				r.NoError(err)
				err = os.Remove("./_test/stringptr_test.go")
				r.NoError(err)
				err = os.Remove("./_test/stringslice_test.go")
				r.NoError(err)
				err = os.Remove("./_test/stringptrslice_test.go")
				r.NoError(err)
			})
		})
	}
}

func Test_generateCode(t *testing.T) {
	rawTestTmpl := `
func test() {
	fmt.Println("Source: {{.Source}}")
	fmt.Println("Type name: {{.Type.Name}}")
	fmt.Println("Type title: {{.Type.Title}}")
}
`
	testTmpl, err := template.
		New("test").
		Parse(rawTestTmpl)
	require.NoError(t, err)
	require.NotNil(t, testTmpl)

	type args struct {
		tmpl *template.Template
		data *Data
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				tmpl: testTmpl,
				data: &Data{
					Type: &Type{
						Name:  "string",
						Title: "String",
					},
					Source: "supplier.go.tmpl",
				},
			},
			want: []byte(`
func test() {
	fmt.Println("Source: supplier.go.tmpl")
	fmt.Println("Type name: string")
	fmt.Println("Type title: String")
}
`),
		},
		{
			name:    "with error  empty template",
			wantErr: true,
		},
		{
			name: "with err  invalid data",
			args: args{
				tmpl: testTmpl,
				data: &Data{
					Type:   nil,
					Source: "",
				},
			},
			wantErr: true,
		},
		{
			name: "with error  code parsing failure",
			args: args{
				tmpl: testTmpl,
				data: &Data{
					Type: &Type{
						Name:  "\\",
						Title: "\\",
					},
					Source: "supplier.go.tmpl",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			fn, err := generateCode(tt.args.tmpl, tt.args.data)
			if tt.wantErr {
				r.Error(err)
			} else {
				r.Equal(tt.want, fn)
				r.NoError(err)
			}
		})
	}
}

func Test_generateFilename(t *testing.T) {
	type args struct {
		templatePath string
		t            *Type
	}
	tests := []struct {
		name string
		args args
		res  string
		err  error
	}{
		{
			name: "ok",
			args: args{
				templatePath: "supplier.go.tmpl",
				t: &Type{
					Name:  "string",
					Title: "String",
				},
			},
			res: "string_supplier.go",
		},
		{
			name: "ok  type  empty title",
			args: args{
				templatePath: "supplier.go.tmpl",
				t: &Type{
					Name:  "string",
					Title: "",
				},
			},
			res: "supplier.go",
		},
		{
			name: "ok  type  name equal interface{}",
			args: args{
				templatePath: "supplier.go.tmpl",
				t: &Type{
					Name:  "interface{}",
					Title: "",
				},
			},
			res: "supplier.go",
		},
		{
			name: "with error  invalid template filename  empty",
			res:  "",
			err:  errors.New("template filename is invalid"),
		},
		{
			name: "with error  invalid template filename  without_extension",
			args: args{
				templatePath: "supplier.go",
			},
			res: "",
			err: errors.New("template filename is invalid"),
		},
		{
			name: "with error  invalid template filename  only extension",
			args: args{
				templatePath: tmplFileExtension,
			},
			res: "",
			err: errors.New("template filename is invalid"),
		},
		{
			name: "with error  invalid type  nil",
			args: args{
				templatePath: "supplier.go.tmpl",
			},
			res: "",
			err: errors.New("type is invalid"),
		},
		{
			name: "with error  invalid type  empty_name",
			args: args{
				templatePath: "supplier.go.tmpl",
				t: &Type{
					Name:  "",
					Title: "",
				},
			},
			res: "",
			err: errors.New("type is invalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			fn, err := generateFilename(tt.args.templatePath, tt.args.t)
			if tt.err != nil {
				r.EqualError(err, tt.err.Error())
			} else {
				r.Equal(tt.res, fn)
				r.NoError(err)
			}
		})
	}
}
