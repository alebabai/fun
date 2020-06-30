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
	resTestInt16SliceSupplier []int16
	errTestInt16SliceSupplier = errors.New("error")
)

func testInt16SliceSupplier() ([]int16, error) {
	return resTestInt16SliceSupplier, nil
}

func testInt16SliceSupplierWithError() ([]int16, error) {
	return resTestInt16SliceSupplier, errTestInt16SliceSupplier
}

func TestInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt16SliceSupplier.Error())
			} else {
				r.Equal(resTestInt16SliceSupplier, v)
			}
		})
	}
}

func TestInt16SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
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
				r.EqualError(err, errTestInt16SliceSupplier.Error())
			} else {
				r.Equal(resTestInt16SliceSupplier, v)
			}
		})
	}
}

func TestInt16SliceSupplier_ToSilentInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt16SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16SliceSupplier, v)
			}
		})
	}
}

func TestInt16SliceSupplier_ToMustInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt16SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt16SliceSupplier, v)
			}
		})
	}
}

func TestSilentInt16SliceSupplier(t *testing.T) {
	var ss SilentInt16SliceSupplier = func() []int16 {
		return resTestInt16SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestInt16SliceSupplier, v)
}

func TestSilentInt16SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt16SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16SliceSupplier, v)
			}
		})
	}
}

func TestMustInt16SliceSupplier(t *testing.T) {
	var ms MustInt16SliceSupplier = func() []int16 {
		return resTestInt16SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestInt16SliceSupplier, v)
}

func TestMustInt16SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt16SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt16SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt16SliceSupplier, v)
			}
		})
	}
}

func TestMustInt16SliceSupplier_ToSilentInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt16SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16SliceSupplier, v)
			}
		})
	}
}

func TestMustInt16SliceSupplier_ToInt16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16SliceSupplier
	}{
		{
			name: "ok",
			s:    testInt16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt16SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16SliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt16SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt16SliceSupplier.Error())
			} else {
				r.Equal(resTestInt16SliceSupplier, v)
			}
		})
	}
}
