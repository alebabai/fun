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
	resTestFloat32PtrSupplier *float32
	errTestFloat32PtrSupplier = errors.New("error")
)

func testFloat32PtrSupplier() (*float32, error) {
	return resTestFloat32PtrSupplier, nil
}

func testFloat32PtrSupplierWithError() (*float32, error) {
	return resTestFloat32PtrSupplier, errTestFloat32PtrSupplier
}

func TestFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSupplier, v)
			}
		})
	}
}

func TestFloat32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSupplierWithError,
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
				r.EqualError(err, errTestFloat32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSupplier, v)
			}
		})
	}
}

func TestFloat32PtrSupplier_ToSilentFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSupplier, v)
			}
		})
	}
}

func TestFloat32PtrSupplier_ToMustFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat32PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat32PtrSupplier, v)
			}
		})
	}
}

func TestSilentFloat32PtrSupplier(t *testing.T) {
	var ss SilentFloat32PtrSupplier = func() *float32 {
		return resTestFloat32PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestFloat32PtrSupplier, v)
}

func TestSilentFloat32PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat32PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSupplier, v)
			}
		})
	}
}

func TestMustFloat32PtrSupplier(t *testing.T) {
	var ms MustFloat32PtrSupplier = func() *float32 {
		return resTestFloat32PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestFloat32PtrSupplier, v)
}

func TestMustFloat32PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat32PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat32PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat32PtrSupplier, v)
			}
		})
	}
}

func TestMustFloat32PtrSupplier_ToSilentFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSupplier, v)
			}
		})
	}
}

func TestMustFloat32PtrSupplier_ToFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSupplier()
			r.NotNil(ms)

			s := ms.ToFloat32PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSupplier, v)
			}
		})
	}
}
