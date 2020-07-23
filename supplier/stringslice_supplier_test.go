// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier_test.go.tmpl
// DO NOT EDIT

package supplier

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	resTestStringSliceSupplier []string
	errTestStringSliceSupplier = errors.New("error")
)

func testStringSliceSupplier() ([]string, error) {
	return resTestStringSliceSupplier, nil
}

func testStringSliceSupplierWithError() ([]string, error) {
	return resTestStringSliceSupplier, errTestStringSliceSupplier
}

func TestStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestStringSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringSliceSupplier, v)
			}
		})
	}
}

func TestStringSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			s := tt.s.ToSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestStringSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringSliceSupplier, v)
			}
		})
	}
}

func TestStringSliceSupplier_ToSilentStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentStringSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringSliceSupplier, v)
			}
		})
	}
}

func TestStringSliceSupplier_ToMustStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestStringSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestStringSliceSupplier, v)
			}
		})
	}
}

func TestSilentStringSliceSupplier(t *testing.T) {
	var ss SilentStringSliceSupplier = func() []string {
		return resTestStringSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestStringSliceSupplier, v)
}

func TestSilentStringSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentStringSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringSliceSupplier, v)
			}
		})
	}
}

func TestMustStringSliceSupplier(t *testing.T) {
	var ms MustStringSliceSupplier = func() []string {
		return resTestStringSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestStringSliceSupplier, v)
}

func TestMustStringSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustStringSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestStringSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestStringSliceSupplier, v)
			}
		})
	}
}

func TestMustStringSliceSupplier_ToSilentStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentStringSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringSliceSupplier, v)
			}
		})
	}
}

func TestMustStringSliceSupplier_ToStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSliceSupplier()
			r.NotNil(ms)

			s := ms.ToStringSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestStringSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringSliceSupplier, v)
			}
		})
	}
}
