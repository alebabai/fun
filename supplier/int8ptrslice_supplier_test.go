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
	resTestInt8PtrSliceSupplier []*int8
	errTestInt8PtrSliceSupplier = errors.New("error")
)

func testInt8PtrSliceSupplier() ([]*int8, error) {
	return resTestInt8PtrSliceSupplier, nil
}

func testInt8PtrSliceSupplierWithError() ([]*int8, error) {
	return resTestInt8PtrSliceSupplier, errTestInt8PtrSliceSupplier
}

func TestInt8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testInt8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt8PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestInt8PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt8PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestInt8PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestInt8PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt8PtrSliceSupplier_ToSilentInt8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt8PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt8PtrSliceSupplier_ToMustInt8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt8PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt8PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentInt8PtrSliceSupplier(t *testing.T) {
	var ss SilentInt8PtrSliceSupplier = func() []*int8 {
		return resTestInt8PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestInt8PtrSliceSupplier, v)
}

func TestSilentInt8PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt8PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt8PtrSliceSupplier(t *testing.T) {
	var ms MustInt8PtrSliceSupplier = func() []*int8 {
		return resTestInt8PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestInt8PtrSliceSupplier, v)
}

func TestMustInt8PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt8PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestInt8PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt8PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt8PtrSliceSupplier_ToSilentInt8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt8PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt8PtrSliceSupplier_ToInt8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testInt8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt8PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestInt8PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestInt8PtrSliceSupplier, v)
			}
		})
	}
}
