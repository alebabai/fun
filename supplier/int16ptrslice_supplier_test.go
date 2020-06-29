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
	testInt16PtrSliceSupplierResult []*int16
	testInt16PtrSliceSupplierError  = errors.New("error")
)

func testInt16PtrSliceSupplier() ([]*int16, error) {
	return testInt16PtrSliceSupplierResult, nil
}

func testInt16PtrSliceSupplierWithError() ([]*int16, error) {
	return testInt16PtrSliceSupplierResult, testInt16PtrSliceSupplierError
}

func TestInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt16PtrSliceSupplierError.Error())
			} else {
				r.Equal(testInt16PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestInt16PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSliceSupplierWithError,
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
				r.EqualError(err, testInt16PtrSliceSupplierError.Error())
			} else {
				r.Equal(testInt16PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestInt16PtrSliceSupplier_ToSilentInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt16PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt16PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestInt16PtrSliceSupplier_ToMustInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt16PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt16PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentInt16PtrSliceSupplier(t *testing.T) {
	var ss SilentInt16PtrSliceSupplier = func() []*int16 {
		return testInt16PtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testInt16PtrSliceSupplierResult, v)
}

func TestMustInt16PtrSliceSupplier(t *testing.T) {
	var ms MustInt16PtrSliceSupplier = func() []*int16 {
		return testInt16PtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testInt16PtrSliceSupplierResult, v)
}

func TestMustInt16PtrSliceSupplier_ToSilentInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt16PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt16PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustInt16PtrSliceSupplier_ToInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt16PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt16PtrSliceSupplierError.Error())
			} else {
				r.Equal(testInt16PtrSliceSupplierResult, v)
			}
		})
	}
}
