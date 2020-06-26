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
	testInt8SupplierResult int8
	testInt8SupplierError  = errors.New("error")
)

func testInt8Supplier() (int8, error) {
	return testInt8SupplierResult, nil
}

func testInt8SupplierWithError() (int8, error) {
	return testInt8SupplierResult, testInt8SupplierError
}

func TestInt8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Supplier
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name: "with_error",
			s:    testInt8SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt8SupplierError.Error())
			} else {
				r.Equal(testInt8SupplierResult, v)
			}
		})
	}
}

func TestInt8Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name: "with_error",
			s:    testInt8SupplierWithError,
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
				r.EqualError(err, testInt8SupplierError.Error())
			} else {
				r.Equal(testInt8SupplierResult, v)
			}
		})
	}
}

func TestInt8Supplier_ToSilentInt8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name: "with_error",
			s:    testInt8SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt8Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt8SupplierResult, v)
			}
		})
	}
}

func TestInt8Supplier_ToMustInt8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name: "with_error",
			s:    testInt8SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt8SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt8SupplierResult, v)
			}
		})
	}
}

func TestSilentInt8Supplier(t *testing.T) {
	var ss SilentInt8Supplier = func() int8 {
		return testInt8SupplierResult
	}
	v := ss()
	require.Equal(t, testInt8SupplierResult, v)
}

func TestMustInt8Supplier(t *testing.T) {
	var ms MustInt8Supplier = func() int8 {
		return testInt8SupplierResult
	}

	v := ms()
	require.Equal(t, testInt8SupplierResult, v)
}

func TestMustInt8Supplier_ToSilentInt8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name: "with_error",
			s:    testInt8SupplierWithError,
			err:  true,
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
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt8SupplierResult, v)
			}
		})
	}
}

func TestMustInt8Supplier_ToInt8Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8Supplier
	}{
		{
			name: "ok",
			s:    testInt8Supplier,
		},
		{
			name: "with_error",
			s:    testInt8SupplierWithError,
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
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt8SupplierError.Error())
			} else {
				r.Equal(testInt8SupplierResult, v)
			}
		})
	}
}
