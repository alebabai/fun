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
	testUint16PtrSliceSupplierResult []*uint16
	testUint16PtrSliceSupplierError  = errors.New("error")
)

func testUint16PtrSliceSupplier() ([]*uint16, error) {
	return testUint16PtrSliceSupplierResult, nil
}

func testUint16PtrSliceSupplierWithError() ([]*uint16, error) {
	return testUint16PtrSliceSupplierResult, testUint16PtrSliceSupplierError
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
				r.EqualError(err, testUint16PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint16PtrSliceSupplierResult, v)
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
				r.EqualError(err, testUint16PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint16PtrSliceSupplierResult, v)
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
				r.Equal(testUint16PtrSliceSupplierResult, v)
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
				r.PanicsWithError(testUint16PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint16PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUint16PtrSliceSupplier(t *testing.T) {
	var ss SilentUint16PtrSliceSupplier = func() []*uint16 {
		return testUint16PtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUint16PtrSliceSupplierResult, v)
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
				r.Equal(testUint16PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint16PtrSliceSupplier(t *testing.T) {
	var ms MustUint16PtrSliceSupplier = func() []*uint16 {
		return testUint16PtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUint16PtrSliceSupplierResult, v)
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
				r.PanicsWithError(testUint16PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint16PtrSliceSupplierResult, v)
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
				r.Equal(testUint16PtrSliceSupplierResult, v)
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
				r.EqualError(err, testUint16PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint16PtrSliceSupplierResult, v)
			}
		})
	}
}
