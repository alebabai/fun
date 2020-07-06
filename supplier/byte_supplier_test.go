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
	resTestByteSupplier byte
	errTestByteSupplier = errors.New("error")
)

func testByteSupplier() (byte, error) {
	return resTestByteSupplier, nil
}

func testByteSupplierWithError() (byte, error) {
	return resTestByteSupplier, errTestByteSupplier
}

func TestByteSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSupplier
	}{
		{
			name: "ok",
			s:    testByteSupplier,
		},
		{
			name: "with_error",
			s:    testByteSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestByteSupplier.Error())
			} else {
				r.Equal(resTestByteSupplier, v)
			}
		})
	}
}

func TestByteSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSupplier,
		},
		{
			name: "with_error",
			s:    testByteSupplierWithError,
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
				r.EqualError(err, errTestByteSupplier.Error())
			} else {
				r.Equal(resTestByteSupplier, v)
			}
		})
	}
}

func TestByteSupplier_ToSilentByteSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSupplier,
		},
		{
			name: "with_error",
			s:    testByteSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentByteSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestByteSupplier, v)
			}
		})
	}
}

func TestByteSupplier_ToMustByteSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSupplier,
		},
		{
			name: "with_error",
			s:    testByteSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustByteSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestByteSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestByteSupplier, v)
			}
		})
	}
}

func TestSilentByteSupplier(t *testing.T) {
	var ss SilentByteSupplier = func() byte {
		return resTestByteSupplier
	}
	v := ss()
	require.Equal(t, resTestByteSupplier, v)
}

func TestSilentByteSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSupplier,
		},
		{
			name: "with_error",
			s:    testByteSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentByteSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestByteSupplier, v)
			}
		})
	}
}

func TestMustByteSupplier(t *testing.T) {
	var ms MustByteSupplier = func() byte {
		return resTestByteSupplier
	}

	v := ms()
	require.Equal(t, resTestByteSupplier, v)
}

func TestMustByteSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSupplier,
		},
		{
			name: "with_error",
			s:    testByteSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustByteSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestByteSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestByteSupplier, v)
			}
		})
	}
}

func TestMustByteSupplier_ToSilentByteSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSupplier,
		},
		{
			name: "with_error",
			s:    testByteSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustByteSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentByteSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestByteSupplier, v)
			}
		})
	}
}

func TestMustByteSupplier_ToByteSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSupplier
	}{
		{
			name: "ok",
			s:    testByteSupplier,
		},
		{
			name: "with_error",
			s:    testByteSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustByteSupplier()
			r.NotNil(ms)

			s := ms.ToByteSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestByteSupplier.Error())
			} else {
				r.Equal(resTestByteSupplier, v)
			}
		})
	}
}
