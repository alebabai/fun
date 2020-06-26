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
	testStringSupplierResult string
	testStringSupplierError  = errors.New("error")
)

func testStringSupplier() (string, error) {
	return testStringSupplierResult, nil
}

func testStringSupplierWithError() (string, error) {
	return testStringSupplierResult, testStringSupplierError
}

func TestStringSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSupplier
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name: "with_error",
			s:    testStringSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testStringSupplierError.Error())
			} else {
				r.Equal(testStringSupplierResult, v)
			}
		})
	}
}

func TestStringSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name: "with_error",
			s:    testStringSupplierWithError,
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
				r.EqualError(err, testStringSupplierError.Error())
			} else {
				r.Equal(testStringSupplierResult, v)
			}
		})
	}
}

func TestStringSupplier_ToSilentStringSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name: "with_error",
			s:    testStringSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentStringSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringSupplierResult, v)
			}
		})
	}
}

func TestStringSupplier_ToMustStringSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name: "with_error",
			s:    testStringSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testStringSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testStringSupplierResult, v)
			}
		})
	}
}

func TestSilentStringSupplier(t *testing.T) {
	var ss SilentStringSupplier = func() string {
		return testStringSupplierResult
	}
	v := ss()
	require.Equal(t, testStringSupplierResult, v)
}

func TestMustStringSupplier(t *testing.T) {
	var ms MustStringSupplier = func() string {
		return testStringSupplierResult
	}

	v := ms()
	require.Equal(t, testStringSupplierResult, v)
}

func TestMustStringSupplier_ToSilentStringSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name: "with_error",
			s:    testStringSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentStringSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringSupplierResult, v)
			}
		})
	}
}

func TestMustStringSupplier_ToStringSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringSupplier
	}{
		{
			name: "ok",
			s:    testStringSupplier,
		},
		{
			name: "with_error",
			s:    testStringSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringSupplier()
			r.NotNil(ms)

			s := ms.ToStringSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testStringSupplierError.Error())
			} else {
				r.Equal(testStringSupplierResult, v)
			}
		})
	}
}
