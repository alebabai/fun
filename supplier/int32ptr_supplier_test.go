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
	testInt32PtrSupplierResult *int32
	testInt32PtrSupplierError  = errors.New("error")
)

func testInt32PtrSupplier() (*int32, error) {
	return testInt32PtrSupplierResult, nil
}

func testInt32PtrSupplierWithError() (*int32, error) {
	return testInt32PtrSupplierResult, testInt32PtrSupplierError
}

func TestInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSupplier
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt32PtrSupplierError.Error())
			} else {
				r.Equal(testInt32PtrSupplierResult, v)
			}
		})
	}
}

func TestInt32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSupplierWithError,
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
				r.EqualError(err, testInt32PtrSupplierError.Error())
			} else {
				r.Equal(testInt32PtrSupplierResult, v)
			}
		})
	}
}

func TestInt32PtrSupplier_ToSilentInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt32PtrSupplierResult, v)
			}
		})
	}
}

func TestInt32PtrSupplier_ToMustInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt32PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt32PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentInt32PtrSupplier(t *testing.T) {
	var ss SilentInt32PtrSupplier = func() *int32 {
		return testInt32PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testInt32PtrSupplierResult, v)
}

func TestMustInt32PtrSupplier(t *testing.T) {
	var ms MustInt32PtrSupplier = func() *int32 {
		return testInt32PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testInt32PtrSupplierResult, v)
}

func TestMustInt32PtrSupplier_ToSilentInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt32PtrSupplierResult, v)
			}
		})
	}
}

func TestMustInt32PtrSupplier_ToInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSupplier
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSupplier()
			r.NotNil(ms)

			s := ms.ToInt32PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt32PtrSupplierError.Error())
			} else {
				r.Equal(testInt32PtrSupplierResult, v)
			}
		})
	}
}
