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
	testInt16SliceSupplierResult []int16
	testInt16SliceSupplierError  = errors.New("error")
)

func testInt16SliceSupplier() ([]int16, error) {
	return testInt16SliceSupplierResult, nil
}

func testInt16SliceSupplierWithError() ([]int16, error) {
	return testInt16SliceSupplierResult, testInt16SliceSupplierError
}

func TestInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt16SliceSupplierError.Error())
			} else {
				r.Equal(testInt16SliceSupplierResult, v)
			}
		})
	}
}

func TestInt16SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
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
				r.EqualError(err, testInt16SliceSupplierError.Error())
			} else {
				r.Equal(testInt16SliceSupplierResult, v)
			}
		})
	}
}

func TestInt16SliceSupplier_ToSilentInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt16SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt16SliceSupplierResult, v)
			}
		})
	}
}

func TestInt16SliceSupplier_ToMustInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt16SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt16SliceSupplierResult, v)
			}
		})
	}
}

func TestSilentInt16SliceSupplier(t *testing.T) {
	var ss SilentInt16SliceSupplier = func() []int16 {
		return testInt16SliceSupplierResult
	}
	v := ss()
	require.Equal(t, testInt16SliceSupplierResult, v)
}

func TestMustInt16SliceSupplier(t *testing.T) {
	var ms MustInt16SliceSupplier = func() []int16 {
		return testInt16SliceSupplierResult
	}

	v := ms()
	require.Equal(t, testInt16SliceSupplierResult, v)
}

func TestMustInt16SliceSupplier_ToSilentInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt16SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt16SliceSupplierResult, v)
			}
		})
	}
}

func TestMustInt16SliceSupplier_ToInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16SliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt16SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt16SliceSupplierError.Error())
			} else {
				r.Equal(testInt16SliceSupplierResult, v)
			}
		})
	}
}
