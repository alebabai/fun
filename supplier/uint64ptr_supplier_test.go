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
	resTestUint64PtrSupplier *uint64
	errTestUint64PtrSupplier = errors.New("error")
)

func testUint64PtrSupplier() (*uint64, error) {
	return resTestUint64PtrSupplier, nil
}

func testUint64PtrSupplierWithError() (*uint64, error) {
	return resTestUint64PtrSupplier, errTestUint64PtrSupplier
}

func TestUint64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSupplier
	}{
		{
			name: "ok",
			s:    testUint64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint64PtrSupplier.Error())
			} else {
				r.Equal(resTestUint64PtrSupplier, v)
			}
		})
	}
}

func TestUint64PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSupplierWithError,
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
				r.EqualError(err, errTestUint64PtrSupplier.Error())
			} else {
				r.Equal(resTestUint64PtrSupplier, v)
			}
		})
	}
}

func TestUint64PtrSupplier_ToSilentUint64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSupplier, v)
			}
		})
	}
}

func TestUint64PtrSupplier_ToMustUint64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint64PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint64PtrSupplier, v)
			}
		})
	}
}

func TestSilentUint64PtrSupplier(t *testing.T) {
	var ss SilentUint64PtrSupplier = func() *uint64 {
		return resTestUint64PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestUint64PtrSupplier, v)
}

func TestSilentUint64PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint64PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSupplier, v)
			}
		})
	}
}

func TestMustUint64PtrSupplier(t *testing.T) {
	var ms MustUint64PtrSupplier = func() *uint64 {
		return resTestUint64PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestUint64PtrSupplier, v)
}

func TestMustUint64PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint64PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestUint64PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint64PtrSupplier, v)
			}
		})
	}
}

func TestMustUint64PtrSupplier_ToSilentUint64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSupplier, v)
			}
		})
	}
}

func TestMustUint64PtrSupplier_ToUint64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSupplier
	}{
		{
			name: "ok",
			s:    testUint64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSupplier()
			r.NotNil(ms)

			s := ms.ToUint64PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestUint64PtrSupplier.Error())
			} else {
				r.Equal(resTestUint64PtrSupplier, v)
			}
		})
	}
}
