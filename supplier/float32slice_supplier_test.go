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
	resTestFloat32SliceSupplier []float32
	errTestFloat32SliceSupplier = errors.New("error")
)

func testFloat32SliceSupplier() ([]float32, error) {
	return resTestFloat32SliceSupplier, nil
}

func testFloat32SliceSupplierWithError() ([]float32, error) {
	return resTestFloat32SliceSupplier, errTestFloat32SliceSupplier
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
				r.EqualError(err, errTestFloat32SliceSupplier.Error())
			} else {
				r.Equal(resTestFloat32SliceSupplier, v)
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
				r.EqualError(err, errTestFloat32SliceSupplier.Error())
			} else {
				r.Equal(resTestFloat32SliceSupplier, v)
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
				r.Equal(resTestFloat32SliceSupplier, v)
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
				r.PanicsWithError(errTestFloat32SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat32SliceSupplier, v)
			}
		})
	}
}

func TestSilentFloat32SliceSupplier(t *testing.T) {
	var ss SilentFloat32SliceSupplier = func() []float32 {
		return resTestFloat32SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestFloat32SliceSupplier, v)
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
				r.Equal(resTestFloat32SliceSupplier, v)
			}
		})
	}
}

func TestMustFloat32SliceSupplier(t *testing.T) {
	var ms MustFloat32SliceSupplier = func() []float32 {
		return resTestFloat32SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestFloat32SliceSupplier, v)
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
				r.PanicsWithError(errTestFloat32SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat32SliceSupplier, v)
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
				r.Equal(resTestFloat32SliceSupplier, v)
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
				r.EqualError(err, errTestFloat32SliceSupplier.Error())
			} else {
				r.Equal(resTestFloat32SliceSupplier, v)
			}
		})
	}
}
