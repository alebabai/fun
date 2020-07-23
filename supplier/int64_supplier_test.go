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
	resTestInt64Supplier int64
	errTestInt64Supplier = errors.New("error")
)

func testInt64Supplier() (int64, error) {
	return resTestInt64Supplier, nil
}

func testInt64SupplierWithError() (int64, error) {
	return resTestInt64Supplier, errTestInt64Supplier
}

func TestInt64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name:    "with error",
			s:       testInt64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt64Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64Supplier, v)
			}
		})
	}
}

func TestInt64Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name:    "with error",
			s:       testInt64SupplierWithError,
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
				r.EqualError(err, errTestInt64Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64Supplier, v)
			}
		})
	}
}

func TestInt64Supplier_ToSilentInt64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name:    "with error",
			s:       testInt64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64Supplier, v)
			}
		})
	}
}

func TestInt64Supplier_ToMustInt64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name:    "with error",
			s:       testInt64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64Supplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt64Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt64Supplier, v)
			}
		})
	}
}

func TestSilentInt64Supplier(t *testing.T) {
	var ss SilentInt64Supplier = func() int64 {
		return resTestInt64Supplier
	}
	v := ss()
	require.Equal(t, resTestInt64Supplier, v)
}

func TestSilentInt64Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name:    "with error",
			s:       testInt64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt64Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64Supplier, v)
			}
		})
	}
}

func TestMustInt64Supplier(t *testing.T) {
	var ms MustInt64Supplier = func() int64 {
		return resTestInt64Supplier
	}

	v := ms()
	require.Equal(t, resTestInt64Supplier, v)
}

func TestMustInt64Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name:    "with error",
			s:       testInt64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt64Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt64Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt64Supplier, v)
			}
		})
	}
}

func TestMustInt64Supplier_ToSilentInt64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name:    "with error",
			s:       testInt64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64Supplier, v)
			}
		})
	}
}

func TestMustInt64Supplier_ToInt64Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name:    "with error",
			s:       testInt64SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64Supplier()
			r.NotNil(ms)

			s := ms.ToInt64Supplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt64Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64Supplier, v)
			}
		})
	}
}
