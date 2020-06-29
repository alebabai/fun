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
	testBytePtrSupplierResult *byte
	testBytePtrSupplierError  = errors.New("error")
)

func testBytePtrSupplier() (*byte, error) {
	return testBytePtrSupplierResult, nil
}

func testBytePtrSupplierWithError() (*byte, error) {
	return testBytePtrSupplierResult, testBytePtrSupplierError
}

func TestBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSupplier
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBytePtrSupplierError.Error())
			} else {
				r.Equal(testBytePtrSupplierResult, v)
			}
		})
	}
}

func TestBytePtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSupplierWithError,
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
				r.EqualError(err, testBytePtrSupplierError.Error())
			} else {
				r.Equal(testBytePtrSupplierResult, v)
			}
		})
	}
}

func TestBytePtrSupplier_ToSilentBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBytePtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBytePtrSupplierResult, v)
			}
		})
	}
}

func TestBytePtrSupplier_ToMustBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBytePtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBytePtrSupplierResult, v)
			}
		})
	}
}

func TestSilentBytePtrSupplier(t *testing.T) {
	var ss SilentBytePtrSupplier = func() *byte {
		return testBytePtrSupplierResult
	}
	v := ss()
	require.Equal(t, testBytePtrSupplierResult, v)
}

func TestMustBytePtrSupplier(t *testing.T) {
	var ms MustBytePtrSupplier = func() *byte {
		return testBytePtrSupplierResult
	}

	v := ms()
	require.Equal(t, testBytePtrSupplierResult, v)
}

func TestMustBytePtrSupplier_ToSilentBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBytePtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBytePtrSupplierResult, v)
			}
		})
	}
}

func TestMustBytePtrSupplier_ToBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSupplier
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSupplier()
			r.NotNil(ms)

			s := ms.ToBytePtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBytePtrSupplierError.Error())
			} else {
				r.Equal(testBytePtrSupplierResult, v)
			}
		})
	}
}
