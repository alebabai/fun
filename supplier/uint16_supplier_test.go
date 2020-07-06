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
	resTestUint16Supplier uint16
	errTestUint16Supplier = errors.New("error")
)

func testUint16Supplier() (uint16, error) {
	return resTestUint16Supplier, nil
}

func testUint16SupplierWithError() (uint16, error) {
	return resTestUint16Supplier, errTestUint16Supplier
}

func TestUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint16Supplier.Error())
			} else {
				r.Equal(resTestUint16Supplier, v)
			}
		})
	}
}

func TestUint16Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			s := tt.s.ToSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint16Supplier.Error())
			} else {
				r.Equal(resTestUint16Supplier, v)
			}
		})
	}
}

func TestUint16Supplier_ToSilentUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint16Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16Supplier, v)
			}
		})
	}
}

func TestUint16Supplier_ToMustUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint16Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint16Supplier, v)
			}
		})
	}
}

func TestSilentUint16Supplier(t *testing.T) {
	var ss SilentUint16Supplier = func() uint16 {
		return resTestUint16Supplier
	}
	v := ss()
	require.Equal(t, resTestUint16Supplier, v)
}

func TestSilentUint16Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint16Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16Supplier, v)
			}
		})
	}
}

func TestMustUint16Supplier(t *testing.T) {
	var ms MustUint16Supplier = func() uint16 {
		return resTestUint16Supplier
	}

	v := ms()
	require.Equal(t, resTestUint16Supplier, v)
}

func TestMustUint16Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint16Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint16Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint16Supplier, v)
			}
		})
	}
}

func TestMustUint16Supplier_ToSilentUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint16Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16Supplier, v)
			}
		})
	}
}

func TestMustUint16Supplier_ToUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16Supplier()
			r.NotNil(ms)

			s := ms.ToUint16Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint16Supplier.Error())
			} else {
				r.Equal(resTestUint16Supplier, v)
			}
		})
	}
}
