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
	resTestSupplier interface{}
	errTestSupplier = errors.New("error")
)

func testSupplier() (interface{}, error) {
	return resTestSupplier, nil
}

func testSupplierWithError() (interface{}, error) {
	return resTestSupplier, errTestSupplier
}

func TestSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name: "with_error",
			s:    testSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestSupplier.Error())
			} else {
				r.Equal(resTestSupplier, v)
			}
		})
	}
}

func TestSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name: "with_error",
			s:    testSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestSupplier, v)
			}
		})
	}
}

func TestSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name: "with_error",
			s:    testSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestSupplier, v)
			}
		})
	}
}

func TestSilentSupplier(t *testing.T) {
	var ss SilentSupplier = func() interface{} {
		return resTestSupplier
	}
	v := ss()
	require.Equal(t, resTestSupplier, v)
}

func TestMustSupplier(t *testing.T) {
	var ms MustSupplier = func() interface{} {
		return resTestSupplier
	}

	v := ms()
	require.Equal(t, resTestSupplier, v)
}

func TestMustSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
		err  bool
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name: "with_error",
			s:    testSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestSupplier, v)
			}
		})
	}
}

func TestMustSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name: "with_error",
			s:    testSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustSupplier()
			r.NotNil(ms)

			s := ms.ToSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestSupplier.Error())
			} else {
				r.Equal(resTestSupplier, v)
			}
		})
	}
}
