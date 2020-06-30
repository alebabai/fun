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
	testUint16SupplierResult uint16
	testUint16SupplierError  = errors.New("error")
)

func testUint16Supplier() (uint16, error) {
	return testUint16SupplierResult, nil
}

func testUint16SupplierWithError() (uint16, error) {
	return testUint16SupplierResult, testUint16SupplierError
}

func TestUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint16SupplierError.Error())
			} else {
				r.Equal(testUint16SupplierResult, v)
			}
		})
	}
}

func TestUint16Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
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
				r.EqualError(err, testUint16SupplierError.Error())
			} else {
				r.Equal(testUint16SupplierResult, v)
			}
		})
	}
}

func TestUint16Supplier_ToSilentUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint16Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint16SupplierResult, v)
			}
		})
	}
}

func TestUint16Supplier_ToMustUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16Supplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint16SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint16SupplierResult, v)
			}
		})
	}
}

func TestSilentUint16Supplier(t *testing.T) {
	var ss SilentUint16Supplier = func() uint16 {
		return testUint16SupplierResult
	}
	v := ss()
	require.Equal(t, testUint16SupplierResult, v)
}

func TestSilentUint16Supplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint16Supplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint16SupplierResult, v)
			}
		})
	}
}

func TestMustUint16Supplier(t *testing.T) {
	var ms MustUint16Supplier = func() uint16 {
		return testUint16SupplierResult
	}

	v := ms()
	require.Equal(t, testUint16SupplierResult, v)
}

func TestMustUint16Supplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint16Supplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint16SupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint16SupplierResult, v)
			}
		})
	}
}

func TestMustUint16Supplier_ToSilentUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16Supplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint16Supplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint16SupplierResult, v)
			}
		})
	}
}

func TestMustUint16Supplier_ToUint16Supplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16Supplier
	}{
		{
			name: "ok",
			s:    testUint16Supplier,
		},
		{
			name: "with_error",
			s:    testUint16SupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16Supplier()
			r.NotNil(ms)

			s := ms.ToUint16Supplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint16SupplierError.Error())
			} else {
				r.Equal(testUint16SupplierResult, v)
			}
		})
	}
}
