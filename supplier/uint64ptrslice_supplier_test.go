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
	testUint64PtrSliceSupplierResult []*uint64
	testUint64PtrSliceSupplierError  = errors.New("error")
)

func testUint64PtrSliceSupplier() ([]*uint64, error) {
	return testUint64PtrSliceSupplierResult, nil
}

func testUint64PtrSliceSupplierWithError() ([]*uint64, error) {
	return testUint64PtrSliceSupplierResult, testUint64PtrSliceSupplierError
}

func TestUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint64PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUint64PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSliceSupplierWithError,
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
				r.EqualError(err, testUint64PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUint64PtrSliceSupplier_ToSilentUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUint64PtrSliceSupplier_ToMustUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint64PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUint64PtrSliceSupplier(t *testing.T) {
	var ss SilentUint64PtrSliceSupplier = func() []*uint64 {
		return testUint64PtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUint64PtrSliceSupplierResult, v)
}

func TestSilentUint64PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint64PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint64PtrSliceSupplier(t *testing.T) {
	var ms MustUint64PtrSliceSupplier = func() []*uint64 {
		return testUint64PtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUint64PtrSliceSupplierResult, v)
}

func TestMustUint64PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint64PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUint64PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUint64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint64PtrSliceSupplier_ToSilentUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUint64PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUint64PtrSliceSupplier_ToUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Uint64PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUint64PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint64PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUint64PtrSliceSupplierError.Error())
			} else {
				r.Equal(testUint64PtrSliceSupplierResult, v)
			}
		})
	}
}
