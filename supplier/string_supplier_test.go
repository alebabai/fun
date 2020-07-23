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
	resTestStringSupplier string
	errTestStringSupplier = errors.New("error")
)

func testStringSupplier() (string, error) {
	return resTestStringSupplier, nil
}

func testStringSupplierWithError() (string, error) {
	return resTestStringSupplier, errTestStringSupplier
}

func TestStringSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name:    "with error",
			s:       testStringSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestStringSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringSupplier, v)
			}
		})
	}
}

func TestStringSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name:    "with error",
			s:       testStringSupplierWithError,
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
				r.EqualError(err, errTestStringSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringSupplier, v)
			}
		})
	}
}

func TestStringSupplier_ToSilentStringSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name:    "with error",
			s:       testStringSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentStringSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringSupplier, v)
			}
		})
	}
}

func TestStringSupplier_ToMustStringSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name:    "with error",
			s:       testStringSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestStringSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestStringSupplier, v)
			}
		})
	}
}

func TestSilentStringSupplier(t *testing.T) {
	var ss SilentStringSupplier = func() string {
		return resTestStringSupplier
	}
	v := ss()
	require.Equal(t, resTestStringSupplier, v)
}

func TestSilentStringSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name:    "with error",
			s:       testStringSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentStringSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringSupplier, v)
			}
		})
	}
}

func TestMustStringSupplier(t *testing.T) {
	var ms MustStringSupplier = func() string {
		return resTestStringSupplier
	}

	v := ms()
	require.Equal(t, resTestStringSupplier, v)
}

func TestMustStringSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name:    "with error",
			s:       testStringSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustStringSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestStringSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestStringSupplier, v)
			}
		})
	}
}

func TestMustStringSupplier_ToSilentStringSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name:    "with error",
			s:       testStringSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentStringSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringSupplier, v)
			}
		})
	}
}

func TestMustStringSupplier_ToStringSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name:    "with error",
			s:       testStringSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSupplier()
			r.NotNil(ms)

			s := ms.ToStringSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestStringSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringSupplier, v)
			}
		})
	}
}
