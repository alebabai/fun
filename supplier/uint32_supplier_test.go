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
	resTestUint32Supplier uint32
	errTestUint32Supplier = errors.New("error")
)

func testUint32Supplier() (uint32, error) {
	return resTestUint32Supplier, nil
}

func testUint32SupplierWithError() (uint32, error) {
	return resTestUint32Supplier, errTestUint32Supplier
}

func TestUint32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name:    "with error",
			s:       testUint32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUint32Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32Supplier, v)
			}
		})
	}
}

func TestUint32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name:    "with error",
			s:       testUint32SupplierWithError,
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
				r.EqualError(err, errTestUint32Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32Supplier, v)
			}
		})
	}
}

func TestUint32Supplier_ToSilentUint32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name:    "with error",
			s:       testUint32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32Supplier, v)
			}
		})
	}
}

func TestUint32Supplier_ToMustUint32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name:    "with error",
			s:       testUint32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32Supplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint32Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint32Supplier, v)
			}
		})
	}
}

func TestSilentUint32Supplier(t *testing.T) {
	var ss SilentUint32Supplier = func() uint32 {
		return resTestUint32Supplier
	}
	v := ss()
	require.Equal(t, resTestUint32Supplier, v)
}

func TestSilentUint32Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name:    "with error",
			s:       testUint32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint32Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32Supplier, v)
			}
		})
	}
}

func TestMustUint32Supplier(t *testing.T) {
	var ms MustUint32Supplier = func() uint32 {
		return resTestUint32Supplier
	}

	v := ms()
	require.Equal(t, resTestUint32Supplier, v)
}

func TestMustUint32Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name:    "with error",
			s:       testUint32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint32Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint32Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint32Supplier, v)
			}
		})
	}
}

func TestMustUint32Supplier_ToSilentUint32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name:    "with error",
			s:       testUint32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32Supplier, v)
			}
		})
	}
}

func TestMustUint32Supplier_ToUint32Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name:    "with error",
			s:       testUint32SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32Supplier()
			r.NotNil(ms)

			s := ms.ToUint32Supplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUint32Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32Supplier, v)
			}
		})
	}
}
