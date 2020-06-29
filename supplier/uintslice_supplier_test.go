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
	testUintSliceSupplierResult []uint
	testUintSliceSupplierError  = errors.New("error")
)

func testUintSliceSupplier() ([]uint, error) {
	return testUintSliceSupplierResult, nil
}

func testUintSliceSupplierWithError() ([]uint, error) {
	return testUintSliceSupplierResult, testUintSliceSupplierError
}

func TestUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintSliceSupplier
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintSliceSupplierError.Error())
			} else {
				r.Equal(testUintSliceSupplierResult, v)
			}
		})
	}
}

func TestUintSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintSliceSupplierWithError,
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
				r.EqualError(err, testUintSliceSupplierError.Error())
			} else {
				r.Equal(testUintSliceSupplierResult, v)
			}
		})
	}
}

func TestUintSliceSupplier_ToSilentUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUintSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintSliceSupplierResult, v)
			}
		})
	}
}

func TestUintSliceSupplier_ToMustUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUintSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUintSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUintSliceSupplier(t *testing.T) {
	var ss SilentUintSliceSupplier = func() []uint {
		return testUintSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUintSliceSupplierResult, v)
}

func TestMustUintSliceSupplier(t *testing.T) {
	var ms MustUintSliceSupplier = func() []uint {
		return testUintSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUintSliceSupplierResult, v)
}

func TestMustUintSliceSupplier_ToSilentUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUintSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUintSliceSupplier_ToUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintSliceSupplier
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUintSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintSliceSupplierError.Error())
			} else {
				r.Equal(testUintSliceSupplierResult, v)
			}
		})
	}
}
