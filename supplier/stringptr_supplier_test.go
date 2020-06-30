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
	resTestStringPtrSupplier *string
	errTestStringPtrSupplier = errors.New("error")
)

func testStringPtrSupplier() (*string, error) {
	return resTestStringPtrSupplier, nil
}

func testStringPtrSupplierWithError() (*string, error) {
	return resTestStringPtrSupplier, errTestStringPtrSupplier
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
				r.EqualError(err, errTestStringPtrSupplier.Error())
			} else {
				r.Equal(resTestStringPtrSupplier, v)
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
				r.EqualError(err, errTestStringPtrSupplier.Error())
			} else {
				r.Equal(resTestStringPtrSupplier, v)
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
				r.Equal(resTestStringPtrSupplier, v)
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
				r.PanicsWithError(errTestStringPtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestStringPtrSupplier, v)
			}
		})
	}
}

func TestSilentStringPtrSupplier(t *testing.T) {
	var ss SilentStringPtrSupplier = func() *string {
		return resTestStringPtrSupplier
	}
	v := ss()
	require.Equal(t, resTestStringPtrSupplier, v)
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
				r.Equal(resTestStringPtrSupplier, v)
			}
		})
	}
}

func TestMustStringPtrSupplier(t *testing.T) {
	var ms MustStringPtrSupplier = func() *string {
		return resTestStringPtrSupplier
	}

	v := ms()
	require.Equal(t, resTestStringPtrSupplier, v)
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
				r.PanicsWithError(errTestStringPtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestStringPtrSupplier, v)
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
				r.Equal(resTestStringPtrSupplier, v)
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
				r.EqualError(err, errTestStringPtrSupplier.Error())
			} else {
				r.Equal(resTestStringPtrSupplier, v)
			}
		})
	}
}
