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
	testUint64SupplierResult uint64
	testUint64SupplierError  = errors.New("error")
)

func testUint64Supplier() (uint64, error) {
	return testUint64SupplierResult, nil
}

func testUint64SupplierWithError() (uint64, error) {
	return testUint64SupplierResult, testUint64SupplierError
}

func TestUint64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Supplier
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name: "with_error",
			s:    testUint64SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint64SupplierError.Error())
			} else {
				r.Equal(testUint64SupplierResult, v)
			}
		})
	}
}

func TestUint64Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name: "with_error",
			s:    testUint64SupplierWithError,
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
				r.EqualError(err, testUint64SupplierError.Error())
			} else {
				r.Equal(testUint64SupplierResult, v)
			}
		})
	}
}

func TestUint64Supplier_ToSilentUint64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name: "with_error",
			s:    testUint64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint64SupplierResult, v)
			}
		})
	}
}

func TestUint64Supplier_ToMustUint64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name: "with_error",
			s:    testUint64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint64SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint64SupplierResult, v)
			}
		})
	}
}

func TestSilentUint64Supplier(t *testing.T) {
	var ss SilentUint64Supplier = func() uint64 {
		return testUint64SupplierResult
	}
	v := ss()
	require.Equal(t, testUint64SupplierResult, v)
}

func TestSilentUint64Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name: "with_error",
			s:    testUint64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint64Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint64SupplierResult, v)
			}
		})
	}
}

func TestMustUint64Supplier(t *testing.T) {
	var ms MustUint64Supplier = func() uint64 {
		return testUint64SupplierResult
	}

	v := ms()
	require.Equal(t, testUint64SupplierResult, v)
}

func TestMustUint64Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name: "with_error",
			s:    testUint64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint64Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint64SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint64SupplierResult, v)
			}
		})
	}
}

func TestMustUint64Supplier_ToSilentUint64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name: "with_error",
			s:    testUint64SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint64Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint64SupplierResult, v)
			}
		})
	}
}

func TestMustUint64Supplier_ToUint64Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64Supplier
	}{
		{
			name: "ok",
			s:    testUint64Supplier,
		},
		{
			name: "with_error",
			s:    testUint64SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64Supplier()
			r.NotNil(ms)

			s := ms.ToUint64Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint64SupplierError.Error())
			} else {
				r.Equal(testUint64SupplierResult, v)
			}
		})
	}
}
