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
	resTestUint64Supplier uint64
	errTestUint64Supplier = errors.New("error")
)

func testUint64Supplier() (uint64, error) {
	return resTestUint64Supplier, nil
}

func testUint64SupplierWithError() (uint64, error) {
	return resTestUint64Supplier, errTestUint64Supplier
}

func TestUint64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name:    "with error",
			s:       testUint64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUint64Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64Supplier, v)
			}
		})
	}
}

func TestUint64Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name:    "with error",
			s:       testUint64SupplierWithError,
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
				r.EqualError(err, errTestUint64Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64Supplier, v)
			}
		})
	}
}

func TestUint64Supplier_ToSilentUint64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name:    "with error",
			s:       testUint64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64Supplier, v)
			}
		})
	}
}

func TestUint64Supplier_ToMustUint64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name:    "with error",
			s:       testUint64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64Supplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint64Supplier, v)
			}
		})
	}
}

func TestSilentUint64Supplier(t *testing.T) {
	var ss SilentUint64Supplier = func() uint64 {
		return resTestUint64Supplier
	}
	v := ss()
	require.Equal(t, resTestUint64Supplier, v)
}

func TestSilentUint64Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name:    "with error",
			s:       testUint64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint64Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64Supplier, v)
			}
		})
	}
}

func TestMustUint64Supplier(t *testing.T) {
	var ms MustUint64Supplier = func() uint64 {
		return resTestUint64Supplier
	}

	v := ms()
	require.Equal(t, resTestUint64Supplier, v)
}

func TestMustUint64Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name:    "with error",
			s:       testUint64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint64Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint64Supplier, v)
			}
		})
	}
}

func TestMustUint64Supplier_ToSilentUint64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name:    "with error",
			s:       testUint64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64Supplier, v)
			}
		})
	}
}

func TestMustUint64Supplier_ToUint64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name:    "with error",
			s:       testUint64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64Supplier()
			r.NotNil(ms)

			s := ms.ToUint64Supplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUint64Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64Supplier, v)
			}
		})
	}
}
