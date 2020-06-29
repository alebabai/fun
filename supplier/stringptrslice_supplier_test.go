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
	testStringPtrSliceSupplierResult []*string
	testStringPtrSliceSupplierError  = errors.New("error")
)

func testStringPtrSliceSupplier() ([]*string, error) {
	return testStringPtrSliceSupplierResult, nil
}

func testStringPtrSliceSupplierWithError() ([]*string, error) {
	return testStringPtrSliceSupplierResult, testStringPtrSliceSupplierError
}

func TestStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testStringPtrSliceSupplierError.Error())
			} else {
				r.Equal(testStringPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestStringPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSliceSupplierWithError,
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
				r.EqualError(err, testStringPtrSliceSupplierError.Error())
			} else {
				r.Equal(testStringPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestStringPtrSliceSupplier_ToSilentStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentStringPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestStringPtrSliceSupplier_ToMustStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testStringPtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testStringPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentStringPtrSliceSupplier(t *testing.T) {
	var ss SilentStringPtrSliceSupplier = func() []*string {
		return testStringPtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testStringPtrSliceSupplierResult, v)
}

func TestMustStringPtrSliceSupplier(t *testing.T) {
	var ms MustStringPtrSliceSupplier = func() []*string {
		return testStringPtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testStringPtrSliceSupplierResult, v)
}

func TestMustStringPtrSliceSupplier_ToSilentStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentStringPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringPtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustStringPtrSliceSupplier_ToStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToStringPtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testStringPtrSliceSupplierError.Error())
			} else {
				r.Equal(testStringPtrSliceSupplierResult, v)
			}
		})
	}
}
