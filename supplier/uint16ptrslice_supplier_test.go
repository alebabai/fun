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
	resTestUint16PtrSliceSupplier []*uint16
	errTestUint16PtrSliceSupplier = errors.New("error")
)

func testUint16PtrSliceSupplier() ([]*uint16, error) {
	return resTestUint16PtrSliceSupplier, nil
}

func testUint16PtrSliceSupplierWithError() ([]*uint16, error) {
	return resTestUint16PtrSliceSupplier, errTestUint16PtrSliceSupplier
}

func TestUint16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint16PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint16PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint16PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestUint16PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint16PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint16PtrSliceSupplier_ToSilentUint16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint16PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint16PtrSliceSupplier_ToMustUint16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint16PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint16PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentUint16PtrSliceSupplier(t *testing.T) {
	var ss SilentUint16PtrSliceSupplier = func() []*uint16 {
		return resTestUint16PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUint16PtrSliceSupplier, v)
}

func TestSilentUint16PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint16PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint16PtrSliceSupplier(t *testing.T) {
	var ms MustUint16PtrSliceSupplier = func() []*uint16 {
		return resTestUint16PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUint16PtrSliceSupplier, v)
}

func TestMustUint16PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint16PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint16PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint16PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint16PtrSliceSupplier_ToSilentUint16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint16PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint16PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint16PtrSliceSupplier_ToUint16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint16PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint16PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint16PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint16PtrSliceSupplier, v)
			}
		})
	}
}
