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
	resTestUint32PtrSliceSupplier []*uint32
	errTestUint32PtrSliceSupplier = errors.New("error")
)

func testUint32PtrSliceSupplier() ([]*uint32, error) {
	return resTestUint32PtrSliceSupplier, nil
}

func testUint32PtrSliceSupplierWithError() ([]*uint32, error) {
	return resTestUint32PtrSliceSupplier, errTestUint32PtrSliceSupplier
}

func TestUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint32PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint32PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint32PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestUint32PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint32PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint32PtrSliceSupplier_ToSilentUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint32PtrSliceSupplier_ToMustUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint32PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint32PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentUint32PtrSliceSupplier(t *testing.T) {
	var ss SilentUint32PtrSliceSupplier = func() []*uint32 {
		return resTestUint32PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUint32PtrSliceSupplier, v)
}

func TestSilentUint32PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint32PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint32PtrSliceSupplier(t *testing.T) {
	var ms MustUint32PtrSliceSupplier = func() []*uint32 {
		return resTestUint32PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUint32PtrSliceSupplier, v)
}

func TestMustUint32PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint32PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint32PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint32PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint32PtrSliceSupplier_ToSilentUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint32PtrSliceSupplier_ToUint32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint32PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint32PtrSliceSupplier.Error())
			} else {
				r.Equal(resTestUint32PtrSliceSupplier, v)
			}
		})
	}
}
