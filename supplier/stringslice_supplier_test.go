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
	testStringSliceSupplierResult []string
	testStringSliceSupplierError  = errors.New("error")
)

func testStringSliceSupplier() ([]string, error) {
	return testStringSliceSupplierResult, nil
}

func testStringSliceSupplierWithError() ([]string, error) {
	return testStringSliceSupplierResult, testStringSliceSupplierError
}

func TestStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSliceSupplier
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testStringSliceSupplierError.Error())
			} else {
				r.Equal(testStringSliceSupplierResult, v)
			}
		})
	}
}

func TestStringSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringSliceSupplierWithError,
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
				r.EqualError(err, testStringSliceSupplierError.Error())
			} else {
				r.Equal(testStringSliceSupplierResult, v)
			}
		})
	}
}

func TestStringSliceSupplier_ToSilentStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentStringSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringSliceSupplierResult, v)
			}
		})
	}
}

func TestStringSliceSupplier_ToMustStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testStringSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testStringSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentStringSliceSupplier(t *testing.T) {
	var ss SilentStringSliceSupplier = func() []string {
		return testStringSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testStringSliceSupplierResult, v)
}

func TestMustStringSliceSupplier(t *testing.T) {
	var ms MustStringSliceSupplier = func() []string {
		return testStringSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testStringSliceSupplierResult, v)
}

func TestMustStringSliceSupplier_ToSilentStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentStringSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringSliceSupplierResult, v)
			}
		})
	}
}

func TestMustStringSliceSupplier_ToStringSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSliceSupplier
	}{
		{
			name: "ok",
			s:    testStringSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSliceSupplier()
			r.NotNil(ms)

			s := ms.ToStringSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testStringSliceSupplierError.Error())
			} else {
				r.Equal(testStringSliceSupplierResult, v)
			}
		})
	}
}
