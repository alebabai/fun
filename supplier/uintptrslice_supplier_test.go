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
	testUintPtrSliceSupplierResult []*uint
	testUintPtrSliceSupplierError  = errors.New("error")
)

func testUintPtrSliceSupplier() ([]*uint, error) {
	return testUintPtrSliceSupplierResult, nil
}

func testUintPtrSliceSupplierWithError() ([]*uint, error) {
	return testUintPtrSliceSupplierResult, testUintPtrSliceSupplierError
}

func TestUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintPtrSliceSupplierError.Error())
			} else {
				r.Equal(testUintPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUintPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSliceSupplierWithError,
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
				r.EqualError(err, testUintPtrSliceSupplierError.Error())
			} else {
				r.Equal(testUintPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUintPtrSliceSupplier_ToSilentUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUintPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestUintPtrSliceSupplier_ToMustUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUintPtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUintPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentUintPtrSliceSupplier(t *testing.T) {
	var ss SilentUintPtrSliceSupplier = func() []*uint {
		return testUintPtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testUintPtrSliceSupplierResult, v)
}

func TestSilentUintPtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUintPtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUintPtrSliceSupplier(t *testing.T) {
	var ms MustUintPtrSliceSupplier = func() []*uint {
		return testUintPtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testUintPtrSliceSupplierResult, v)
}

func TestMustUintPtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUintPtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUintPtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUintPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUintPtrSliceSupplier_ToSilentUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUintPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustUintPtrSliceSupplier_ToUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUintPtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintPtrSliceSupplierError.Error())
			} else {
				r.Equal(testUintPtrSliceSupplierResult, v)
			}
		})
	}
}
