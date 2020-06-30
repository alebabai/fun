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
	testInt32PtrSliceSupplierResult []*int32
	testInt32PtrSliceSupplierError  = errors.New("error")
)

func testInt32PtrSliceSupplier() ([]*int32, error) {
	return testInt32PtrSliceSupplierResult, nil
}

func testInt32PtrSliceSupplierWithError() ([]*int32, error) {
	return testInt32PtrSliceSupplierResult, testInt32PtrSliceSupplierError
}

func TestInt32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testInt32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testInt32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestInt32PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSliceSupplierWithError,
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
				r.EqualError(err, testInt32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testInt32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestInt32PtrSliceSupplier_ToSilentInt32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestInt32PtrSliceSupplier_ToMustInt32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt32PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentInt32PtrSliceSupplier(t *testing.T) {
	var ss SilentInt32PtrSliceSupplier = func() []*int32 {
		return testInt32PtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testInt32PtrSliceSupplierResult, v)
}

func TestSilentInt32PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt32PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustInt32PtrSliceSupplier(t *testing.T) {
	var ms MustInt32PtrSliceSupplier = func() []*int32 {
		return testInt32PtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testInt32PtrSliceSupplierResult, v)
}

func TestMustInt32PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt32PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testInt32PtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testInt32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustInt32PtrSliceSupplier_ToSilentInt32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt32PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testInt32PtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustInt32PtrSliceSupplier_ToInt32PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Int32PtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testInt32PtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testInt32PtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt32PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testInt32PtrSliceSupplierError.Error())
			} else {
				r.Equal(testInt32PtrSliceSupplierResult, v)
			}
		})
	}
}
