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
	testByteSliceSupplierResult []byte
	testByteSliceSupplierError  = errors.New("error")
)

func testByteSliceSupplier() ([]byte, error) {
	return testByteSliceSupplierResult, nil
}

func testByteSliceSupplierWithError() ([]byte, error) {
	return testByteSliceSupplierResult, testByteSliceSupplierError
}

func TestByteSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSliceSupplier
	}{
		{
			name: "ok",
			s:    testByteSliceSupplier,
		},
		{
			name: "with_error",
			s:    testByteSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testByteSliceSupplierError.Error())
			} else {
				r.Equal(testByteSliceSupplierResult, v)
			}
		})
	}
}

func TestByteSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSliceSupplier,
		},
		{
			name: "with_error",
			s:    testByteSliceSupplierWithError,
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
				r.EqualError(err, testByteSliceSupplierError.Error())
			} else {
				r.Equal(testByteSliceSupplierResult, v)
			}
		})
	}
}

func TestByteSliceSupplier_ToSilentByteSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSliceSupplier,
		},
		{
			name: "with_error",
			s:    testByteSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentByteSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testByteSliceSupplierResult, v)
			}
		})
	}
}

func TestByteSliceSupplier_ToMustByteSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSliceSupplier,
		},
		{
			name: "with_error",
			s:    testByteSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustByteSliceSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testByteSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testByteSliceSupplierResult, v)
			}
		})
	}
}

func TestSilentByteSliceSupplier(t *testing.T) {
	var ss SilentByteSliceSupplier = func() []byte {
		return testByteSliceSupplierResult
	}
	v := ss()
	require.Equal(t, testByteSliceSupplierResult, v)
}

func TestSilentByteSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSliceSupplier,
		},
		{
			name: "with_error",
			s:    testByteSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentByteSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testByteSliceSupplierResult, v)
			}
		})
	}
}

func TestMustByteSliceSupplier(t *testing.T) {
	var ms MustByteSliceSupplier = func() []byte {
		return testByteSliceSupplierResult
	}

	v := ms()
	require.Equal(t, testByteSliceSupplierResult, v)
}

func TestMustByteSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSliceSupplier,
		},
		{
			name: "with_error",
			s:    testByteSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustByteSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.err {
				r.PanicsWithError(testByteSliceSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testByteSliceSupplierResult, v)
			}
		})
	}
}

func TestMustByteSliceSupplier_ToSilentByteSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSliceSupplier
		err  bool
	}{
		{
			name: "ok",
			s:    testByteSliceSupplier,
		},
		{
			name: "with_error",
			s:    testByteSliceSupplierWithError,
			err:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustByteSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentByteSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.err {
				r.Empty(v)
			} else {
				r.Equal(testByteSliceSupplierResult, v)
			}
		})
	}
}

func TestMustByteSliceSupplier_ToByteSliceSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    ByteSliceSupplier
	}{
		{
			name: "ok",
			s:    testByteSliceSupplier,
		},
		{
			name: "with_error",
			s:    testByteSliceSupplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustByteSliceSupplier()
			r.NotNil(ms)

			s := ms.ToByteSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.Empty(v)
				r.EqualError(err, testByteSliceSupplierError.Error())
			} else {
				r.Equal(testByteSliceSupplierResult, v)
			}
		})
	}
}
