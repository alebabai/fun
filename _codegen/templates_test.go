package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"text/template"
)

func Test_createTemplate(t *testing.T) {
	type args struct {
		name string
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *template.Template
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				name: "test",
				path: "./_test/test.go.tmpl",
			},
		},
		{
			name:    "with_error__empty_args",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			tmpl, err := createTemplate(tt.args.name, tt.args.path)
			if tt.wantErr {
				r.Nil(tmpl)
				r.Error(err)
			} else {
				r.NotNil(tmpl)
				r.NoError(err)
				r.Equal(tt.args.name, tmpl.Name())
			}
		})
	}
}
