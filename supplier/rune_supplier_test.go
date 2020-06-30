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
	testRuneSupplierResult rune
	testRuneSupplierError  = errors.New("error")
)

func testRuneSupplier() (rune, error) {
	return testRuneSupplierResult, nil
}

func testRuneSupplierWithError() (rune, error) {
	return testRuneSupplierResult, testRuneSupplierError
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
				r.EqualError(err, testRuneSupplierError.Error())
			} else {
				r.Equal(testRuneSupplierResult, v)
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
				r.EqualError(err, testRuneSupplierError.Error())
			} else {
				r.Equal(testRuneSupplierResult, v)
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
				r.Equal(testRuneSupplierResult, v)
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
				r.PanicsWithError(testRuneSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testRuneSupplierResult, v)
			}
		})
	}
}

func TestSilentRuneSupplier(t *testing.T) {
	var ss SilentRuneSupplier = func() rune {
		return testRuneSupplierResult
	}
	v := ss()
	require.Equal(t, testRuneSupplierResult, v)
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
				r.Equal(testRuneSupplierResult, v)
			}
		})
	}
}

func TestMustRuneSupplier(t *testing.T) {
	var ms MustRuneSupplier = func() rune {
		return testRuneSupplierResult
	}

	v := ms()
	require.Equal(t, testRuneSupplierResult, v)
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
				r.PanicsWithError(testRuneSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testRuneSupplierResult, v)
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
				r.Equal(testRuneSupplierResult, v)
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
				r.EqualError(err, testRuneSupplierError.Error())
			} else {
				r.Equal(testRuneSupplierResult, v)
			}
		})
	}
}
