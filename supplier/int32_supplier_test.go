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
	testInt32SupplierResult int32
	testInt32SupplierError  = errors.New("error")
)

func testInt32Supplier() (int32, error) {
	return testInt32SupplierResult, nil
}

func testInt32SupplierWithError() (int32, error) {
	return testInt32SupplierResult, testInt32SupplierError
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
				r.EqualError(err, testInt32SupplierError.Error())
			} else {
				r.Equal(testInt32SupplierResult, v)
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
				r.EqualError(err, testInt32SupplierError.Error())
			} else {
				r.Equal(testInt32SupplierResult, v)
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
				r.Equal(testInt32SupplierResult, v)
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
				r.PanicsWithError(testInt32SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt32SupplierResult, v)
			}
		})
	}
}

func TestSilentInt32Supplier(t *testing.T) {
	var ss SilentInt32Supplier = func() int32 {
		return testInt32SupplierResult
	}
	v := ss()
	require.Equal(t, testInt32SupplierResult, v)
}

func TestMustInt32Supplier(t *testing.T) {
	var ms MustInt32Supplier = func() int32 {
		return testInt32SupplierResult
	}

	v := ms()
	require.Equal(t, testInt32SupplierResult, v)
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
				r.Equal(testInt32SupplierResult, v)
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
				r.EqualError(err, testInt32SupplierError.Error())
			} else {
				r.Equal(testInt32SupplierResult, v)
			}
		})
	}
}
