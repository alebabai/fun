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
	testBoolPtrSliceSupplierResult []*bool
	testBoolPtrSliceSupplierError  = errors.New("error")
)

func testBoolPtrSliceSupplier() ([]*bool, error) {
	return testBoolPtrSliceSupplierResult, nil
}

func testBoolPtrSliceSupplierWithError() ([]*bool, error) {
	return testBoolPtrSliceSupplierResult, testBoolPtrSliceSupplierError
}

func TestBoolPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testBoolPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBoolPtrSliceSupplierError.Error())
			} else {
				r.Equal(testBoolPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestBoolPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolPtrSliceSupplierWithError,
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
				r.EqualError(err, testBoolPtrSliceSupplierError.Error())
			} else {
				r.Equal(testBoolPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestBoolPtrSliceSupplier_ToSilentBoolPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBoolPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestBoolPtrSliceSupplier_ToMustBoolPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolPtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBoolPtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBoolPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentBoolPtrSliceSupplier(t *testing.T) {
	var ss SilentBoolPtrSliceSupplier = func() []*bool {
		return testBoolPtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testBoolPtrSliceSupplierResult, v)
}

func TestSilentBoolPtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentBoolPtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBoolPtrSliceSupplier(t *testing.T) {
	var ms MustBoolPtrSliceSupplier = func() []*bool {
		return testBoolPtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testBoolPtrSliceSupplierResult, v)
}

func TestMustBoolPtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustBoolPtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBoolPtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBoolPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBoolPtrSliceSupplier_ToSilentBoolPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolPtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBoolPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBoolPtrSliceSupplier_ToBoolPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testBoolPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolPtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToBoolPtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBoolPtrSliceSupplierError.Error())
			} else {
				r.Equal(testBoolPtrSliceSupplierResult, v)
			}
		})
	}
}
