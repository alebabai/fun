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
	testUint32PtrSliceSupplierResult []*uint32
	testUint32PtrSliceSupplierError  = errors.New("error")
)

func testUint32PtrSliceSupplier() ([]*uint32, error) {
	return testUint32PtrSliceSupplierResult, nil
}

func testUint32PtrSliceSupplierWithError() ([]*uint32, error) {
	return testUint32PtrSliceSupplierResult, testUint32PtrSliceSupplierError
}

func TestUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUint32PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
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
				r.EqualError(err, testUint32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUint32PtrSliceSupplier_ToSilentUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUint32PtrSliceSupplier_ToMustUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint32PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUint32PtrSliceSupplier(t *testing.T) {
	var ss SilentUint32PtrSliceSupplier = func() []*uint32 {
		return testUint32PtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUint32PtrSliceSupplierResult, v)
}

func TestMustUint32PtrSliceSupplier(t *testing.T) {
	var ms MustUint32PtrSliceSupplier = func() []*uint32 {
		return testUint32PtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUint32PtrSliceSupplierResult, v)
}

func TestMustUint32PtrSliceSupplier_ToSilentUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint32PtrSliceSupplier_ToUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint32PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint32PtrSliceSupplierResult, v)
			}
		})
	}
}
