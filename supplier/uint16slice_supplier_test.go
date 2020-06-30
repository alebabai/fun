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
	resTestUint16SliceSupplier []uint16
	errTestUint16SliceSupplier = errors.New("error")
)

func testUint16SliceSupplier() ([]uint16, error) {
	return resTestUint16SliceSupplier, nil
}

func testUint16SliceSupplierWithError() ([]uint16, error) {
	return resTestUint16SliceSupplier, errTestUint16SliceSupplier
}

func TestUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint16SliceSupplier.Error())
			} else {
				r.Equal(resTestUint16SliceSupplier, v)
			}
		})
	}
}

func TestUint16SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
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
				r.EqualError(err, errTestUint16SliceSupplier.Error())
			} else {
				r.Equal(resTestUint16SliceSupplier, v)
			}
		})
	}
}

func TestUint16SliceSupplier_ToSilentUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint16SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16SliceSupplier, v)
			}
		})
	}
}

func TestUint16SliceSupplier_ToMustUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint16SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint16SliceSupplier, v)
			}
		})
	}
}

func TestSilentUint16SliceSupplier(t *testing.T) {
	var ss SilentUint16SliceSupplier = func() []uint16 {
		return resTestUint16SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUint16SliceSupplier, v)
}

func TestSilentUint16SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint16SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16SliceSupplier, v)
			}
		})
	}
}

func TestMustUint16SliceSupplier(t *testing.T) {
	var ms MustUint16SliceSupplier = func() []uint16 {
		return resTestUint16SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUint16SliceSupplier, v)
}

func TestMustUint16SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint16SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint16SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint16SliceSupplier, v)
			}
		})
	}
}

func TestMustUint16SliceSupplier_ToSilentUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint16SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16SliceSupplier, v)
			}
		})
	}
}

func TestMustUint16SliceSupplier_ToUint16SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint16SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16SliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint16SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint16SliceSupplier.Error())
			} else {
				r.Equal(resTestUint16SliceSupplier, v)
			}
		})
	}
}
