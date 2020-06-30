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
	testFloat32SliceSupplierResult []float32
	testFloat32SliceSupplierError  = errors.New("error")
)

func testFloat32SliceSupplier() ([]float32, error) {
	return testFloat32SliceSupplierResult, nil
}

func testFloat32SliceSupplierWithError() ([]float32, error) {
	return testFloat32SliceSupplierResult, testFloat32SliceSupplierError
}

func TestFloat32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32SliceSupplier
	}{
		{
			name: "ok",
			s:    testFloat32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat32SliceSupplierError.Error())
			} else {
				r.Equal(testFloat32SliceSupplierResult, v)
			}
		})
	}
}

func TestFloat32SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32SliceSupplierWithError,
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
				r.EqualError(err, testFloat32SliceSupplierError.Error())
			} else {
				r.Equal(testFloat32SliceSupplierResult, v)
			}
		})
	}
}

func TestFloat32SliceSupplier_ToSilentFloat32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat32SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32SliceSupplierResult, v)
			}
		})
	}
}

func TestFloat32SliceSupplier_ToMustFloat32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat32SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat32SliceSupplierResult, v)
			}
		})
	}
}

func TestSilentFloat32SliceSupplier(t *testing.T) {
	var ss SilentFloat32SliceSupplier = func() []float32 {
		return testFloat32SliceSupplierResult
	}
	v := ss()
	require.Equal(t, testFloat32SliceSupplierResult, v)
}

func TestSilentFloat32SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat32SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32SliceSupplierResult, v)
			}
		})
	}
}

func TestMustFloat32SliceSupplier(t *testing.T) {
	var ms MustFloat32SliceSupplier = func() []float32 {
		return testFloat32SliceSupplierResult
	}

	v := ms()
	require.Equal(t, testFloat32SliceSupplierResult, v)
}

func TestMustFloat32SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat32SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testFloat32SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testFloat32SliceSupplierResult, v)
			}
		})
	}
}

func TestMustFloat32SliceSupplier_ToSilentFloat32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testFloat32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat32SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testFloat32SliceSupplierResult, v)
			}
		})
	}
}

func TestMustFloat32SliceSupplier_ToFloat32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Float32SliceSupplier
	}{
		{
			name: "ok",
			s:    testFloat32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testFloat32SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat32SliceSupplier()
			r.NotNil(ms)

			s := ms.ToFloat32SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testFloat32SliceSupplierError.Error())
			} else {
				r.Equal(testFloat32SliceSupplierResult, v)
			}
		})
	}
}
