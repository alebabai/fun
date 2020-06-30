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
	resTestUint8SliceSupplier []uint8
	errTestUint8SliceSupplier = errors.New("error")
)

func testUint8SliceSupplier() ([]uint8, error) {
	return resTestUint8SliceSupplier, nil
}

func testUint8SliceSupplierWithError() ([]uint8, error) {
	return resTestUint8SliceSupplier, errTestUint8SliceSupplier
}

func TestUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint8SliceSupplier.Error())
			} else {
				r.Equal(resTestUint8SliceSupplier, v)
			}
		})
	}
}

func TestUint8SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
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
				r.EqualError(err, errTestUint8SliceSupplier.Error())
			} else {
				r.Equal(resTestUint8SliceSupplier, v)
			}
		})
	}
}

func TestUint8SliceSupplier_ToSilentUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint8SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8SliceSupplier, v)
			}
		})
	}
}

func TestUint8SliceSupplier_ToMustUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint8SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint8SliceSupplier, v)
			}
		})
	}
}

func TestSilentUint8SliceSupplier(t *testing.T) {
	var ss SilentUint8SliceSupplier = func() []uint8 {
		return resTestUint8SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUint8SliceSupplier, v)
}

func TestSilentUint8SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint8SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8SliceSupplier, v)
			}
		})
	}
}

func TestMustUint8SliceSupplier(t *testing.T) {
	var ms MustUint8SliceSupplier = func() []uint8 {
		return resTestUint8SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUint8SliceSupplier, v)
}

func TestMustUint8SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint8SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint8SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint8SliceSupplier, v)
			}
		})
	}
}

func TestMustUint8SliceSupplier_ToSilentUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint8SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8SliceSupplier, v)
			}
		})
	}
}

func TestMustUint8SliceSupplier_ToUint8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8SliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint8SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint8SliceSupplier.Error())
			} else {
				r.Equal(resTestUint8SliceSupplier, v)
			}
		})
	}
}
