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
	testFloat32SupplierResult float32
	testFloat32SupplierError  = errors.New("error")
)

func testFloat32Supplier() (float32, error) {
	return testFloat32SupplierResult, nil
}

func testFloat32SupplierWithError() (float32, error) {
	return testFloat32SupplierResult, testFloat32SupplierError
}

func TestFloat32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Supplier
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name: "with_error",
			s:    testFloat32SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat32SupplierError.Error())
			} else {
				r.Equal(testFloat32SupplierResult, v)
			}
		})
	}
}

func TestFloat32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name: "with_error",
			s:    testFloat32SupplierWithError,
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
				r.EqualError(err, testFloat32SupplierError.Error())
			} else {
				r.Equal(testFloat32SupplierResult, v)
			}
		})
	}
}

func TestFloat32Supplier_ToSilentFloat32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name: "with_error",
			s:    testFloat32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32SupplierResult, v)
			}
		})
	}
}

func TestFloat32Supplier_ToMustFloat32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name: "with_error",
			s:    testFloat32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat32SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat32SupplierResult, v)
			}
		})
	}
}

func TestSilentFloat32Supplier(t *testing.T) {
	var ss SilentFloat32Supplier = func() float32 {
		return testFloat32SupplierResult
	}
	v := ss()
	require.Equal(t, testFloat32SupplierResult, v)
}

func TestMustFloat32Supplier(t *testing.T) {
	var ms MustFloat32Supplier = func() float32 {
		return testFloat32SupplierResult
	}

	v := ms()
	require.Equal(t, testFloat32SupplierResult, v)
}

func TestMustFloat32Supplier_ToSilentFloat32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name: "with_error",
			s:    testFloat32SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat32Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32SupplierResult, v)
			}
		})
	}
}

func TestMustFloat32Supplier_ToFloat32Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32Supplier
	}{
		{
			name: "ok",
			s:    testFloat32Supplier,
		},
		{
			name: "with_error",
			s:    testFloat32SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32Supplier()
			r.NotNil(ms)

			s := ms.ToFloat32Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat32SupplierError.Error())
			} else {
				r.Equal(testFloat32SupplierResult, v)
			}
		})
	}
}
