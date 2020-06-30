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
	testUintPtrSupplierResult *uint
	testUintPtrSupplierError  = errors.New("error")
)

func testUintPtrSupplier() (*uint, error) {
	return testUintPtrSupplierResult, nil
}

func testUintPtrSupplierWithError() (*uint, error) {
	return testUintPtrSupplierResult, testUintPtrSupplierError
}

func TestUintPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSupplier
	}{
		{
			name: "ok",
			s:    testUintPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintPtrSupplierError.Error())
			} else {
				r.Equal(testUintPtrSupplierResult, v)
			}
		})
	}
}

func TestUintPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSupplierWithError,
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
				r.EqualError(err, testUintPtrSupplierError.Error())
			} else {
				r.Equal(testUintPtrSupplierResult, v)
			}
		})
	}
}

func TestUintPtrSupplier_ToSilentUintPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUintPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintPtrSupplierResult, v)
			}
		})
	}
}

func TestUintPtrSupplier_ToMustUintPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUintPtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUintPtrSupplierResult, v)
			}
		})
	}
}

func TestSilentUintPtrSupplier(t *testing.T) {
	var ss SilentUintPtrSupplier = func() *uint {
		return testUintPtrSupplierResult
	}
	v := ss()
	require.Equal(t, testUintPtrSupplierResult, v)
}

func TestSilentUintPtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUintPtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintPtrSupplierResult, v)
			}
		})
	}
}

func TestMustUintPtrSupplier(t *testing.T) {
	var ms MustUintPtrSupplier = func() *uint {
		return testUintPtrSupplierResult
	}

	v := ms()
	require.Equal(t, testUintPtrSupplierResult, v)
}

func TestMustUintPtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUintPtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testUintPtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testUintPtrSupplierResult, v)
			}
		})
	}
}

func TestMustUintPtrSupplier_ToSilentUintPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testUintPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUintPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testUintPtrSupplierResult, v)
			}
		})
	}
}

func TestMustUintPtrSupplier_ToUintPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    UintPtrSupplier
	}{
		{
			name: "ok",
			s:    testUintPtrSupplier,
		},
		{
			name: "with_error",
			s:    testUintPtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSupplier()
			r.NotNil(ms)

			s := ms.ToUintPtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testUintPtrSupplierError.Error())
			} else {
				r.Equal(testUintPtrSupplierResult, v)
			}
		})
	}
}
