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
	testInt64PtrSupplierResult *int64
	testInt64PtrSupplierError  = errors.New("error")
)

func testInt64PtrSupplier() (*int64, error) {
	return testInt64PtrSupplierResult, nil
}

func testInt64PtrSupplierWithError() (*int64, error) {
	return testInt64PtrSupplierResult, testInt64PtrSupplierError
}

func TestInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64PtrSupplier
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt64PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt64PtrSupplierError.Error())
			} else {
				r.Equal(testInt64PtrSupplierResult, v)
			}
		})
	}
}

func TestInt64PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt64PtrSupplierWithError,
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
				r.EqualError(err, testInt64PtrSupplierError.Error())
			} else {
				r.Equal(testInt64PtrSupplierResult, v)
			}
		})
	}
}

func TestInt64PtrSupplier_ToSilentInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt64PtrSupplierResult, v)
			}
		})
	}
}

func TestInt64PtrSupplier_ToMustInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt64PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt64PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentInt64PtrSupplier(t *testing.T) {
	var ss SilentInt64PtrSupplier = func() *int64 {
		return testInt64PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testInt64PtrSupplierResult, v)
}

func TestSilentInt64PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt64PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt64PtrSupplierResult, v)
			}
		})
	}
}

func TestMustInt64PtrSupplier(t *testing.T) {
	var ms MustInt64PtrSupplier = func() *int64 {
		return testInt64PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testInt64PtrSupplierResult, v)
}

func TestMustInt64PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt64PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt64PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt64PtrSupplierResult, v)
			}
		})
	}
}

func TestMustInt64PtrSupplier_ToSilentInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt64PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt64PtrSupplierResult, v)
			}
		})
	}
}

func TestMustInt64PtrSupplier_ToInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int64PtrSupplier
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt64PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSupplier()
			r.NotNil(ms)

			s := ms.ToInt64PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt64PtrSupplierError.Error())
			} else {
				r.Equal(testInt64PtrSupplierResult, v)
			}
		})
	}
}
