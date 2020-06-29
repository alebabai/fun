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
	testFloat64PtrSupplierResult *float64
	testFloat64PtrSupplierError  = errors.New("error")
)

func testFloat64PtrSupplier() (*float64, error) {
	return testFloat64PtrSupplierResult, nil
}

func testFloat64PtrSupplierWithError() (*float64, error) {
	return testFloat64PtrSupplierResult, testFloat64PtrSupplierError
}

func TestFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSupplier
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat64PtrSupplierError.Error())
			} else {
				r.Equal(testFloat64PtrSupplierResult, v)
			}
		})
	}
}

func TestFloat64PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSupplierWithError,
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
				r.EqualError(err, testFloat64PtrSupplierError.Error())
			} else {
				r.Equal(testFloat64PtrSupplierResult, v)
			}
		})
	}
}

func TestFloat64PtrSupplier_ToSilentFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat64PtrSupplierResult, v)
			}
		})
	}
}

func TestFloat64PtrSupplier_ToMustFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat64PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat64PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentFloat64PtrSupplier(t *testing.T) {
	var ss SilentFloat64PtrSupplier = func() *float64 {
		return testFloat64PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testFloat64PtrSupplierResult, v)
}

func TestMustFloat64PtrSupplier(t *testing.T) {
	var ms MustFloat64PtrSupplier = func() *float64 {
		return testFloat64PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testFloat64PtrSupplierResult, v)
}

func TestMustFloat64PtrSupplier_ToSilentFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat64PtrSupplierResult, v)
			}
		})
	}
}

func TestMustFloat64PtrSupplier_ToFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float64PtrSupplier
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat64PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSupplier()
			r.NotNil(ms)

			s := ms.ToFloat64PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat64PtrSupplierError.Error())
			} else {
				r.Equal(testFloat64PtrSupplierResult, v)
			}
		})
	}
}
