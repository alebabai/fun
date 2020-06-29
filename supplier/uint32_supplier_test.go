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
	testUint32SupplierResult uint32
	testUint32SupplierError  = errors.New("error")
)

func testUint32Supplier() (uint32, error) {
	return testUint32SupplierResult, nil
}

func testUint32SupplierWithError() (uint32, error) {
	return testUint32SupplierResult, testUint32SupplierError
}

func TestUint32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Supplier
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name: "with_error",
			s:    testUint32SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint32SupplierError.Error())
			} else {
				r.Equal(testUint32SupplierResult, v)
			}
		})
	}
}

func TestUint32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name: "with_error",
			s:    testUint32SupplierWithError,
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
				r.EqualError(err, testUint32SupplierError.Error())
			} else {
				r.Equal(testUint32SupplierResult, v)
			}
		})
	}
}

func TestUint32Supplier_ToSilentUint32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name: "with_error",
			s:    testUint32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32SupplierResult, v)
			}
		})
	}
}

func TestUint32Supplier_ToMustUint32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name: "with_error",
			s:    testUint32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint32SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint32SupplierResult, v)
			}
		})
	}
}

func TestSilentUint32Supplier(t *testing.T) {
	var ss SilentUint32Supplier = func() uint32 {
		return testUint32SupplierResult
	}
	v := ss()
	require.Equal(t, testUint32SupplierResult, v)
}

func TestMustUint32Supplier(t *testing.T) {
	var ms MustUint32Supplier = func() uint32 {
		return testUint32SupplierResult
	}

	v := ms()
	require.Equal(t, testUint32SupplierResult, v)
}

func TestMustUint32Supplier_ToSilentUint32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name: "with_error",
			s:    testUint32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32SupplierResult, v)
			}
		})
	}
}

func TestMustUint32Supplier_ToUint32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32Supplier
	}{
		{
			name: "ok",
			s:    testUint32Supplier,
		},
		{
			name: "with_error",
			s:    testUint32SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32Supplier()
			r.NotNil(ms)

			s := ms.ToUint32Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint32SupplierError.Error())
			} else {
				r.Equal(testUint32SupplierResult, v)
			}
		})
	}
}
