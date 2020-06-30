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
	testInt16PtrSupplierResult *int16
	testInt16PtrSupplierError  = errors.New("error")
)

func testInt16PtrSupplier() (*int16, error) {
	return testInt16PtrSupplierResult, nil
}

func testInt16PtrSupplierWithError() (*int16, error) {
	return testInt16PtrSupplierResult, testInt16PtrSupplierError
}

func TestInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSupplier
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt16PtrSupplierError.Error())
			} else {
				r.Equal(testInt16PtrSupplierResult, v)
			}
		})
	}
}

func TestInt16PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSupplierWithError,
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
				r.EqualError(err, testInt16PtrSupplierError.Error())
			} else {
				r.Equal(testInt16PtrSupplierResult, v)
			}
		})
	}
}

func TestInt16PtrSupplier_ToSilentInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt16PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt16PtrSupplierResult, v)
			}
		})
	}
}

func TestInt16PtrSupplier_ToMustInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt16PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt16PtrSupplierResult, v)
			}
		})
	}
}

func TestSilentInt16PtrSupplier(t *testing.T) {
	var ss SilentInt16PtrSupplier = func() *int16 {
		return testInt16PtrSupplierResult
	}
	v := ss()
	require.Equal(t, testInt16PtrSupplierResult, v)
}

func TestSilentInt16PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt16PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt16PtrSupplierResult, v)
			}
		})
	}
}

func TestMustInt16PtrSupplier(t *testing.T) {
	var ms MustInt16PtrSupplier = func() *int16 {
		return testInt16PtrSupplierResult
	}

	v := ms()
	require.Equal(t, testInt16PtrSupplierResult, v)
}

func TestMustInt16PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt16PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt16PtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt16PtrSupplierResult, v)
			}
		})
	}
}

func TestMustInt16PtrSupplier_ToSilentInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt16PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt16PtrSupplierResult, v)
			}
		})
	}
}

func TestMustInt16PtrSupplier_ToInt16PtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int16PtrSupplier
	}{
		{
			name: "ok",
			s:    testInt16PtrSupplier,
		},
		{
			name: "with_error",
			s:    testInt16PtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSupplier()
			r.NotNil(ms)

			s := ms.ToInt16PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt16PtrSupplierError.Error())
			} else {
				r.Equal(testInt16PtrSupplierResult, v)
			}
		})
	}
}
