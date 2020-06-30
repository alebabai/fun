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
	testFloat64SupplierResult float64
	testFloat64SupplierError  = errors.New("error")
)

func testFloat64Supplier() (float64, error) {
	return testFloat64SupplierResult, nil
}

func testFloat64SupplierWithError() (float64, error) {
	return testFloat64SupplierResult, testFloat64SupplierError
}

func TestFloat64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Supplier
	}{
		{
			name: "ok",
			s:    testFloat64Supplier,
		},
		{
			name: "with_error",
			s:    testFloat64SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat64SupplierError.Error())
			} else {
				r.Equal(testFloat64SupplierResult, v)
			}
		})
	}
}

func TestFloat64Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64Supplier,
		},
		{
			name: "with_error",
			s:    testFloat64SupplierWithError,
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
				r.EqualError(err, testFloat64SupplierError.Error())
			} else {
				r.Equal(testFloat64SupplierResult, v)
			}
		})
	}
}

func TestFloat64Supplier_ToSilentFloat64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64Supplier,
		},
		{
			name: "with_error",
			s:    testFloat64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat64SupplierResult, v)
			}
		})
	}
}

func TestFloat64Supplier_ToMustFloat64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64Supplier,
		},
		{
			name: "with_error",
			s:    testFloat64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat64SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat64SupplierResult, v)
			}
		})
	}
}

func TestSilentFloat64Supplier(t *testing.T) {
	var ss SilentFloat64Supplier = func() float64 {
		return testFloat64SupplierResult
	}
	v := ss()
	require.Equal(t, testFloat64SupplierResult, v)
}

func TestSilentFloat64Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64Supplier,
		},
		{
			name: "with_error",
			s:    testFloat64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat64Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat64SupplierResult, v)
			}
		})
	}
}

func TestMustFloat64Supplier(t *testing.T) {
	var ms MustFloat64Supplier = func() float64 {
		return testFloat64SupplierResult
	}

	v := ms()
	require.Equal(t, testFloat64SupplierResult, v)
}

func TestMustFloat64Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64Supplier,
		},
		{
			name: "with_error",
			s:    testFloat64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat64Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat64SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat64SupplierResult, v)
			}
		})
	}
}

func TestMustFloat64Supplier_ToSilentFloat64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64Supplier,
		},
		{
			name: "with_error",
			s:    testFloat64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat64SupplierResult, v)
			}
		})
	}
}

func TestMustFloat64Supplier_ToFloat64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64Supplier
	}{
		{
			name: "ok",
			s:    testFloat64Supplier,
		},
		{
			name: "with_error",
			s:    testFloat64SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64Supplier()
			r.NotNil(ms)

			s := ms.ToFloat64Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat64SupplierError.Error())
			} else {
				r.Equal(testFloat64SupplierResult, v)
			}
		})
	}
}
