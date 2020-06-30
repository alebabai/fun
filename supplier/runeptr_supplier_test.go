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
	testRunePtrSupplierResult *rune
	testRunePtrSupplierError  = errors.New("error")
)

func testRunePtrSupplier() (*rune, error) {
	return testRunePtrSupplierResult, nil
}

func testRunePtrSupplierWithError() (*rune, error) {
	return testRunePtrSupplierResult, testRunePtrSupplierError
}

func TestRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSupplier
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testRunePtrSupplierError.Error())
			} else {
				r.Equal(testRunePtrSupplierResult, v)
			}
		})
	}
}

func TestRunePtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSupplierWithError,
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
				r.EqualError(err, testRunePtrSupplierError.Error())
			} else {
				r.Equal(testRunePtrSupplierResult, v)
			}
		})
	}
}

func TestRunePtrSupplier_ToSilentRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentRunePtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testRunePtrSupplierResult, v)
			}
		})
	}
}

func TestRunePtrSupplier_ToMustRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testRunePtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testRunePtrSupplierResult, v)
			}
		})
	}
}

func TestSilentRunePtrSupplier(t *testing.T) {
	var ss SilentRunePtrSupplier = func() *rune {
		return testRunePtrSupplierResult
	}
	v := ss()
	require.Equal(t, testRunePtrSupplierResult, v)
}

func TestSilentRunePtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentRunePtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testRunePtrSupplierResult, v)
			}
		})
	}
}

func TestMustRunePtrSupplier(t *testing.T) {
	var ms MustRunePtrSupplier = func() *rune {
		return testRunePtrSupplierResult
	}

	v := ms()
	require.Equal(t, testRunePtrSupplierResult, v)
}

func TestMustRunePtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustRunePtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testRunePtrSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testRunePtrSupplierResult, v)
			}
		})
	}
}

func TestMustRunePtrSupplier_ToSilentRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentRunePtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testRunePtrSupplierResult, v)
			}
		})
	}
}

func TestMustRunePtrSupplier_ToRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSupplier
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSupplier()
			r.NotNil(ms)

			s := ms.ToRunePtrSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testRunePtrSupplierError.Error())
			} else {
				r.Equal(testRunePtrSupplierResult, v)
			}
		})
	}
}
