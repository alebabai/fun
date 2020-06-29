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
	testUint8SupplierResult uint8
	testUint8SupplierError  = errors.New("error")
)

func testUint8Supplier() (uint8, error) {
	return testUint8SupplierResult, nil
}

func testUint8SupplierWithError() (uint8, error) {
	return testUint8SupplierResult, testUint8SupplierError
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
				r.EqualError(err, testUint8SupplierError.Error())
			} else {
				r.Equal(testUint8SupplierResult, v)
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
				r.EqualError(err, testUint8SupplierError.Error())
			} else {
				r.Equal(testUint8SupplierResult, v)
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
				r.Equal(testUint8SupplierResult, v)
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
				r.PanicsWithError(testUint8SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint8SupplierResult, v)
			}
		})
	}
}

func TestSilentUint8Supplier(t *testing.T) {
	var ss SilentUint8Supplier = func() uint8 {
		return testUint8SupplierResult
	}
	v := ss()
	require.Equal(t, testUint8SupplierResult, v)
}

func TestMustUint8Supplier(t *testing.T) {
	var ms MustUint8Supplier = func() uint8 {
		return testUint8SupplierResult
	}

	v := ms()
	require.Equal(t, testUint8SupplierResult, v)
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
				r.Equal(testUint8SupplierResult, v)
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
				r.EqualError(err, testUint8SupplierError.Error())
			} else {
				r.Equal(testUint8SupplierResult, v)
			}
		})
	}
}
