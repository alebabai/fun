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
	resTestInt32Supplier int32
	errTestInt32Supplier = errors.New("error")
)

func testInt32Supplier() (int32, error) {
	return resTestInt32Supplier, nil
}

func testInt32SupplierWithError() (int32, error) {
	return resTestInt32Supplier, errTestInt32Supplier
}

func TestInt32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Supplier
	}{
		{
			name: "ok",
			s:    testInt32Supplier,
		},
		{
			name: "with_error",
			s:    testInt32SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt32Supplier.Error())
			} else {
				r.Equal(resTestInt32Supplier, v)
			}
		})
	}
}

func TestInt32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32Supplier,
		},
		{
			name: "with_error",
			s:    testInt32SupplierWithError,
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
				r.EqualError(err, errTestInt32Supplier.Error())
			} else {
				r.Equal(resTestInt32Supplier, v)
			}
		})
	}
}

func TestInt32Supplier_ToSilentInt32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32Supplier,
		},
		{
			name: "with_error",
			s:    testInt32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32Supplier, v)
			}
		})
	}
}

func TestInt32Supplier_ToMustInt32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32Supplier,
		},
		{
			name: "with_error",
			s:    testInt32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt32Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt32Supplier, v)
			}
		})
	}
}

func TestSilentInt32Supplier(t *testing.T) {
	var ss SilentInt32Supplier = func() int32 {
		return resTestInt32Supplier
	}
	v := ss()
	require.Equal(t, resTestInt32Supplier, v)
}

func TestSilentInt32Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32Supplier,
		},
		{
			name: "with_error",
			s:    testInt32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt32Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32Supplier, v)
			}
		})
	}
}

func TestMustInt32Supplier(t *testing.T) {
	var ms MustInt32Supplier = func() int32 {
		return resTestInt32Supplier
	}

	v := ms()
	require.Equal(t, resTestInt32Supplier, v)
}

func TestMustInt32Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32Supplier,
		},
		{
			name: "with_error",
			s:    testInt32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt32Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt32Supplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt32Supplier, v)
			}
		})
	}
}

func TestMustInt32Supplier_ToSilentInt32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32Supplier,
		},
		{
			name: "with_error",
			s:    testInt32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32Supplier, v)
			}
		})
	}
}

func TestMustInt32Supplier_ToInt32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32Supplier
	}{
		{
			name: "ok",
			s:    testInt32Supplier,
		},
		{
			name: "with_error",
			s:    testInt32SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32Supplier()
			r.NotNil(ms)

			s := ms.ToInt32Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt32Supplier.Error())
			} else {
				r.Equal(resTestInt32Supplier, v)
			}
		})
	}
}
