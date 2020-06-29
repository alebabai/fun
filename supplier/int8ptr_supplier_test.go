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
	testInt8PtrSupplierResult *int8
	testInt8PtrSupplierError  = errors.New("error")
)

func testInt8PtrSupplier() (*int8, error) {
	return testInt8PtrSupplierResult, nil
}

func testInt8PtrSupplierWithError() (*int8, error) {
	return testInt8PtrSupplierResult, testInt8PtrSupplierError
}

func TestInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSupplier
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt8PtrSupplierError.Error())
			} else {
				r.Equal(testInt8PtrSupplierResult, v)
			}
		})
	}
}

func TestInt8PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSupplierWithError,
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
				r.EqualError(err, testInt8PtrSupplierError.Error())
			} else {
				r.Equal(testInt8PtrSupplierResult, v)
			}
		})
	}
}

func TestInt8PtrSupplier_ToSilentInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt8PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt8PtrSupplierResult, v)
			}
		})
	}
}

func TestInt8PtrSupplier_ToMustInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt8PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt8PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentInt8PtrSupplier(t *testing.T) {
	var ss SilentInt8PtrSupplier = func() *int8 {
		return testInt8PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testInt8PtrSupplierResult, v)
}

func TestMustInt8PtrSupplier(t *testing.T) {
	var ms MustInt8PtrSupplier = func() *int8 {
		return testInt8PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testInt8PtrSupplierResult, v)
}

func TestMustInt8PtrSupplier_ToSilentInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt8PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt8PtrSupplierResult, v)
			}
		})
	}
}

func TestMustInt8PtrSupplier_ToInt8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSupplier
	}{
		{
			name: "ok",
			s:    testInt8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSupplier()
			r.NotNil(ms)

			s := ms.ToInt8PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt8PtrSupplierError.Error())
			} else {
				r.Equal(testInt8PtrSupplierResult, v)
			}
		})
	}
}
