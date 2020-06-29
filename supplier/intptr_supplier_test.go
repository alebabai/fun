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
	testIntPtrSupplierResult *int
	testIntPtrSupplierError  = errors.New("error")
)

func testIntPtrSupplier() (*int, error) {
	return testIntPtrSupplierResult, nil
}

func testIntPtrSupplierWithError() (*int, error) {
	return testIntPtrSupplierResult, testIntPtrSupplierError
}

func TestIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSupplier
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testIntPtrSupplierError.Error())
			} else {
				r.Equal(testIntPtrSupplierResult, v)
			}
		})
	}
}

func TestIntPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSupplierWithError,
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
				r.EqualError(err, testIntPtrSupplierError.Error())
			} else {
				r.Equal(testIntPtrSupplierResult, v)
			}
		})
	}
}

func TestIntPtrSupplier_ToSilentIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentIntPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testIntPtrSupplierResult, v)
			}
		})
	}
}

func TestIntPtrSupplier_ToMustIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testIntPtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testIntPtrSupplierResult, v)
			}
		})
	}
}

func TestSilentIntPtrSupplier(t *testing.T) {
	var ss SilentIntPtrSupplier = func() *int {
		return testIntPtrSupplierResult
	}
	v := ss()
	require.Equal(t, testIntPtrSupplierResult, v)
}

func TestMustIntPtrSupplier(t *testing.T) {
	var ms MustIntPtrSupplier = func() *int {
		return testIntPtrSupplierResult
	}

	v := ms()
	require.Equal(t, testIntPtrSupplierResult, v)
}

func TestMustIntPtrSupplier_ToSilentIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentIntPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testIntPtrSupplierResult, v)
			}
		})
	}
}

func TestMustIntPtrSupplier_ToIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSupplier
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSupplier()
			r.NotNil(ms)

			s := ms.ToIntPtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testIntPtrSupplierError.Error())
			} else {
				r.Equal(testIntPtrSupplierResult, v)
			}
		})
	}
}
