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
	resTestRunePtrSliceSupplier []*rune
	errTestRunePtrSliceSupplier = errors.New("error")
)

func testRunePtrSliceSupplier() ([]*rune, error) {
	return resTestRunePtrSliceSupplier, nil
}

func testRunePtrSliceSupplierWithError() ([]*rune, error) {
	return resTestRunePtrSliceSupplier, errTestRunePtrSliceSupplier
}

func TestRunePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testRunePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestRunePtrSliceSupplier.Error())
			} else {
				r.Equal(resTestRunePtrSliceSupplier, v)
			}
		})
	}
}

func TestRunePtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSliceSupplierWithError,
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
				r.EqualError(err, errTestRunePtrSliceSupplier.Error())
			} else {
				r.Equal(resTestRunePtrSliceSupplier, v)
			}
		})
	}
}

func TestRunePtrSliceSupplier_ToSilentRunePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentRunePtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSliceSupplier, v)
			}
		})
	}
}

func TestRunePtrSliceSupplier_ToMustRunePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestRunePtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestRunePtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentRunePtrSliceSupplier(t *testing.T) {
	var ss SilentRunePtrSliceSupplier = func() []*rune {
		return resTestRunePtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestRunePtrSliceSupplier, v)
}

func TestSilentRunePtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentRunePtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSliceSupplier, v)
			}
		})
	}
}

func TestMustRunePtrSliceSupplier(t *testing.T) {
	var ms MustRunePtrSliceSupplier = func() []*rune {
		return resTestRunePtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestRunePtrSliceSupplier, v)
}

func TestMustRunePtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustRunePtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(errTestRunePtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestRunePtrSliceSupplier, v)
			}
		})
	}
}

func TestMustRunePtrSliceSupplier_ToSilentRunePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testRunePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentRunePtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSliceSupplier, v)
			}
		})
	}
}

func TestMustRunePtrSliceSupplier_ToRunePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    RunePtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testRunePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testRunePtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToRunePtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, errTestRunePtrSliceSupplier.Error())
			} else {
				r.Equal(resTestRunePtrSliceSupplier, v)
			}
		})
	}
}
