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
	testIntSliceSupplierResult []int
	testIntSliceSupplierError  = errors.New("error")
)

func testIntSliceSupplier() ([]int, error) {
	return testIntSliceSupplierResult, nil
}

func testIntSliceSupplierWithError() ([]int, error) {
	return testIntSliceSupplierResult, testIntSliceSupplierError
}

func TestIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSliceSupplier
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testIntSliceSupplierError.Error())
			} else {
				r.Equal(testIntSliceSupplierResult, v)
			}
		})
	}
}

func TestIntSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntSliceSupplierWithError,
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
				r.EqualError(err, testIntSliceSupplierError.Error())
			} else {
				r.Equal(testIntSliceSupplierResult, v)
			}
		})
	}
}

func TestIntSliceSupplier_ToSilentIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentIntSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testIntSliceSupplierResult, v)
			}
		})
	}
}

func TestIntSliceSupplier_ToMustIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testIntSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testIntSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentIntSliceSupplier(t *testing.T) {
	var ss SilentIntSliceSupplier = func() []int {
		return testIntSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testIntSliceSupplierResult, v)
}

func TestMustIntSliceSupplier(t *testing.T) {
	var ms MustIntSliceSupplier = func() []int {
		return testIntSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testIntSliceSupplierResult, v)
}

func TestMustIntSliceSupplier_ToSilentIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentIntSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testIntSliceSupplierResult, v)
			}
		})
	}
}

func TestMustIntSliceSupplier_ToIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSliceSupplier
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSliceSupplier()
			r.NotNil(ms)

			s := ms.ToIntSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testIntSliceSupplierError.Error())
			} else {
				r.Equal(testIntSliceSupplierResult, v)
			}
		})
	}
}
