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
	testUint32SliceSupplierResult []uint32
	testUint32SliceSupplierError  = errors.New("error")
)

func testUint32SliceSupplier() ([]uint32, error) {
	return testUint32SliceSupplierResult, nil
}

func testUint32SliceSupplierWithError() ([]uint32, error) {
	return testUint32SliceSupplierResult, testUint32SliceSupplierError
}

func TestUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint32SliceSupplierError.Error())
			} else {
				r.Equal(testUint32SliceSupplierResult, v)
			}
		})
	}
}

func TestUint32SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32SliceSupplierWithError,
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
				r.EqualError(err, testUint32SliceSupplierError.Error())
			} else {
				r.Equal(testUint32SliceSupplierResult, v)
			}
		})
	}
}

func TestUint32SliceSupplier_ToSilentUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint32SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32SliceSupplierResult, v)
			}
		})
	}
}

func TestUint32SliceSupplier_ToMustUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint32SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint32SliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUint32SliceSupplier(t *testing.T) {
	var ss SilentUint32SliceSupplier = func() []uint32 {
		return testUint32SliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUint32SliceSupplierResult, v)
}

func TestSilentUint32SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint32SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32SliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint32SliceSupplier(t *testing.T) {
	var ms MustUint32SliceSupplier = func() []uint32 {
		return testUint32SliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUint32SliceSupplierResult, v)
}

func TestMustUint32SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint32SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint32SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint32SliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint32SliceSupplier_ToSilentUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint32SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint32SliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint32SliceSupplier_ToUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint32SliceSupplier
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint32SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32SliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint32SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint32SliceSupplierError.Error())
			} else {
				r.Equal(testUint32SliceSupplierResult, v)
			}
		})
	}
}
