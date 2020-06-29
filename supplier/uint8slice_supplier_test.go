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
	testUint8SliceSupplierResult []uint8
	testUint8SliceSupplierError  = errors.New("error")
)

func testUint8SliceSupplier() ([]uint8, error) {
	return testUint8SliceSupplierResult, nil
}

func testUint8SliceSupplierWithError() ([]uint8, error) {
	return testUint8SliceSupplierResult, testUint8SliceSupplierError
}

func TestUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint8SliceSupplierError.Error())
			} else {
				r.Equal(testUint8SliceSupplierResult, v)
			}
		})
	}
}

func TestUint8SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
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
				r.EqualError(err, testUint8SliceSupplierError.Error())
			} else {
				r.Equal(testUint8SliceSupplierResult, v)
			}
		})
	}
}

func TestUint8SliceSupplier_ToSilentUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint8SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint8SliceSupplierResult, v)
			}
		})
	}
}

func TestUint8SliceSupplier_ToMustUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint8SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint8SliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUint8SliceSupplier(t *testing.T) {
	var ss SilentUint8SliceSupplier = func() []uint8 {
		return testUint8SliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUint8SliceSupplierResult, v)
}

func TestMustUint8SliceSupplier(t *testing.T) {
	var ms MustUint8SliceSupplier = func() []uint8 {
		return testUint8SliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUint8SliceSupplierResult, v)
}

func TestMustUint8SliceSupplier_ToSilentUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint8SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint8SliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint8SliceSupplier_ToUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8SliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint8SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint8SliceSupplierError.Error())
			} else {
				r.Equal(testUint8SliceSupplierResult, v)
			}
		})
	}
}
