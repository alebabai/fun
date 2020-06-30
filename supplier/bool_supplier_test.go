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
	testBoolSupplierResult bool
	testBoolSupplierError  = errors.New("error")
)

func testBoolSupplier() (bool, error) {
	return testBoolSupplierResult, nil
}

func testBoolSupplierWithError() (bool, error) {
	return testBoolSupplierResult, testBoolSupplierError
}

func TestBoolSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSupplier
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBoolSupplierError.Error())
			} else {
				r.Equal(testBoolSupplierResult, v)
			}
		})
	}
}

func TestBoolSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSupplierWithError,
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
				r.EqualError(err, testBoolSupplierError.Error())
			} else {
				r.Equal(testBoolSupplierResult, v)
			}
		})
	}
}

func TestBoolSupplier_ToSilentBoolSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBoolSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolSupplierResult, v)
			}
		})
	}
}

func TestBoolSupplier_ToMustBoolSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBoolSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBoolSupplierResult, v)
			}
		})
	}
}

func TestSilentBoolSupplier(t *testing.T) {
	var ss SilentBoolSupplier = func() bool {
		return testBoolSupplierResult
	}
	v := ss()
	require.Equal(t, testBoolSupplierResult, v)
}

func TestSilentBoolSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentBoolSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolSupplierResult, v)
			}
		})
	}
}

func TestMustBoolSupplier(t *testing.T) {
	var ms MustBoolSupplier = func() bool {
		return testBoolSupplierResult
	}

	v := ms()
	require.Equal(t, testBoolSupplierResult, v)
}

func TestMustBoolSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustBoolSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBoolSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBoolSupplierResult, v)
			}
		})
	}
}

func TestMustBoolSupplier_ToSilentBoolSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBoolSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBoolSupplierResult, v)
			}
		})
	}
}

func TestMustBoolSupplier_ToBoolSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BoolSupplier
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name: "with_error",
			s:    testBoolSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSupplier()
			r.NotNil(ms)

			s := ms.ToBoolSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBoolSupplierError.Error())
			} else {
				r.Equal(testBoolSupplierResult, v)
			}
		})
	}
}
