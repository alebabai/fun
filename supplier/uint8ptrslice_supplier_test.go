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
	resTestUint8PtrSliceSupplier []*uint8
	errTestUint8PtrSliceSupplier = errors.New("error")
)

func testUint8PtrSliceSupplier() ([]*uint8, error) {
	return resTestUint8PtrSliceSupplier, nil
}

func testUint8PtrSliceSupplierWithError() ([]*uint8, error) {
	return resTestUint8PtrSliceSupplier, errTestUint8PtrSliceSupplier
}

func TestUint8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint8PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint8PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint8PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestUint8PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint8PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint8PtrSliceSupplier_ToSilentUint8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint8PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint8PtrSliceSupplier_ToMustUint8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint8PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint8PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentUint8PtrSliceSupplier(t *testing.T) {
	var ss SilentUint8PtrSliceSupplier = func() []*uint8 {
		return resTestUint8PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUint8PtrSliceSupplier, v)
}

func TestSilentUint8PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint8PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint8PtrSliceSupplier(t *testing.T) {
	var ms MustUint8PtrSliceSupplier = func() []*uint8 {
		return resTestUint8PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUint8PtrSliceSupplier, v)
}

func TestMustUint8PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint8PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint8PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint8PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint8PtrSliceSupplier_ToSilentUint8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint8PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint8PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint8PtrSliceSupplier_ToUint8PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint8PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint8PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint8PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint8PtrSliceSupplier, v)
			}
		})
	}
}
