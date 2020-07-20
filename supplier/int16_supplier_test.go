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
	resTestInt16Supplier int16
	errTestInt16Supplier = errors.New("error")
)

func testInt16Supplier() (int16, error) {
	return resTestInt16Supplier, nil
}

func testInt16SupplierWithError() (int16, error) {
	return resTestInt16Supplier, errTestInt16Supplier
}

func TestInt16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Supplier
	}{
		{
			name: "ok",
			s:    testInt16Supplier,
		},
		{
			name: "with_error",
			s:    testInt16SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt16Supplier.Error())
			} else {
				r.Equal(resTestInt16Supplier, v)
			}
		})
	}
}

func TestInt16Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16Supplier,
		},
		{
			name: "with_error",
			s:    testInt16SupplierWithError,
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
				r.EqualError(err, errTestInt16Supplier.Error())
			} else {
				r.Equal(resTestInt16Supplier, v)
			}
		})
	}
}

func TestInt16Supplier_ToSilentInt16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16Supplier,
		},
		{
			name: "with_error",
			s:    testInt16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt16Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16Supplier, v)
			}
		})
	}
}

func TestInt16Supplier_ToMustInt16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16Supplier,
		},
		{
			name: "with_error",
			s:    testInt16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt16Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt16Supplier, v)
			}
		})
	}
}

func TestSilentInt16Supplier(t *testing.T) {
	var ss SilentInt16Supplier = func() int16 {
		return resTestInt16Supplier
	}
	v := ss()
	require.Equal(t, resTestInt16Supplier, v)
}

func TestSilentInt16Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16Supplier,
		},
		{
			name: "with_error",
			s:    testInt16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt16Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16Supplier, v)
			}
		})
	}
}

func TestMustInt16Supplier(t *testing.T) {
	var ms MustInt16Supplier = func() int16 {
		return resTestInt16Supplier
	}

	v := ms()
	require.Equal(t, resTestInt16Supplier, v)
}

func TestMustInt16Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16Supplier,
		},
		{
			name: "with_error",
			s:    testInt16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt16Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt16Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt16Supplier, v)
			}
		})
	}
}

func TestMustInt16Supplier_ToSilentInt16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16Supplier,
		},
		{
			name: "with_error",
			s:    testInt16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt16Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16Supplier, v)
			}
		})
	}
}

func TestMustInt16Supplier_ToInt16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16Supplier
	}{
		{
			name: "ok",
			s:    testInt16Supplier,
		},
		{
			name: "with_error",
			s:    testInt16SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16Supplier()
			r.NotNil(ms)

			s := ms.ToInt16Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt16Supplier.Error())
			} else {
				r.Equal(resTestInt16Supplier, v)
			}
		})
	}
}