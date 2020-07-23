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
	resTestFloat32PtrSliceSupplier []*float32
	errTestFloat32PtrSliceSupplier = errors.New("error")
)

func testFloat32PtrSliceSupplier() ([]*float32, error) {
	return resTestFloat32PtrSliceSupplier, nil
}

func testFloat32PtrSliceSupplierWithError() ([]*float32, error) {
	return resTestFloat32PtrSliceSupplier, errTestFloat32PtrSliceSupplier
}

func TestFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat32PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSliceSupplier, v)
			}
		})
	}
}

func TestFloat32PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestFloat32PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSliceSupplier, v)
			}
		})
	}
}

func TestFloat32PtrSliceSupplier_ToSilentFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSliceSupplier, v)
			}
		})
	}
}

func TestFloat32PtrSliceSupplier_ToMustFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat32PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat32PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentFloat32PtrSliceSupplier(t *testing.T) {
	var ss SilentFloat32PtrSliceSupplier = func() []*float32 {
		return resTestFloat32PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestFloat32PtrSliceSupplier, v)
}

func TestSilentFloat32PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat32PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustFloat32PtrSliceSupplier(t *testing.T) {
	var ms MustFloat32PtrSliceSupplier = func() []*float32 {
		return resTestFloat32PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestFloat32PtrSliceSupplier, v)
}

func TestMustFloat32PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat32PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat32PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat32PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustFloat32PtrSliceSupplier_ToSilentFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustFloat32PtrSliceSupplier_ToFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat32PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToFloat32PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat32PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32PtrSliceSupplier, v)
			}
		})
	}
}
