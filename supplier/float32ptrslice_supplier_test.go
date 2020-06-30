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
	testFloat32PtrSliceSupplierResult []*float32
	testFloat32PtrSliceSupplierError  = errors.New("error")
)

func testFloat32PtrSliceSupplier() ([]*float32, error) {
	return testFloat32PtrSliceSupplierResult, nil
}

func testFloat32PtrSliceSupplierWithError() ([]*float32, error) {
	return testFloat32PtrSliceSupplierResult, testFloat32PtrSliceSupplierError
}

func TestFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testFloat32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestFloat32PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSliceSupplierWithError,
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
				r.EqualError(err, testFloat32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testFloat32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestFloat32PtrSliceSupplier_ToSilentFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestFloat32PtrSliceSupplier_ToMustFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat32PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentFloat32PtrSliceSupplier(t *testing.T) {
	var ss SilentFloat32PtrSliceSupplier = func() []*float32 {
		return testFloat32PtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testFloat32PtrSliceSupplierResult, v)
}

func TestSilentFloat32PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat32PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustFloat32PtrSliceSupplier(t *testing.T) {
	var ms MustFloat32PtrSliceSupplier = func() []*float32 {
		return testFloat32PtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testFloat32PtrSliceSupplierResult, v)
}

func TestMustFloat32PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat32PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat32PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustFloat32PtrSliceSupplier_ToSilentFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustFloat32PtrSliceSupplier_ToFloat32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testFloat32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToFloat32PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testFloat32PtrSliceSupplierResult, v)
			}
		})
	}
}
