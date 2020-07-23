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
	resTestInt8PtrSupplier *int8
	errTestInt8PtrSupplier = errors.New("error")
)

func testInt8PtrSupplier() (*int8, error) {
	return resTestInt8PtrSupplier, nil
}

func testInt8PtrSupplierWithError() (*int8, error) {
	return resTestInt8PtrSupplier, errTestInt8PtrSupplier
}

func TestInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt8PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt8PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSupplier, v)
			}
		})
	}
}

func TestInt8PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt8PtrSupplierWithError,
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
				r.EqualError(err, errTestInt8PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSupplier, v)
			}
		})
	}
}

func TestInt8PtrSupplier_ToSilentInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt8PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt8PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSupplier, v)
			}
		})
	}
}

func TestInt8PtrSupplier_ToMustInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt8PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt8PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt8PtrSupplier, v)
			}
		})
	}
}

func TestSilentInt8PtrSupplier(t *testing.T) {
	var ss SilentInt8PtrSupplier = func() *int8 {
		return resTestInt8PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestInt8PtrSupplier, v)
}

func TestSilentInt8PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt8PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt8PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSupplier, v)
			}
		})
	}
}

func TestMustInt8PtrSupplier(t *testing.T) {
	var ms MustInt8PtrSupplier = func() *int8 {
		return resTestInt8PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestInt8PtrSupplier, v)
}

func TestMustInt8PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt8PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt8PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt8PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt8PtrSupplier, v)
			}
		})
	}
}

func TestMustInt8PtrSupplier_ToSilentInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt8PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt8PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSupplier, v)
			}
		})
	}
}

func TestMustInt8PtrSupplier_ToInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt8PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSupplier()
			r.NotNil(ms)

			s := ms.ToInt8PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt8PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSupplier, v)
			}
		})
	}
}
