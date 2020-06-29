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
	testUint64SliceSupplierResult []uint64
	testUint64SliceSupplierError  = errors.New("error")
)

func testUint64SliceSupplier() ([]uint64, error) {
	return testUint64SliceSupplierResult, nil
}

func testUint64SliceSupplierWithError() ([]uint64, error) {
	return testUint64SliceSupplierResult, testUint64SliceSupplierError
}

func TestUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint64SliceSupplierError.Error())
			} else {
				r.Equal(testUint64SliceSupplierResult, v)
			}
		})
	}
}

func TestUint64SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64SliceSupplierWithError,
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
				r.EqualError(err, testUint64SliceSupplierError.Error())
			} else {
				r.Equal(testUint64SliceSupplierResult, v)
			}
		})
	}
}

func TestUint64SliceSupplier_ToSilentUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint64SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint64SliceSupplierResult, v)
			}
		})
	}
}

func TestUint64SliceSupplier_ToMustUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint64SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint64SliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUint64SliceSupplier(t *testing.T) {
	var ss SilentUint64SliceSupplier = func() []uint64 {
		return testUint64SliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUint64SliceSupplierResult, v)
}

func TestMustUint64SliceSupplier(t *testing.T) {
	var ms MustUint64SliceSupplier = func() []uint64 {
		return testUint64SliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUint64SliceSupplierResult, v)
}

func TestMustUint64SliceSupplier_ToSilentUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint64SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint64SliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint64SliceSupplier_ToUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64SliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint64SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint64SliceSupplierError.Error())
			} else {
				r.Equal(testUint64SliceSupplierResult, v)
			}
		})
	}
}
