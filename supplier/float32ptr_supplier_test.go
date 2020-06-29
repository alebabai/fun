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
	testFloat32PtrSupplierResult *float32
	testFloat32PtrSupplierError  = errors.New("error")
)

func testFloat32PtrSupplier() (*float32, error) {
	return testFloat32PtrSupplierResult, nil
}

func testFloat32PtrSupplierWithError() (*float32, error) {
	return testFloat32PtrSupplierResult, testFloat32PtrSupplierError
}

func TestFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSupplier
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat32PtrSupplierError.Error())
			} else {
				r.Equal(testFloat32PtrSupplierResult, v)
			}
		})
	}
}

func TestFloat32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSupplierWithError,
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
				r.EqualError(err, testFloat32PtrSupplierError.Error())
			} else {
				r.Equal(testFloat32PtrSupplierResult, v)
			}
		})
	}
}

func TestFloat32PtrSupplier_ToSilentFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32PtrSupplierResult, v)
			}
		})
	}
}

func TestFloat32PtrSupplier_ToMustFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat32PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat32PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentFloat32PtrSupplier(t *testing.T) {
	var ss SilentFloat32PtrSupplier = func() *float32 {
		return testFloat32PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testFloat32PtrSupplierResult, v)
}

func TestMustFloat32PtrSupplier(t *testing.T) {
	var ms MustFloat32PtrSupplier = func() *float32 {
		return testFloat32PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testFloat32PtrSupplierResult, v)
}

func TestMustFloat32PtrSupplier_ToSilentFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32PtrSupplierResult, v)
			}
		})
	}
}

func TestMustFloat32PtrSupplier_ToFloat32PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSupplier
	}{
		{
			name: "ok",
			s:    testFloat32PtrSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSupplier()
			r.NotNil(ms)

			s := ms.ToFloat32PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat32PtrSupplierError.Error())
			} else {
				r.Equal(testFloat32PtrSupplierResult, v)
			}
		})
	}
}
