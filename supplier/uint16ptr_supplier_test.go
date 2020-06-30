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
	testUint16PtrSupplierResult *uint16
	testUint16PtrSupplierError  = errors.New("error")
)

func testUint16PtrSupplier() (*uint16, error) {
	return testUint16PtrSupplierResult, nil
}

func testUint16PtrSupplierWithError() (*uint16, error) {
	return testUint16PtrSupplierResult, testUint16PtrSupplierError
}

func TestUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSupplier
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint16PtrSupplierError.Error())
			} else {
				r.Equal(testUint16PtrSupplierResult, v)
			}
		})
	}
}

func TestUint16PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSupplierWithError,
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
				r.EqualError(err, testUint16PtrSupplierError.Error())
			} else {
				r.Equal(testUint16PtrSupplierResult, v)
			}
		})
	}
}

func TestUint16PtrSupplier_ToSilentUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint16PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint16PtrSupplierResult, v)
			}
		})
	}
}

func TestUint16PtrSupplier_ToMustUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint16PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint16PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentUint16PtrSupplier(t *testing.T) {
	var ss SilentUint16PtrSupplier = func() *uint16 {
		return testUint16PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testUint16PtrSupplierResult, v)
}

func TestSilentUint16PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint16PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint16PtrSupplierResult, v)
			}
		})
	}
}

func TestMustUint16PtrSupplier(t *testing.T) {
	var ms MustUint16PtrSupplier = func() *uint16 {
		return testUint16PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testUint16PtrSupplierResult, v)
}

func TestMustUint16PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint16PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint16PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint16PtrSupplierResult, v)
			}
		})
	}
}

func TestMustUint16PtrSupplier_ToSilentUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint16PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint16PtrSupplierResult, v)
			}
		})
	}
}

func TestMustUint16PtrSupplier_ToUint16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint16PtrSupplier
	}{
		{
			name: "ok",
			s:    testUint16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testUint16PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint16PtrSupplier()
			r.NotNil(ms)

			s := ms.ToUint16PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint16PtrSupplierError.Error())
			} else {
				r.Equal(testUint16PtrSupplierResult, v)
			}
		})
	}
}
