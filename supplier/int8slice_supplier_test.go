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
	testInt8SliceSupplierResult []int8
	testInt8SliceSupplierError  = errors.New("error")
)

func testInt8SliceSupplier() ([]int8, error) {
	return testInt8SliceSupplierResult, nil
}

func testInt8SliceSupplierWithError() ([]int8, error) {
	return testInt8SliceSupplierResult, testInt8SliceSupplierError
}

func TestInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8SliceSupplier
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt8SliceSupplierError.Error())
			} else {
				r.Equal(testInt8SliceSupplierResult, v)
			}
		})
	}
}

func TestInt8SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8SliceSupplierWithError,
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
				r.EqualError(err, testInt8SliceSupplierError.Error())
			} else {
				r.Equal(testInt8SliceSupplierResult, v)
			}
		})
	}
}

func TestInt8SliceSupplier_ToSilentInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt8SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt8SliceSupplierResult, v)
			}
		})
	}
}

func TestInt8SliceSupplier_ToMustInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8SliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt8SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt8SliceSupplierResult, v)
			}
		})
	}
}

func TestSilentInt8SliceSupplier(t *testing.T) {
	var ss SilentInt8SliceSupplier = func() []int8 {
		return testInt8SliceSupplierResult
	}
	v := ss()
	require.Equal(t, testInt8SliceSupplierResult, v)
}

func TestSilentInt8SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt8SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt8SliceSupplierResult, v)
			}
		})
	}
}

func TestMustInt8SliceSupplier(t *testing.T) {
	var ms MustInt8SliceSupplier = func() []int8 {
		return testInt8SliceSupplierResult
	}

	v := ms()
	require.Equal(t, testInt8SliceSupplierResult, v)
}

func TestMustInt8SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt8SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt8SliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt8SliceSupplierResult, v)
			}
		})
	}
}

func TestMustInt8SliceSupplier_ToSilentInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8SliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8SliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt8SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt8SliceSupplierResult, v)
			}
		})
	}
}

func TestMustInt8SliceSupplier_ToInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int8SliceSupplier
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt8SliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8SliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt8SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt8SliceSupplierError.Error())
			} else {
				r.Equal(testInt8SliceSupplierResult, v)
			}
		})
	}
}
