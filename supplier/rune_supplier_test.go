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
	resTestRuneSupplier rune
	errTestRuneSupplier = errors.New("error")
)

func testRuneSupplier() (rune, error) {
	return resTestRuneSupplier, nil
}

func testRuneSupplierWithError() (rune, error) {
	return resTestRuneSupplier, errTestRuneSupplier
}

func TestRuneSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSupplier
	}{
		{
			name: "ok",
			s:    testRuneSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestRuneSupplier.Error())
			} else {
				r.Equal(resTestRuneSupplier, v)
			}
		})
	}
}

func TestRuneSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSupplierWithError,
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
				r.EqualError(err, errTestRuneSupplier.Error())
			} else {
				r.Equal(resTestRuneSupplier, v)
			}
		})
	}
}

func TestRuneSupplier_ToSilentRuneSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentRuneSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSupplier, v)
			}
		})
	}
}

func TestRuneSupplier_ToMustRuneSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestRuneSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestRuneSupplier, v)
			}
		})
	}
}

func TestSilentRuneSupplier(t *testing.T) {
	var ss SilentRuneSupplier = func() rune {
		return resTestRuneSupplier
	}
	v := ss()
	require.Equal(t, resTestRuneSupplier, v)
}

func TestSilentRuneSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentRuneSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSupplier, v)
			}
		})
	}
}

func TestMustRuneSupplier(t *testing.T) {
	var ms MustRuneSupplier = func() rune {
		return resTestRuneSupplier
	}

	v := ms()
	require.Equal(t, resTestRuneSupplier, v)
}

func TestMustRuneSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustRuneSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestRuneSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestRuneSupplier, v)
			}
		})
	}
}

func TestMustRuneSupplier_ToSilentRuneSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentRuneSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSupplier, v)
			}
		})
	}
}

func TestMustRuneSupplier_ToRuneSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSupplier
	}{
		{
			name: "ok",
			s:    testRuneSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSupplier()
			r.NotNil(ms)

			s := ms.ToRuneSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestRuneSupplier.Error())
			} else {
				r.Equal(resTestRuneSupplier, v)
			}
		})
	}
}
