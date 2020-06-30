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
	resTestInt64SliceSupplier []int64
	errTestInt64SliceSupplier = errors.New("error")
)

func testInt64SliceSupplier() ([]int64, error) {
	return resTestInt64SliceSupplier, nil
}

func testInt64SliceSupplierWithError() ([]int64, error) {
	return resTestInt64SliceSupplier, errTestInt64SliceSupplier
}

func TestInt64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64SliceSupplier
	}{
		{
			name: "ok",
			s:    testInt64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt64SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt64SliceSupplier.Error())
			} else {
				r.Equal(resTestInt64SliceSupplier, v)
			}
		})
	}
}

func TestInt64SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt64SliceSupplierWithError,
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
				r.EqualError(err, errTestInt64SliceSupplier.Error())
			} else {
				r.Equal(resTestInt64SliceSupplier, v)
			}
		})
	}
}

func TestInt64SliceSupplier_ToSilentInt64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt64SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt64SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64SliceSupplier, v)
			}
		})
	}
}

func TestInt64SliceSupplier_ToMustInt64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt64SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt64SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt64SliceSupplier, v)
			}
		})
	}
}

func TestSilentInt64SliceSupplier(t *testing.T) {
	var ss SilentInt64SliceSupplier = func() []int64 {
		return resTestInt64SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestInt64SliceSupplier, v)
}

func TestSilentInt64SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt64SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt64SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64SliceSupplier, v)
			}
		})
	}
}

func TestMustInt64SliceSupplier(t *testing.T) {
	var ms MustInt64SliceSupplier = func() []int64 {
		return resTestInt64SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestInt64SliceSupplier, v)
}

func TestMustInt64SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt64SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt64SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt64SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt64SliceSupplier, v)
			}
		})
	}
}

func TestMustInt64SliceSupplier_ToSilentInt64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt64SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt64SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64SliceSupplier, v)
			}
		})
	}
}

func TestMustInt64SliceSupplier_ToInt64SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64SliceSupplier
	}{
		{
			name: "ok",
			s:    testInt64SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt64SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64SliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt64SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt64SliceSupplier.Error())
			} else {
				r.Equal(resTestInt64SliceSupplier, v)
			}
		})
	}
}
