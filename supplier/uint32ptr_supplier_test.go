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
	resTestUint32PtrSupplier *uint32
	errTestUint32PtrSupplier = errors.New("error")
)

func testUint32PtrSupplier() (*uint32, error) {
	return resTestUint32PtrSupplier, nil
}

func testUint32PtrSupplierWithError() (*uint32, error) {
	return resTestUint32PtrSupplier, errTestUint32PtrSupplier
}

func TestUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUint32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSupplier, v)
			}
		})
	}
}

func TestUint32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint32PtrSupplierWithError,
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
				r.EqualError(err, errTestUint32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSupplier, v)
			}
		})
	}
}

func TestUint32PtrSupplier_ToSilentUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSupplier, v)
			}
		})
	}
}

func TestUint32PtrSupplier_ToMustUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint32PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint32PtrSupplier, v)
			}
		})
	}
}

func TestSilentUint32PtrSupplier(t *testing.T) {
	var ss SilentUint32PtrSupplier = func() *uint32 {
		return resTestUint32PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestUint32PtrSupplier, v)
}

func TestSilentUint32PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint32PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSupplier, v)
			}
		})
	}
}

func TestMustUint32PtrSupplier(t *testing.T) {
	var ms MustUint32PtrSupplier = func() *uint32 {
		return resTestUint32PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestUint32PtrSupplier, v)
}

func TestMustUint32PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint32PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint32PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint32PtrSupplier, v)
			}
		})
	}
}

func TestMustUint32PtrSupplier_ToSilentUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSupplier, v)
			}
		})
	}
}

func TestMustUint32PtrSupplier_ToUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testUint32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSupplier()
			r.NotNil(ms)

			s := ms.ToUint32PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUint32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSupplier, v)
			}
		})
	}
}
