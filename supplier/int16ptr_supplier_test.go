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
	resTestInt16PtrSupplier *int16
	errTestInt16PtrSupplier = errors.New("error")
)

func testInt16PtrSupplier() (*int16, error) {
	return resTestInt16PtrSupplier, nil
}

func testInt16PtrSupplierWithError() (*int16, error) {
	return resTestInt16PtrSupplier, errTestInt16PtrSupplier
}

func TestInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt16PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSupplier, v)
			}
		})
	}
}

func TestInt16PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSupplierWithError,
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
				r.EqualError(err, errTestInt16PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSupplier, v)
			}
		})
	}
}

func TestInt16PtrSupplier_ToSilentInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt16PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSupplier, v)
			}
		})
	}
}

func TestInt16PtrSupplier_ToMustInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt16PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt16PtrSupplier, v)
			}
		})
	}
}

func TestSilentInt16PtrSupplier(t *testing.T) {
	var ss SilentInt16PtrSupplier = func() *int16 {
		return resTestInt16PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestInt16PtrSupplier, v)
}

func TestSilentInt16PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt16PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSupplier, v)
			}
		})
	}
}

func TestMustInt16PtrSupplier(t *testing.T) {
	var ms MustInt16PtrSupplier = func() *int16 {
		return resTestInt16PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestInt16PtrSupplier, v)
}

func TestMustInt16PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt16PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt16PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt16PtrSupplier, v)
			}
		})
	}
}

func TestMustInt16PtrSupplier_ToSilentInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt16PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSupplier, v)
			}
		})
	}
}

func TestMustInt16PtrSupplier_ToInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSupplier()
			r.NotNil(ms)

			s := ms.ToInt16PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt16PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSupplier, v)
			}
		})
	}
}
