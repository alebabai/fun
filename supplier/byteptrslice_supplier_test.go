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
	testBytePtrSliceSupplierResult []*byte
	testBytePtrSliceSupplierError  = errors.New("error")
)

func testBytePtrSliceSupplier() ([]*byte, error) {
	return testBytePtrSliceSupplierResult, nil
}

func testBytePtrSliceSupplierWithError() ([]*byte, error) {
	return testBytePtrSliceSupplierResult, testBytePtrSliceSupplierError
}

func TestBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBytePtrSliceSupplierError.Error())
			} else {
				r.Equal(testBytePtrSliceSupplierResult, v)
			}
		})
	}
}

func TestBytePtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSliceSupplierWithError,
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
				r.EqualError(err, testBytePtrSliceSupplierError.Error())
			} else {
				r.Equal(testBytePtrSliceSupplierResult, v)
			}
		})
	}
}

func TestBytePtrSliceSupplier_ToSilentBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBytePtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBytePtrSliceSupplierResult, v)
			}
		})
	}
}

func TestBytePtrSliceSupplier_ToMustBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBytePtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBytePtrSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentBytePtrSliceSupplier(t *testing.T) {
	var ss SilentBytePtrSliceSupplier = func() []*byte {
		return testBytePtrSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testBytePtrSliceSupplierResult, v)
}

func TestSilentBytePtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentBytePtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBytePtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBytePtrSliceSupplier(t *testing.T) {
	var ms MustBytePtrSliceSupplier = func() []*byte {
		return testBytePtrSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testBytePtrSliceSupplierResult, v)
}

func TestMustBytePtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustBytePtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testBytePtrSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testBytePtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBytePtrSliceSupplier_ToSilentBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBytePtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testBytePtrSliceSupplierResult, v)
			}
		})
	}
}

func TestMustBytePtrSliceSupplier_ToBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    BytePtrSliceSupplier
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name: "with_error",
			s:    testBytePtrSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToBytePtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testBytePtrSliceSupplierError.Error())
			} else {
				r.Equal(testBytePtrSliceSupplierResult, v)
			}
		})
	}
}
