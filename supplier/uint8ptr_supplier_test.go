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
	testUint8PtrSupplierResult *uint8
	testUint8PtrSupplierError  = errors.New("error")
)

func testUint8PtrSupplier() (*uint8, error) {
	return testUint8PtrSupplierResult, nil
}

func testUint8PtrSupplierWithError() (*uint8, error) {
	return testUint8PtrSupplierResult, testUint8PtrSupplierError
}

func TestUint8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSupplier
	}{
		{
			name: "ok",
			s:    testUint8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint8PtrSupplierError.Error())
			} else {
				r.Equal(testUint8PtrSupplierResult, v)
			}
		})
	}
}

func TestUint8PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSupplierWithError,
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
				r.EqualError(err, testUint8PtrSupplierError.Error())
			} else {
				r.Equal(testUint8PtrSupplierResult, v)
			}
		})
	}
}

func TestUint8PtrSupplier_ToSilentUint8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint8PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint8PtrSupplierResult, v)
			}
		})
	}
}

func TestUint8PtrSupplier_ToMustUint8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint8PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint8PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentUint8PtrSupplier(t *testing.T) {
	var ss SilentUint8PtrSupplier = func() *uint8 {
		return testUint8PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testUint8PtrSupplierResult, v)
}

func TestSilentUint8PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint8PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint8PtrSupplierResult, v)
			}
		})
	}
}

func TestMustUint8PtrSupplier(t *testing.T) {
	var ms MustUint8PtrSupplier = func() *uint8 {
		return testUint8PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testUint8PtrSupplierResult, v)
}

func TestMustUint8PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint8PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint8PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint8PtrSupplierResult, v)
			}
		})
	}
}

func TestMustUint8PtrSupplier_ToSilentUint8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint8PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint8PtrSupplierResult, v)
			}
		})
	}
}

func TestMustUint8PtrSupplier_ToUint8PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint8PtrSupplier
	}{
		{
			name: "ok",
			s:    testUint8PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint8PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint8PtrSupplier()
			r.NotNil(ms)

			s := ms.ToUint8PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint8PtrSupplierError.Error())
			} else {
				r.Equal(testUint8PtrSupplierResult, v)
			}
		})
	}
}
