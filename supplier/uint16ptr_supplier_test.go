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
	resTestUint16PtrSupplier *uint16
	errTestUint16PtrSupplier = errors.New("error")
)

func testUint16PtrSupplier() (*uint16, error) {
	return resTestUint16PtrSupplier, nil
}

func testUint16PtrSupplierWithError() (*uint16, error) {
	return resTestUint16PtrSupplier, errTestUint16PtrSupplier
}

func TestUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUint16PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSupplier, v)
			}
		})
	}
}

func TestUint16PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint16PtrSupplierWithError,
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
				r.EqualError(err, errTestUint16PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSupplier, v)
			}
		})
	}
}

func TestUint16PtrSupplier_ToSilentUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint16PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSupplier, v)
			}
		})
	}
}

func TestUint16PtrSupplier_ToMustUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint16PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint16PtrSupplier, v)
			}
		})
	}
}

func TestSilentUint16PtrSupplier(t *testing.T) {
	var ss SilentUint16PtrSupplier = func() *uint16 {
		return resTestUint16PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestUint16PtrSupplier, v)
}

func TestSilentUint16PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint16PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSupplier, v)
			}
		})
	}
}

func TestMustUint16PtrSupplier(t *testing.T) {
	var ms MustUint16PtrSupplier = func() *uint16 {
		return resTestUint16PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestUint16PtrSupplier, v)
}

func TestMustUint16PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint16PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint16PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint16PtrSupplier, v)
			}
		})
	}
}

func TestMustUint16PtrSupplier_ToSilentUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint16PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSupplier, v)
			}
		})
	}
}

func TestMustUint16PtrSupplier_ToUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint16PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint16PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSupplier()
			r.NotNil(ms)

			s := ms.ToUint16PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUint16PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSupplier, v)
			}
		})
	}
}
