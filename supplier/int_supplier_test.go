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
	resTestIntSupplier int
	errTestIntSupplier = errors.New("error")
)

func testIntSupplier() (int, error) {
	return resTestIntSupplier, nil
}

func testIntSupplierWithError() (int, error) {
	return resTestIntSupplier, errTestIntSupplier
}

func TestIntSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSupplier
	}{
		{
			name: "ok",
			s:    testIntSupplier,
		},
		{
			name: "with_error",
			s:    testIntSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestIntSupplier.Error())
			} else {
				r.Equal(resTestIntSupplier, v)
			}
		})
	}
}

func TestIntSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSupplier,
		},
		{
			name: "with_error",
			s:    testIntSupplierWithError,
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
				r.EqualError(err, errTestIntSupplier.Error())
			} else {
				r.Equal(resTestIntSupplier, v)
			}
		})
	}
}

func TestIntSupplier_ToSilentIntSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSupplier,
		},
		{
			name: "with_error",
			s:    testIntSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentIntSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestIntSupplier, v)
			}
		})
	}
}

func TestIntSupplier_ToMustIntSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSupplier,
		},
		{
			name: "with_error",
			s:    testIntSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestIntSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestIntSupplier, v)
			}
		})
	}
}

func TestSilentIntSupplier(t *testing.T) {
	var ss SilentIntSupplier = func() int {
		return resTestIntSupplier
	}
	v := ss()
	require.Equal(t, resTestIntSupplier, v)
}

func TestSilentIntSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSupplier,
		},
		{
			name: "with_error",
			s:    testIntSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentIntSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestIntSupplier, v)
			}
		})
	}
}

func TestMustIntSupplier(t *testing.T) {
	var ms MustIntSupplier = func() int {
		return resTestIntSupplier
	}

	v := ms()
	require.Equal(t, resTestIntSupplier, v)
}

func TestMustIntSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSupplier,
		},
		{
			name: "with_error",
			s:    testIntSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustIntSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestIntSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestIntSupplier, v)
			}
		})
	}
}

func TestMustIntSupplier_ToSilentIntSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testIntSupplier,
		},
		{
			name: "with_error",
			s:    testIntSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentIntSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestIntSupplier, v)
			}
		})
	}
}

func TestMustIntSupplier_ToIntSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    IntSupplier
	}{
		{
			name: "ok",
			s:    testIntSupplier,
		},
		{
			name: "with_error",
			s:    testIntSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSupplier()
			r.NotNil(ms)

			s := ms.ToIntSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestIntSupplier.Error())
			} else {
				r.Equal(resTestIntSupplier, v)
			}
		})
	}
}
