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
	testUint16SliceSupplierResult []uint16
	testUint16SliceSupplierError  = errors.New("error")
)

func testUint16SliceSupplier() ([]uint16, error) {
	return testUint16SliceSupplierResult, nil
}

func testUint16SliceSupplierWithError() ([]uint16, error) {
	return testUint16SliceSupplierResult, testUint16SliceSupplierError
}

func TestUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint16SliceSupplierError.Error())
			} else {
				r.Equal(testUint16SliceSupplierResult, v)
			}
		})
	}
}

func TestUint16SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
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
				r.EqualError(err, testUint16SliceSupplierError.Error())
			} else {
				r.Equal(testUint16SliceSupplierResult, v)
			}
		})
	}
}

func TestUint16SliceSupplier_ToSilentUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint16SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint16SliceSupplierResult, v)
			}
		})
	}
}

func TestUint16SliceSupplier_ToMustUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint16SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint16SliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUint16SliceSupplier(t *testing.T) {
	var ss SilentUint16SliceSupplier = func() []uint16 {
		return testUint16SliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUint16SliceSupplierResult, v)
}

func TestMustUint16SliceSupplier(t *testing.T) {
	var ms MustUint16SliceSupplier = func() []uint16 {
		return testUint16SliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUint16SliceSupplierResult, v)
}

func TestMustUint16SliceSupplier_ToSilentUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint16SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint16SliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint16SliceSupplier_ToUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16SliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint16SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint16SliceSupplierError.Error())
			} else {
				r.Equal(testUint16SliceSupplierResult, v)
			}
		})
	}
}
