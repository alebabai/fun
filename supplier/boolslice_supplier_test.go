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
	testBoolSliceSupplierResult []bool
	testBoolSliceSupplierError  = errors.New("error")
)

func testBoolSliceSupplier() ([]bool, error) {
	return testBoolSliceSupplierResult, nil
}

func testBoolSliceSupplierWithError() ([]bool, error) {
	return testBoolSliceSupplierResult, testBoolSliceSupplierError
}

func TestBoolSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSliceSupplier
	}{
		{
			name: "ok",
			s:    testBoolSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBoolSliceSupplierError.Error())
			} else {
				r.Equal(testBoolSliceSupplierResult, v)
			}
		})
	}
}

func TestBoolSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSliceSupplierWithError,
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
				r.EqualError(err, testBoolSliceSupplierError.Error())
			} else {
				r.Equal(testBoolSliceSupplierResult, v)
			}
		})
	}
}

func TestBoolSliceSupplier_ToSilentBoolSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBoolSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolSliceSupplierResult, v)
			}
		})
	}
}

func TestBoolSliceSupplier_ToMustBoolSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBoolSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBoolSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentBoolSliceSupplier(t *testing.T) {
	var ss SilentBoolSliceSupplier = func() []bool {
		return testBoolSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testBoolSliceSupplierResult, v)
}

func TestSilentBoolSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentBoolSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBoolSliceSupplier(t *testing.T) {
	var ms MustBoolSliceSupplier = func() []bool {
		return testBoolSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testBoolSliceSupplierResult, v)
}

func TestMustBoolSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustBoolSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBoolSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBoolSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBoolSliceSupplier_ToSilentBoolSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBoolSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBoolSliceSupplier_ToBoolSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSliceSupplier
	}{
		{
			name: "ok",
			s:    testBoolSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSliceSupplier()
			r.NotNil(ms)

			s := ms.ToBoolSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBoolSliceSupplierError.Error())
			} else {
				r.Equal(testBoolSliceSupplierResult, v)
			}
		})
	}
}
