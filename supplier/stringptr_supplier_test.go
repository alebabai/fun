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
	testStringPtrSupplierResult *string
	testStringPtrSupplierError  = errors.New("error")
)

func testStringPtrSupplier() (*string, error) {
	return testStringPtrSupplierResult, nil
}

func testStringPtrSupplierWithError() (*string, error) {
	return testStringPtrSupplierResult, testStringPtrSupplierError
}

func TestStringPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSupplier
	}{
		{
			name: "ok",
			s:    testStringPtrSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testStringPtrSupplierError.Error())
			} else {
				r.Equal(testStringPtrSupplierResult, v)
			}
		})
	}
}

func TestStringPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSupplierWithError,
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
				r.EqualError(err, testStringPtrSupplierError.Error())
			} else {
				r.Equal(testStringPtrSupplierResult, v)
			}
		})
	}
}

func TestStringPtrSupplier_ToSilentStringPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentStringPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringPtrSupplierResult, v)
			}
		})
	}
}

func TestStringPtrSupplier_ToMustStringPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testStringPtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testStringPtrSupplierResult, v)
			}
		})
	}
}

func TestSilentStringPtrSupplier(t *testing.T) {
	var ss SilentStringPtrSupplier = func() *string {
		return testStringPtrSupplierResult
	}
	v := ss()
	require.Equal(t, testStringPtrSupplierResult, v)
}

func TestSilentStringPtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentStringPtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringPtrSupplierResult, v)
			}
		})
	}
}

func TestMustStringPtrSupplier(t *testing.T) {
	var ms MustStringPtrSupplier = func() *string {
		return testStringPtrSupplierResult
	}

	v := ms()
	require.Equal(t, testStringPtrSupplierResult, v)
}

func TestMustStringPtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustStringPtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testStringPtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testStringPtrSupplierResult, v)
			}
		})
	}
}

func TestMustStringPtrSupplier_ToSilentStringPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testStringPtrSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentStringPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testStringPtrSupplierResult, v)
			}
		})
	}
}

func TestMustStringPtrSupplier_ToStringPtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    StringPtrSupplier
	}{
		{
			name: "ok",
			s:    testStringPtrSupplier,
		},
		{
			name: "with_error",
			s:    testStringPtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSupplier()
			r.NotNil(ms)

			s := ms.ToStringPtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testStringPtrSupplierError.Error())
			} else {
				r.Equal(testStringPtrSupplierResult, v)
			}
		})
	}
}
