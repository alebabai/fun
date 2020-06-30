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
	testRuneSliceSupplierResult []rune
	testRuneSliceSupplierError  = errors.New("error")
)

func testRuneSliceSupplier() ([]rune, error) {
	return testRuneSliceSupplierResult, nil
}

func testRuneSliceSupplierWithError() ([]rune, error) {
	return testRuneSliceSupplierResult, testRuneSliceSupplierError
}

func TestRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSliceSupplier
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testRuneSliceSupplierError.Error())
			} else {
				r.Equal(testRuneSliceSupplierResult, v)
			}
		})
	}
}

func TestRuneSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSliceSupplierWithError,
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
				r.EqualError(err, testRuneSliceSupplierError.Error())
			} else {
				r.Equal(testRuneSliceSupplierResult, v)
			}
		})
	}
}

func TestRuneSliceSupplier_ToSilentRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentRuneSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testRuneSliceSupplierResult, v)
			}
		})
	}
}

func TestRuneSliceSupplier_ToMustRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testRuneSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testRuneSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentRuneSliceSupplier(t *testing.T) {
	var ss SilentRuneSliceSupplier = func() []rune {
		return testRuneSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testRuneSliceSupplierResult, v)
}

func TestSilentRuneSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentRuneSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testRuneSliceSupplierResult, v)
			}
		})
	}
}

func TestMustRuneSliceSupplier(t *testing.T) {
	var ms MustRuneSliceSupplier = func() []rune {
		return testRuneSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testRuneSliceSupplierResult, v)
}

func TestMustRuneSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustRuneSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testRuneSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testRuneSliceSupplierResult, v)
			}
		})
	}
}

func TestMustRuneSliceSupplier_ToSilentRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentRuneSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testRuneSliceSupplierResult, v)
			}
		})
	}
}

func TestMustRuneSliceSupplier_ToRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RuneSliceSupplier
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRuneSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSliceSupplier()
			r.NotNil(ms)

			s := ms.ToRuneSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testRuneSliceSupplierError.Error())
			} else {
				r.Equal(testRuneSliceSupplierResult, v)
			}
		})
	}
}
