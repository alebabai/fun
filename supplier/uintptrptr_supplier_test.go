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
	testUintptrPtrSupplierResult *uintptr
	testUintptrPtrSupplierError  = errors.New("error")
)

func testUintptrPtrSupplier() (*uintptr, error) {
	return testUintptrPtrSupplierResult, nil
}

func testUintptrPtrSupplierWithError() (*uintptr, error) {
	return testUintptrPtrSupplierResult, testUintptrPtrSupplierError
}

func TestUintptrPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSupplier
	}{
		{
			name: "ok",
			s:    testUintptrPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintptrPtrSupplierError.Error())
			} else {
				r.Equal(testUintptrPtrSupplierResult, v)
			}
		})
	}
}

func TestUintptrPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintptrPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSupplierWithError,
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
				r.EqualError(err, testUintptrPtrSupplierError.Error())
			} else {
				r.Equal(testUintptrPtrSupplierResult, v)
			}
		})
	}
}

func TestUintptrPtrSupplier_ToSilentUintptrPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintptrPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUintptrPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintptrPtrSupplierResult, v)
			}
		})
	}
}

func TestUintptrPtrSupplier_ToMustUintptrPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintptrPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintptrPtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUintptrPtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUintptrPtrSupplierResult, v)
			}
		})
	}
}

func TestSilentUintptrPtrSupplier(t *testing.T) {
	var ss SilentUintptrPtrSupplier = func() *uintptr {
		return testUintptrPtrSupplierResult
	}
	v := ss()
	require.Equal(t, testUintptrPtrSupplierResult, v)
}

func TestMustUintptrPtrSupplier(t *testing.T) {
	var ms MustUintptrPtrSupplier = func() *uintptr {
		return testUintptrPtrSupplierResult
	}

	v := ms()
	require.Equal(t, testUintptrPtrSupplierResult, v)
}

func TestMustUintptrPtrSupplier_ToSilentUintptrPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintptrPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintptrPtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUintptrPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintptrPtrSupplierResult, v)
			}
		})
	}
}

func TestMustUintptrPtrSupplier_ToUintptrPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintptrPtrSupplier
	}{
		{
			name: "ok",
			s:    testUintptrPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintptrPtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintptrPtrSupplier()
			r.NotNil(ms)

			s := ms.ToUintptrPtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintptrPtrSupplierError.Error())
			} else {
				r.Equal(testUintptrPtrSupplierResult, v)
			}
		})
	}
}
