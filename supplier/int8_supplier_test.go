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
	resTestInt8Supplier int8
	errTestInt8Supplier = errors.New("error")
)

func testInt8Supplier() (int8, error) {
	return resTestInt8Supplier, nil
}

func testInt8SupplierWithError() (int8, error) {
	return resTestInt8Supplier, errTestInt8Supplier
}

func TestInt8Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name:    "with error",
			s:       testInt8SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt8Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8Supplier, v)
			}
		})
	}
}

func TestInt8Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name:    "with error",
			s:       testInt8SupplierWithError,
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
				r.EqualError(err, errTestInt8Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8Supplier, v)
			}
		})
	}
}

func TestInt8Supplier_ToSilentInt8Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name:    "with error",
			s:       testInt8SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt8Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8Supplier, v)
			}
		})
	}
}

func TestInt8Supplier_ToMustInt8Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name:    "with error",
			s:       testInt8SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8Supplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt8Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt8Supplier, v)
			}
		})
	}
}

func TestSilentInt8Supplier(t *testing.T) {
	var ss SilentInt8Supplier = func() int8 {
		return resTestInt8Supplier
	}
	v := ss()
	require.Equal(t, resTestInt8Supplier, v)
}

func TestSilentInt8Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name:    "with error",
			s:       testInt8SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt8Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8Supplier, v)
			}
		})
	}
}

func TestMustInt8Supplier(t *testing.T) {
	var ms MustInt8Supplier = func() int8 {
		return resTestInt8Supplier
	}

	v := ms()
	require.Equal(t, resTestInt8Supplier, v)
}

func TestMustInt8Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name:    "with error",
			s:       testInt8SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt8Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt8Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt8Supplier, v)
			}
		})
	}
}

func TestMustInt8Supplier_ToSilentInt8Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name:    "with error",
			s:       testInt8SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt8Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8Supplier, v)
			}
		})
	}
}

func TestMustInt8Supplier_ToInt8Supplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name:    "with error",
			s:       testInt8SupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8Supplier()
			r.NotNil(ms)

			s := ms.ToInt8Supplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt8Supplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8Supplier, v)
			}
		})
	}
}
