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
	testFloat64PtrSliceSupplierResult []*float64
	testFloat64PtrSliceSupplierError  = errors.New("error")
)

func testFloat64PtrSliceSupplier() ([]*float64, error) {
	return testFloat64PtrSliceSupplierResult, nil
}

func testFloat64PtrSliceSupplierWithError() ([]*float64, error) {
	return testFloat64PtrSliceSupplierResult, testFloat64PtrSliceSupplierError
}

func TestFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat64PtrSliceSupplierError.Error())
			} else {
				r.Equal(testFloat64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestFloat64PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSliceSupplierWithError,
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
				r.EqualError(err, testFloat64PtrSliceSupplierError.Error())
			} else {
				r.Equal(testFloat64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestFloat64PtrSliceSupplier_ToSilentFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestFloat64PtrSliceSupplier_ToMustFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat64PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentFloat64PtrSliceSupplier(t *testing.T) {
	var ss SilentFloat64PtrSliceSupplier = func() []*float64 {
		return testFloat64PtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testFloat64PtrSliceSupplierResult, v)
}

func TestMustFloat64PtrSliceSupplier(t *testing.T) {
	var ms MustFloat64PtrSliceSupplier = func() []*float64 {
		return testFloat64PtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testFloat64PtrSliceSupplierResult, v)
}

func TestMustFloat64PtrSliceSupplier_ToSilentFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustFloat64PtrSliceSupplier_ToFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToFloat64PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat64PtrSliceSupplierError.Error())
			} else {
				r.Equal(testFloat64PtrSliceSupplierResult, v)
			}
		})
	}
}
