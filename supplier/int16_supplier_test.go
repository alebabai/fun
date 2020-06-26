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
	testInt16SupplierResult int16
	testInt16SupplierError  = errors.New("error")
)

func testInt16Supplier() (int16, error) {
	return testInt16SupplierResult, nil
}

func testInt16SupplierWithError() (int16, error) {
	return testInt16SupplierResult, testInt16SupplierError
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
				r.EqualError(err, testInt16SupplierError.Error())
			} else {
				r.Equal(testInt16SupplierResult, v)
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
				r.EqualError(err, testInt16SupplierError.Error())
			} else {
				r.Equal(testInt16SupplierResult, v)
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
				r.Equal(testInt16SupplierResult, v)
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
				r.PanicsWithError(testInt16SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt16SupplierResult, v)
			}
		})
	}
}

func TestSilentInt16Supplier(t *testing.T) {
	var ss SilentInt16Supplier = func() int16 {
		return testInt16SupplierResult
	}
	v := ss()
	require.Equal(t, testInt16SupplierResult, v)
}

func TestMustInt16Supplier(t *testing.T) {
	var ms MustInt16Supplier = func() int16 {
		return testInt16SupplierResult
	}

	v := ms()
	require.Equal(t, testInt16SupplierResult, v)
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
				r.Equal(testInt16SupplierResult, v)
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
				r.EqualError(err, testInt16SupplierError.Error())
			} else {
				r.Equal(testInt16SupplierResult, v)
			}
		})
	}
}
