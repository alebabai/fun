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
	testUint32PtrSupplierResult *uint32
	testUint32PtrSupplierError  = errors.New("error")
)

func testUint32PtrSupplier() (*uint32, error) {
	return testUint32PtrSupplierResult, nil
}

func testUint32PtrSupplierWithError() (*uint32, error) {
	return testUint32PtrSupplierResult, testUint32PtrSupplierError
}

func TestUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSupplier
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint32PtrSupplierError.Error())
			} else {
				r.Equal(testUint32PtrSupplierResult, v)
			}
		})
	}
}

func TestUint32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSupplierWithError,
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
				r.EqualError(err, testUint32PtrSupplierError.Error())
			} else {
				r.Equal(testUint32PtrSupplierResult, v)
			}
		})
	}
}

func TestUint32PtrSupplier_ToSilentUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32PtrSupplierResult, v)
			}
		})
	}
}

func TestUint32PtrSupplier_ToMustUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint32PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint32PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentUint32PtrSupplier(t *testing.T) {
	var ss SilentUint32PtrSupplier = func() *uint32 {
		return testUint32PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testUint32PtrSupplierResult, v)
}

func TestMustUint32PtrSupplier(t *testing.T) {
	var ms MustUint32PtrSupplier = func() *uint32 {
		return testUint32PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testUint32PtrSupplierResult, v)
}

func TestMustUint32PtrSupplier_ToSilentUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32PtrSupplierResult, v)
			}
		})
	}
}

func TestMustUint32PtrSupplier_ToUint32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSupplier
	}{
		{
			name: "ok",
			s:    testUint32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSupplier()
			r.NotNil(ms)

			s := ms.ToUint32PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint32PtrSupplierError.Error())
			} else {
				r.Equal(testUint32PtrSupplierResult, v)
			}
		})
	}
}
