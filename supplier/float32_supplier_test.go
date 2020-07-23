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
	resTestFloat32Supplier float32
	errTestFloat32Supplier = errors.New("error")
)

func testFloat32Supplier() (float32, error) {
	return resTestFloat32Supplier, nil
}

func testFloat32SupplierWithError() (float32, error) {
	return resTestFloat32Supplier, errTestFloat32Supplier
}

func TestFloat32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name:    "with error",
			s:       testFloat32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat32Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32Supplier, v)
			}
		})
	}
}

func TestFloat32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name:    "with error",
			s:       testFloat32SupplierWithError,
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
				r.EqualError(err, errTestFloat32Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32Supplier, v)
			}
		})
	}
}

func TestFloat32Supplier_ToSilentFloat32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name:    "with error",
			s:       testFloat32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32Supplier, v)
			}
		})
	}
}

func TestFloat32Supplier_ToMustFloat32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name:    "with error",
			s:       testFloat32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32Supplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat32Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat32Supplier, v)
			}
		})
	}
}

func TestSilentFloat32Supplier(t *testing.T) {
	var ss SilentFloat32Supplier = func() float32 {
		return resTestFloat32Supplier
	}
	v := ss()
	require.Equal(t, resTestFloat32Supplier, v)
}

func TestSilentFloat32Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name:    "with error",
			s:       testFloat32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat32Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32Supplier, v)
			}
		})
	}
}

func TestMustFloat32Supplier(t *testing.T) {
	var ms MustFloat32Supplier = func() float32 {
		return resTestFloat32Supplier
	}

	v := ms()
	require.Equal(t, resTestFloat32Supplier, v)
}

func TestMustFloat32Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name:    "with error",
			s:       testFloat32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat32Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat32Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat32Supplier, v)
			}
		})
	}
}

func TestMustFloat32Supplier_ToSilentFloat32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name:    "with error",
			s:       testFloat32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32Supplier, v)
			}
		})
	}
}

func TestMustFloat32Supplier_ToFloat32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name:    "with error",
			s:       testFloat32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32Supplier()
			r.NotNil(ms)

			s := ms.ToFloat32Supplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat32Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat32Supplier, v)
			}
		})
	}
}
