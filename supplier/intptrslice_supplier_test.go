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
	testIntPtrSliceSupplierResult []*int
	testIntPtrSliceSupplierError  = errors.New("error")
)

func testIntPtrSliceSupplier() ([]*int, error) {
	return testIntPtrSliceSupplierResult, nil
}

func testIntPtrSliceSupplierWithError() ([]*int, error) {
	return testIntPtrSliceSupplierResult, testIntPtrSliceSupplierError
}

func TestIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testIntPtrSliceSupplierError.Error())
			} else {
				r.Equal(testIntPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestIntPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSliceSupplierWithError,
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
				r.EqualError(err, testIntPtrSliceSupplierError.Error())
			} else {
				r.Equal(testIntPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestIntPtrSliceSupplier_ToSilentIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentIntPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testIntPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestIntPtrSliceSupplier_ToMustIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testIntPtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testIntPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentIntPtrSliceSupplier(t *testing.T) {
	var ss SilentIntPtrSliceSupplier = func() []*int {
		return testIntPtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testIntPtrSliceSupplierResult, v)
}

func TestSilentIntPtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentIntPtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testIntPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustIntPtrSliceSupplier(t *testing.T) {
	var ms MustIntPtrSliceSupplier = func() []*int {
		return testIntPtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testIntPtrSliceSupplierResult, v)
}

func TestMustIntPtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustIntPtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testIntPtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testIntPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustIntPtrSliceSupplier_ToSilentIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentIntPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testIntPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustIntPtrSliceSupplier_ToIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testIntPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToIntPtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testIntPtrSliceSupplierError.Error())
			} else {
				r.Equal(testIntPtrSliceSupplierResult, v)
			}
		})
	}
}
