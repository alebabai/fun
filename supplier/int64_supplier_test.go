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
	testInt64SupplierResult int64
	testInt64SupplierError  = errors.New("error")
)

func testInt64Supplier() (int64, error) {
	return testInt64SupplierResult, nil
}

func testInt64SupplierWithError() (int64, error) {
	return testInt64SupplierResult, testInt64SupplierError
}

func TestInt64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Supplier
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name: "with_error",
			s:    testInt64SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt64SupplierError.Error())
			} else {
				r.Equal(testInt64SupplierResult, v)
			}
		})
	}
}

func TestInt64Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name: "with_error",
			s:    testInt64SupplierWithError,
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
				r.EqualError(err, testInt64SupplierError.Error())
			} else {
				r.Equal(testInt64SupplierResult, v)
			}
		})
	}
}

func TestInt64Supplier_ToSilentInt64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name: "with_error",
			s:    testInt64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt64SupplierResult, v)
			}
		})
	}
}

func TestInt64Supplier_ToMustInt64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name: "with_error",
			s:    testInt64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt64SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt64SupplierResult, v)
			}
		})
	}
}

func TestSilentInt64Supplier(t *testing.T) {
	var ss SilentInt64Supplier = func() int64 {
		return testInt64SupplierResult
	}
	v := ss()
	require.Equal(t, testInt64SupplierResult, v)
}

func TestSilentInt64Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name: "with_error",
			s:    testInt64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt64Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt64SupplierResult, v)
			}
		})
	}
}

func TestMustInt64Supplier(t *testing.T) {
	var ms MustInt64Supplier = func() int64 {
		return testInt64SupplierResult
	}

	v := ms()
	require.Equal(t, testInt64SupplierResult, v)
}

func TestMustInt64Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name: "with_error",
			s:    testInt64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt64Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt64SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt64SupplierResult, v)
			}
		})
	}
}

func TestMustInt64Supplier_ToSilentInt64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name: "with_error",
			s:    testInt64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt64SupplierResult, v)
			}
		})
	}
}

func TestMustInt64Supplier_ToInt64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64Supplier
	}{
		{
			name: "ok",
			s:    testInt64Supplier,
		},
		{
			name: "with_error",
			s:    testInt64SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64Supplier()
			r.NotNil(ms)

			s := ms.ToInt64Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt64SupplierError.Error())
			} else {
				r.Equal(testInt64SupplierResult, v)
			}
		})
	}
}
