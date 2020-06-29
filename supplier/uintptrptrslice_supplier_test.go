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
	testUintptrPtrSliceSupplierResult []*uintptr
	testUintptrPtrSliceSupplierError  = errors.New("error")
)

func testUintptrPtrSliceSupplier() ([]*uintptr, error) {
	return testUintptrPtrSliceSupplierResult, nil
}

func testUintptrPtrSliceSupplierWithError() ([]*uintptr, error) {
	return testUintptrPtrSliceSupplierResult, testUintptrPtrSliceSupplierError
}

func TestUintptrPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUintptrPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintptrPtrSliceSupplierError.Error())
			} else {
				r.Equal(testUintptrPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUintptrPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintptrPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSliceSupplierWithError,
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
				r.EqualError(err, testUintptrPtrSliceSupplierError.Error())
			} else {
				r.Equal(testUintptrPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUintptrPtrSliceSupplier_ToSilentUintptrPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintptrPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUintptrPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintptrPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUintptrPtrSliceSupplier_ToMustUintptrPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintptrPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintptrPtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUintptrPtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUintptrPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUintptrPtrSliceSupplier(t *testing.T) {
	var ss SilentUintptrPtrSliceSupplier = func() []*uintptr {
		return testUintptrPtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUintptrPtrSliceSupplierResult, v)
}

func TestMustUintptrPtrSliceSupplier(t *testing.T) {
	var ms MustUintptrPtrSliceSupplier = func() []*uintptr {
		return testUintptrPtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUintptrPtrSliceSupplierResult, v)
}

func TestMustUintptrPtrSliceSupplier_ToSilentUintptrPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintptrPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintptrPtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUintptrPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintptrPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUintptrPtrSliceSupplier_ToUintptrPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUintptrPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintptrPtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUintptrPtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintptrPtrSliceSupplierError.Error())
			} else {
				r.Equal(testUintptrPtrSliceSupplierResult, v)
			}
		})
	}
}
