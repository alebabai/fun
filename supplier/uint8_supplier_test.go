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
	resTestUint8Supplier uint8
	errTestUint8Supplier = errors.New("error")
)

func testUint8Supplier() (uint8, error) {
	return resTestUint8Supplier, nil
}

func testUint8SupplierWithError() (uint8, error) {
	return resTestUint8Supplier, errTestUint8Supplier
}

func TestUint8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Supplier
	}{
		{
			name: "ok",
			s:    testUint8Supplier,
		},
		{
			name: "with_error",
			s:    testUint8SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint8Supplier.Error())
			} else {
				r.Equal(resTestUint8Supplier, v)
			}
		})
	}
}

func TestUint8Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8Supplier,
		},
		{
			name: "with_error",
			s:    testUint8SupplierWithError,
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
				r.EqualError(err, errTestUint8Supplier.Error())
			} else {
				r.Equal(resTestUint8Supplier, v)
			}
		})
	}
}

func TestUint8Supplier_ToSilentUint8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8Supplier,
		},
		{
			name: "with_error",
			s:    testUint8SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint8Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8Supplier, v)
			}
		})
	}
}

func TestUint8Supplier_ToMustUint8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8Supplier,
		},
		{
			name: "with_error",
			s:    testUint8SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint8Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint8Supplier, v)
			}
		})
	}
}

func TestSilentUint8Supplier(t *testing.T) {
	var ss SilentUint8Supplier = func() uint8 {
		return resTestUint8Supplier
	}
	v := ss()
	require.Equal(t, resTestUint8Supplier, v)
}

func TestSilentUint8Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8Supplier,
		},
		{
			name: "with_error",
			s:    testUint8SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint8Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8Supplier, v)
			}
		})
	}
}

func TestMustUint8Supplier(t *testing.T) {
	var ms MustUint8Supplier = func() uint8 {
		return resTestUint8Supplier
	}

	v := ms()
	require.Equal(t, resTestUint8Supplier, v)
}

func TestMustUint8Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8Supplier,
		},
		{
			name: "with_error",
			s:    testUint8SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint8Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint8Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint8Supplier, v)
			}
		})
	}
}

func TestMustUint8Supplier_ToSilentUint8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8Supplier,
		},
		{
			name: "with_error",
			s:    testUint8SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint8Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8Supplier, v)
			}
		})
	}
}

func TestMustUint8Supplier_ToUint8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8Supplier
	}{
		{
			name: "ok",
			s:    testUint8Supplier,
		},
		{
			name: "with_error",
			s:    testUint8SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8Supplier()
			r.NotNil(ms)

			s := ms.ToUint8Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint8Supplier.Error())
			} else {
				r.Equal(resTestUint8Supplier, v)
			}
		})
	}
}
